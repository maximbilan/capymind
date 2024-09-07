package scheduler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/capymind/internal/analysis"
	"github.com/capymind/internal/firestore"
	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
)

func Schedule(w http.ResponseWriter, r *http.Request) {
	log.Println("Schedule capymind...")

	typeStr := r.URL.Query().Get("type")
	messageType := MessageType(typeStr)

	var message string
	switch messageType {
	case Morning:
		message = "how_are_you_morning"
	case Evening:
		message = "how_are_you_evening"
	case WeeklyAnalysis:
		// Personalized for each user
		message = ""
	default:
		log.Println("Missing message type parameter")
		return
	}

	ctx := context.Background()

	// Firestore
	dbClient := createDBClient(ctx)
	defer dbClient.Close()

	// Cloud Tasks
	tasksClient := createTasksClient(ctx)
	defer tasksClient.Close()

	var isCloud = false
	if os.Getenv("CLOUD") == "true" {
		isCloud = true
	}

	firestore.ForEachUser(ctx, dbClient, func(users []firestore.User) error {
		for _, user := range users {
			log.Printf("[Scheduler] Schedule a message for user: %s", user.ID)
			if user.LastChatId == nil || user.Locale == nil || user.SecondsFromUTC == nil {
				continue
			}

			userLocale := translator.Locale(*user.Locale)

			var localizedMessage string
			if messageType == WeeklyAnalysis {
				notes, err := firestore.GetNotes(ctx, dbClient, user.ID)
				if err != nil {
					log.Printf("[Scheduler] Error getting notes from firestore, %s", err.Error())
					continue
				}

				if len(notes) > 0 {
					var strings []string
					for _, note := range notes {
						if note.Text != "" {
							strings = append(strings, note.Text)
						}
					}
					localizedMessage = *analysis.Request(strings, userLocale)
				} else {
					continue
				}
			} else {
				localizedMessage = translator.Translate(userLocale, message)
			}

			var scheduledTime time.Time
			if isCloud {
				scheduledTime = time.Now().Add(9 * time.Hour)
				scheduledTime = scheduledTime.Add(-time.Duration(*user.SecondsFromUTC) * time.Second)
			} else {
				// For local testing, schedule the message in 10 seconds
				scheduledTime = time.Now().Add(10 * time.Second)
			}

			scheduledMessage := ScheduledMessage{
				ChatId: *user.LastChatId,
				Text:   localizedMessage,
				Type:   messageType,
				Locale: userLocale,
			}

			scheduleTask(ctx, tasksClient, scheduledMessage, scheduledTime)
		}
		return nil
	})
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	var msg ScheduledMessage
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		log.Printf("[Scheduler] Could not parse message %s", err.Error())
		return
	}

	var reply *telegram.InlineKeyboardMarkup
	switch msg.Type {
	case Morning, Evening:
		reply = &telegram.InlineKeyboardMarkup{
			InlineKeyboard: [][]telegram.InlineKeyboardButton{
				{
					{Text: translator.Translate(msg.Locale, "make_record_to_journal"), CallbackData: "note_make"},
				},
			},
		}
	default:
		reply = nil
	}

	telegram.SendMessage(msg.ChatId, msg.Text, reply)
}
