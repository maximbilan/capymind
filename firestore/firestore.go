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
	ID   string `firestore:"id"`
	Name string `firestore:"name"`
}

const (
	users string = "users"
	notes string = "notes"
)

func credentialsPath() string {
	var path = "credentials.json"
	if os.Getenv("DEBUG_MODE") == "true" {
		path = "./../" + path
	} else {
		path = "./" + path
	}
	return path
}

func NewClient(ctx context.Context) (*firestore.Client, error) {
	projectID := os.Getenv("CAPY_PROJECT_ID")
	path := credentialsPath()
	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(path))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewRecord(ctx context.Context, client *firestore.Client, user User, note Note) error {
	err := newUser(ctx, client, user)
	if err != nil {
		return err
	}
	err = newNote(ctx, client, user, note)
	if err != nil {
		return err
	}
	return nil
}

func newUser(ctx context.Context, client *firestore.Client, user User) error {
	_, err := client.Collection(users).Doc(user.ID).Set(ctx, map[string]interface{}{
		"name": user.Name,
	})
	return err
}

func newNote(ctx context.Context, client *firestore.Client, user User, note Note) error {
	userRef := client.Collection(users).Doc(user.ID)
	_, _, err := client.Collection(notes).Add(ctx, map[string]interface{}{
		"text":      note.Text,
		"timestamp": note.Timestamp,
		"user":      userRef,
	})
	return err
}

func LastNote(ctx context.Context, client *firestore.Client, userId string) (*Note, error) {
	userRef := client.Collection(users).Doc(userId)
	query := client.Collection(notes).OrderBy("timestamp", firestore.Desc).Where("user", "==", userRef).Limit(1)

	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	if len(docs) > 0 {
		var note Note
		docs[0].DataTo(&note)
		return &note, nil
	}
	return nil, nil
}