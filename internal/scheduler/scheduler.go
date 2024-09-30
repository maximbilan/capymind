package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/capymind/internal/analysis"
	"github.com/capymind/internal/firestore"
	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
)

// Schedule a message for all users
func Schedule(w http.ResponseWriter, r *http.Request) {
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
	messageType := MessageType(typeStr)

	var message string
	switch messageType {
	case Morning:
		message = "how_are_you_morning"
	case Evening:
		message = "how_are_you_evening"
	case WeeklyAnalysis, UserStats:
		// Personalized for each user
		message = ""
	default:
		log.Println("Missing message type parameter")
		return
	}

	ctx := context.Background()
	firestore.CreateClient(&ctx)

	// Cloud Tasks
	CreateTasks(&ctx)

	var isCloud = false
	if os.Getenv("CLOUD") == "true" {
		isCloud = true
	}

	firestore.ForEachUser(&ctx, func(users []firestore.User) error {
		for _, user := range users {
			log.Printf("[Scheduler] Schedule a message for user: %s", user.ID)
			if user.Locale == nil || user.SecondsFromUTC == nil {
				continue
			}

			userLocale := translator.Locale(*user.Locale)

			var localizedMessage string
			if messageType == WeeklyAnalysis {
				notes, err := firestore.GetNotesForLastWeek(&ctx, user.ID)
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
					header := "weekly_analysis"
					localizedMessage = *analysis.AnalyzeJournal(strings, userLocale, &ctx, &header)
				} else {
					continue
				}
			} else if messageType == UserStats {
				count, err := firestore.NotesCount(&ctx, user.ID)
				if err != nil {
					log.Printf("[Scheduler] Error getting notes count from firestore, %s", err.Error())
					continue
				}
				// Send only if the user has more than one note in the journal
				if count > 1 {
					localizedMessage = fmt.Sprintf(translator.Translate(userLocale, "user_progress_message"), count)
				} else {
					continue
				}
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

			scheduledMessage := ScheduledMessage{
				ChatID: user.ChatID,
				Text:   localizedMessage,
				Type:   messageType,
				Locale: userLocale,
			}

			scheduleTask(&ctx, scheduledMessage, scheduledTime)
		}
		return nil
	})

	// Close Firestore client
	CloseTasks()
	// Close Tasks client
	firestore.CloseClient()
}

// Send a message to a user
func SendMessage(w http.ResponseWriter, r *http.Request) {
	var msg ScheduledMessage
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		log.Printf("[Scheduler] Could not parse message %s", err.Error())
		return
	}

	var reply *telegram.InlineKeyboardMarkup
	switch msg.Type {
	case Morning, Evening:
		callbackData := "/note"
		reply = &telegram.InlineKeyboardMarkup{
			InlineKeyboard: [][]telegram.InlineKeyboardButton{
				{
					{Text: translator.Translate(msg.Locale, "make_record_to_journal"), CallbackData: &callbackData},
				},
			},
		}
	default:
		reply = nil
	}

	telegram.SendMessage(msg.ChatID, msg.Text, reply)
}
