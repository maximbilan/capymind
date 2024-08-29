package scheduler

import (
	"context"
	"log"
	"net/http"

	google "cloud.google.com/go/firestore"
	"github.com/capymind/internal/firestore"
	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
)

func createClient() (*google.Client, context.Context) {
	ctx := context.Background()
	var client, err = firestore.NewClient(ctx)
	if err != nil {
		log.Printf("[Scheduler] Error creating firestore client, %s", err.Error())
	}
	return client, ctx
}

func Schedule(w http.ResponseWriter, r *http.Request) {
	log.Println("Schedule capymind...")

	client, ctx := createClient()
	defer client.Close()

	firestore.ForEachUser(ctx, client, func(users []firestore.User) error {
		for _, user := range users {
			log.Printf("[Scheduler] User: %s", user.ID)
			// Handle empty fields later
			userLocale := translator.Locale(user.Locale)
			localizedMessage := translator.Translate(userLocale, "how_are_you")
			telegram.SendMessage(user.LastChatId, localizedMessage, nil)
		}
		return nil
	})
}
