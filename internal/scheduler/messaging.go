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

	var localizedMessage *string
	if messageType == taskservice.WeeklyAnalysis {
		localizedMessage = prepareWeeklyAnalysis(user, ctx, userLocale, aiService)
	} else if messageType == taskservice.UserStats {
		// Send only to active users
		if !user.IsActive() {
			return
		}
		localizedMessage = prepareUserStats(user, ctx, userLocale)
	} else if messageType == taskservice.AdminStats {
		// Send only to admins
		if !database.IsAdmin(user.Role) {
			return
		}
		localizedMessage = prepareAdminStats(user, ctx, userLocale)
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
		return
	}

	result := prepareBotResult(msg)
	bot.SendResult(msg.ChatID, result)
	log.Printf("[Scheduler] Message sent to user: %d", msg.ChatID)
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
