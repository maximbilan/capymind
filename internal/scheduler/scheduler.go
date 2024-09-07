package scheduler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/capymind/internal/firestore"
	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
)

func Schedule(w http.ResponseWriter, r *http.Request) {
	log.Println("Schedule capymind...")

	ctx := context.Background()

	// Firestore
	dbClient := createDBClient(ctx)
	defer dbClient.Close()

	// Cloud Tasks
	tasksClient := createTasksClient(ctx)
	defer tasksClient.Close()

	firestore.ForEachUser(ctx, dbClient, func(users []firestore.User) error {
		for _, user := range users {
			log.Printf("[Scheduler] Schedule a message for user: %s", user.ID)
			if user.LastChatId == nil || user.Locale == nil || user.SecondsFromUTC == nil {
				break
			}

			userLocale := translator.Locale(*user.Locale)
			localizedMessage := translator.Translate(userLocale, "how_are_you")
			scheduledTime := time.Now().Add(9 * time.Hour)
			scheduledTime = scheduledTime.Add(-time.Duration(*user.SecondsFromUTC) * time.Second)

			scheduleTask(ctx, tasksClient, *user.LastChatId, localizedMessage, scheduledTime)
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
	telegram.SendMessage(msg.ChatId, msg.Text, nil)
}
