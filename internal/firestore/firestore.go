package firestore

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
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
	var client *firestore.Client
	var err error

	if os.Getenv("CLOUD") == "true" {
		client, err = firestore.NewClient(ctx, projectID)
	} else {
		path := credentialsPath()
		client, err = firestore.NewClient(ctx, projectID, option.WithCredentialsFile(path))
	}

	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewRecord(ctx context.Context, client *firestore.Client, user User, note Note) error {
	err := NewUser(ctx, client, user)
	if err != nil {
		return err
	}
	err = newNote(ctx, client, user, note)
	if err != nil {
		return err
	}
	return nil
}

func NewUser(ctx context.Context, client *firestore.Client, user User) error {
	_, err := client.Collection(users.String()).Doc(user.ID).Set(ctx, map[string]interface{}{
		"name": user.Name,
	}, firestore.MergeAll)
	return err
}

func newNote(ctx context.Context, client *firestore.Client, user User, note Note) error {
	userRef := client.Collection(users.String()).Doc(user.ID)
	_, _, err := client.Collection(notes.String()).Add(ctx, map[string]interface{}{
		"text":      note.Text,
		"timestamp": note.Timestamp,
		"user":      userRef,
	})
	return err
}

func LastNote(ctx context.Context, client *firestore.Client, userId string) (*Note, error) {
	userRef := client.Collection(users.String()).Doc(userId)
	query := client.Collection(notes.String()).OrderBy("timestamp", firestore.Desc).Where("user", "==", userRef).Limit(1)

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

func GetNotes(ctx context.Context, client *firestore.Client, userId string) ([]Note, error) {
	userRef := client.Collection(users.String()).Doc(userId)
	query := client.Collection(notes.String()).Where("user", "==", userRef).Limit(100)

	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	var notes []Note
	for _, doc := range docs {
		var note Note
		doc.DataTo(&note)
		notes = append(notes, note)
	}
	return notes, nil
}
