package bot

import (
	"context"
	"log"

	google "cloud.google.com/go/firestore"
	"github.com/capymind/internal/firestore"
)

func createClient() (*google.Client, context.Context) {
	ctx := context.Background()
	var client, err = firestore.NewClient(ctx)
	if err != nil {
		log.Printf("[Database] Error creating firestore client, %s", err.Error())
	}
	return client, ctx
}
