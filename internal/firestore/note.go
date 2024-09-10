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

func NewNote(ctx context.Context, client *firestore.Client, user User, note Note) error {
	userRef := client.Collection(users.String()).Doc(user.ID)
	_, _, err := client.Collection(notes.String()).Add(ctx, map[string]interface{}{
		"text":      note.Text,
		"timestamp": note.Timestamp,
		"user":      userRef,
	})
	return err
}

func LastNote(ctx context.Context, client *firestore.Client, userID string) (*Note, error) {
	userRef := client.Collection(users.String()).Doc(userID)
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

func GetNotes(ctx context.Context, client *firestore.Client, userID string) ([]Note, error) {
	userRef := client.Collection(users.String()).Doc(userID)
	query := client.Collection(notes.String()).OrderBy("timestamp", firestore.Desc).Where("user", "==", userRef).Limit(35) // Suppose that the user posts 5 notes per day max by 7 days

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
