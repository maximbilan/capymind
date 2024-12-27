package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/capymind/internal/analysis"
	"github.com/capymind/internal/botservice"
	"github.com/capymind/internal/database"
	"github.com/capymind/internal/helpers"
	"github.com/capymind/internal/taskservice"
	"github.com/capymind/internal/translator"
	"github.com/capymind/third_party/googletasks"
)

var wg sync.WaitGroup
var tasks googletasks.GoogleTasks

// Schedule a message for all users
func Schedule(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	log.Println("Schedule capymind...")

	typeStr := r.URL.Query().Get("type")
	offsetStr := r.URL.Query().Get("offset") // hours (from UTC 0)
	var offset int = 0
	if offsetStr != "" {
		_, err := fmt.Sscanf(offsetStr, "%d", &offset)
		if err != nil {
			log.Printf("[Scheduler] Error getting offset parameter, %s", err.Error())
		}
	}
	messageType := taskservice.MessageType(typeStr)

	var message string
	switch messageType {
	case taskservice.Morning, taskservice.Evening:
		message = taskservice.GetMessage(messageType, time.Now().Weekday())
	case taskservice.Feedback:
		message = ""
	case taskservice.WeeklyAnalysis, taskservice.UserStats, taskservice.AdminStats:
		// Personalized for each user
		message = ""
	default:
		log.Println("Missing message type parameter")
		return
	}

	ctx := context.Background()
	db.CreateClient(&ctx)
	tasks.Connect(&ctx)

	var isCloud = false
	if os.Getenv("CLOUD") == "true" {
		isCloud = true
	}

	userStorage.ForEachUser(&ctx, func(users []database.User) error {
		for _, user := range users {
			// Skip users without locale or timezone
			if user.Locale == nil || user.SecondsFromUTC == nil {
				continue
			}
			// For debugging locally
			if !isCloud && !database.IsAdmin(user.Role) {
				continue
			}

			// Prepare message concurrently
			wg.Add(1)
			go prepareMessage(&user, &ctx, offset, messageType, message, isCloud)
		}
		return nil
	})

	wg.Wait()

	// Close Firestore client
	tasks.Close()
	// Close Tasks client
	db.CloseClient()

	// Calculate how seconds this function takes to execute
	elapsed := time.Since(start)
	log.Printf("[Scheduler] Execution time for %s: %s", messageType, elapsed)
}

func prepareMessage(user *database.User, ctx *context.Context, offset int, messageType taskservice.MessageType, message string, isCloud bool) {
	defer wg.Done()

	log.Printf("[Scheduler] Schedule a message for user: %s", user.ID)

	userLocale := translator.Locale(*user.Locale)

	var localizedMessage string
	if messageType == taskservice.WeeklyAnalysis {
		notes, err := noteStorage.GetNotesForLastWeek(ctx, user.ID)
		if err != nil {
			log.Printf("[Scheduler] Error getting notes from firestore, %s", err.Error())
			return
		}

		if len(notes) > 0 {
			var strings []string
			for _, note := range notes {
				if note.Text != "" {
					strings = append(strings, note.Text)
				}
			}
			localizedMessage = *analysis.AnalyzeLastWeek(strings, userLocale, ctx)
		} else {
			return
		}
	} else if messageType == taskservice.UserStats {
		count, err := noteStorage.NotesCount(ctx, user.ID)
		if err != nil {
			log.Printf("[Scheduler] Error getting notes count from firestore, %s", err.Error())
			return
		}
		// Send only if the user has more than one note in the journal
		if count > 1 {
			localizedMessage = fmt.Sprintf(translator.Translate(userLocale, "user_progress_message"), count)
		} else {
			return
		}
	} else if messageType == taskservice.AdminStats {
		// Send only to admins
		if !database.IsAdmin(user.Role) {
			return
		}
		stats := helpers.GetStats(ctx, userLocale)

		var finalString string
		for _, stat := range stats {
			finalString += stat + "\n"
		}
		localizedMessage = finalString
	} else {
		localizedMessage = translator.Translate(userLocale, message)
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
		Text:   localizedMessage,
		Type:   messageType,
		Locale: userLocale,
	}

	tasks.Schedule(ctx, scheduledMessage, scheduledTime)
}

// Send a message to a user
func SendMessage(w http.ResponseWriter, r *http.Request) {
	var msg taskservice.ScheduledTask
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		log.Printf("[Scheduler] Could not parse message %s", err.Error())
		return
	}

	switch msg.Type {
	case taskservice.Morning, taskservice.Evening:
		var button botservice.BotResultTextButton = botservice.BotResultTextButton{
			TextID:   "make_record_to_journal",
			Locale:   msg.Locale,
			Callback: "/note",
		}
		result := botservice.BotResult{
			TextID:  msg.Text,
			Locale:  msg.Locale,
			Buttons: []botservice.BotResultTextButton{button},
		}
		bot.SendResult(msg.ChatID, result)
	case taskservice.Feedback:
		var button botservice.BotResultTextButton = botservice.BotResultTextButton{
			TextID:   "feedback_button",
			Locale:   msg.Locale,
			Callback: "/support",
		}
		result := botservice.BotResult{
			TextID:  msg.Text,
			Locale:  msg.Locale,
			Buttons: []botservice.BotResultTextButton{button},
		}
		bot.SendResult(msg.ChatID, result)
	default:
		bot.SendMessage(msg.ChatID, msg.Text)
	}
}
