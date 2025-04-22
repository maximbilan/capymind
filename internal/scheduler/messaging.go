package scheduler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/capymind/internal/botservice"
	"github.com/capymind/internal/database"
	"github.com/capymind/internal/taskservice"
	"github.com/capymind/internal/translator"
)

func prepareMessage(user *database.User, ctx *context.Context, offset int, messageType taskservice.MessageType, message string, isCloud bool) {
	//coverage:ignore
	defer wg.Done()

	log.Printf("[Scheduler] Schedule a message for user: %s", user.ID)

	userLocale := translator.Locale(*user.Locale)
	settings, _ := settingsStorage.GetSettings(ctx, user.ID)
	if settings == nil {
		// Create an empty settings object if it does not exist
		emptySettings := database.Settings{}
		settings = &emptySettings
	}

	var localizedMessage *string
	if messageType == taskservice.WeeklyAnalysis {
		localizedMessage = prepareWeeklyAnalysis(user, ctx, userLocale, noteStorage, aiService)
	} else if messageType == taskservice.UserStats {
		// Send only to active users
		if !user.IsActive() {
			return
		}
		localizedMessage = prepareUserStats(user, ctx, userLocale, noteStorage)
	} else if messageType == taskservice.AdminStats {
		// Send only to admins
		if !database.IsAdmin(user.Role) {
			return
		}
		localizedMessage = prepareAdminStats(ctx, userLocale, adminStorage, feedbackStorage)
	} else if messageType == taskservice.Feedback {
		// Send only to active users
		if user.IsNonActive() {
			return
		}
		msg := translator.Translate(userLocale, message)
		localizedMessage = &msg
	} else if messageType == taskservice.Morning {
		// Send only to active users
		if user.IsNonActive() {
			return
		}
		// Send only if the morning reminder is enabled
		if !settings.IsMorningReminderEnabled() {
			return
		}
		msg := translator.Translate(userLocale, message)
		localizedMessage = &msg
		if settings.MorningReminderOffset != nil {
			offset = *settings.MorningReminderOffset
		}
	} else if messageType == taskservice.Evening {
		// Send only to active users
		if user.IsNonActive() {
			return
		}
		// Send only if the evening reminder is enabled
		if !settings.IsEveningReminderEnabled() {
			return
		}
		msg := translator.Translate(userLocale, message)
		localizedMessage = &msg
		if settings.EveningReminderOffset != nil {
			offset = *settings.EveningReminderOffset
		}
	} else {
		msg := translator.Translate(userLocale, message)
		localizedMessage = &msg
	}

	if localizedMessage == nil {
		return
	}

	var scheduledTime time.Time
	if isCloud {
		scheduledTime = time.Now().Add(time.Duration(offset) * time.Hour)
		scheduledTime = scheduledTime.Add(-time.Duration(*user.SecondsFromUTC) * time.Second)
	} else {
		// For local testing, schedule the message in 10 seconds
		scheduledTime = time.Now().Add(10 * time.Second)
	}

	scheduledMessage := taskservice.ScheduledTask{
		ChatID: user.ChatID,
		Text:   *localizedMessage,
		Type:   messageType,
		Locale: userLocale,
	}

	tasks.Schedule(ctx, scheduledMessage, scheduledTime)
}

// Send a message to a user
func SendMessage(w http.ResponseWriter, r *http.Request) {
	//coverage:ignore
	var msg taskservice.ScheduledTask
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		log.Printf("[Scheduler] Could not parse message %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not parse message"))
		return
	}

	result := prepareBotResult(msg)
	bot.SendResult(msg.ChatID, result)
	log.Printf("[Scheduler] Message sent to user: %d", msg.ChatID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func prepareBotResult(scheduledTask taskservice.ScheduledTask) botservice.BotResult {
	var result botservice.BotResult
	switch scheduledTask.Type {
	case taskservice.Morning, taskservice.Evening:
		var button botservice.BotResultTextButton = botservice.BotResultTextButton{
			TextID:   "make_record_to_journal",
			Locale:   scheduledTask.Locale,
			Callback: "/note",
		}
		result = botservice.BotResult{
			TextID:  scheduledTask.Text,
			Locale:  scheduledTask.Locale,
			Buttons: []botservice.BotResultTextButton{button},
		}
	case taskservice.Feedback:
		var button botservice.BotResultTextButton = botservice.BotResultTextButton{
			TextID:   "feedback_button",
			Locale:   scheduledTask.Locale,
			Callback: "/support",
		}
		result = botservice.BotResult{
			TextID:  scheduledTask.Text,
			Locale:  scheduledTask.Locale,
			Buttons: []botservice.BotResultTextButton{button},
		}
	default:
		result = botservice.BotResult{
			TextID: scheduledTask.Text,
			Locale: scheduledTask.Locale,
		}
	}
	return result
}
