package scheduler

import (
	"context"
	"log"

	google "cloud.google.com/go/firestore"
	"github.com/capymind/internal/firestore"
	"github.com/capymind/internal/telegram"
)

func createClient() (*google.Client, context.Context) {
	ctx := context.Background()
	var client, err = firestore.NewClient(ctx)
	if err != nil {
		log.Printf("[Scheduler] Error creating firestore client, %s", err.Error())
	}
	return client, ctx
}

func Schedule() {
	log.Println("Schedule capymind...")

	client, ctx := createClient()
	defer client.Close()

	firestore.ForEachUser(ctx, client, func(users []firestore.User) error {
		for _, user := range users {
			log.Printf("[Scheduler] User: %s", user.ID)
			telegram.SendMessage(user.LastChatId, "Hello from scheduler!", nil)
		}
		return nil
	})
}
