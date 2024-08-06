package firestore

import (
	"context"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
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
	projectID := os.Getenv("CAPY_PROJECT_ID")
	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile("./credentials.json"))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func SaveNote(ctx context.Context, client *firestore.Client, userId string, note Note) error {
	userRef := client.Collection("users").Doc(userId)
	_, _, err := client.Collection("notes").Add(ctx, map[string]interface{}{
		"id":        note.ID,
		"text":      note.Text,
		"timestamp": note.Timestamp,
		"user":      userRef,
	})
	return err
}
