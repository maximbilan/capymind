package scheduler

import (
	"context"
	"log"

	firestoreDB "cloud.google.com/go/firestore"
	"github.com/capymind/internal/firestore"
)

// Create a Firestore client
func createDBClient(ctx context.Context) *firestoreDB.Client {
	var client, err = firestore.NewClient(ctx)
	if err != nil {
		log.Printf("[Scheduler] Error creating firestore client, %s", err.Error())
	}
	return client
}
