package firestore

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

type Note struct {
	ID        string                 `firestore:"id"`
	Text      string                 `firestore:"text"`
	Timestamp time.Time              `firestore:"timestamp"`
	User      *firestore.DocumentRef `firestore:"user"`
}

type User struct {
	ID string `firestore:"id"`
}

func NewClient(ctx context.Context) (*firestore.Client, error) {
	projectID := "YOUR_PROJECT_ID"

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func SaveNote(ctx context.Context, client *firestore.Client, note Note) error {
	userRef := client.Collection("users").Doc(note.ID)
	_, _, err := client.Collection("notes").Add(ctx, map[string]interface{}{
		"id":        note.ID,
		"text":      note.Text,
		"timestamp": note.Timestamp,
		"user":      userRef,
	})
	return err
}
