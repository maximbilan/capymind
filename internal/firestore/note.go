package firestore

import (
	"context"
	"errors"
	"time"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/firestore/apiv1/firestorepb"
)

type Note struct {
	Text      string                 `firestore:"text"`
	Timestamp time.Time              `firestore:"timestamp"`
	User      *firestore.DocumentRef `firestore:"user"`
}

// Create a new note
func NewNote(ctx *context.Context, user User, note Note) error {
	userRef := client.Collection(users.String()).Doc(user.ID)
	_, _, err := client.Collection(notes.String()).Add(*ctx, map[string]interface{}{
		"text":      note.Text,
		"timestamp": note.Timestamp,
		"user":      userRef,
	})
	return err
}

// Get the last note
func LastNote(ctx *context.Context, userID string) (*Note, error) {
	userRef := client.Collection(users.String()).Doc(userID)
	query := client.Collection(notes.String()).OrderBy("timestamp", firestore.Desc).Where("user", "==", userRef).Limit(1)

	docs, err := query.Documents(*ctx).GetAll()
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

// Get notes for the last 7 days
func GetNotesForLastWeek(ctx *context.Context, userID string) ([]Note, error) {
	userRef := client.Collection(users.String()).Doc(userID)
	query := client.Collection(notes.String()).OrderBy("timestamp", firestore.Desc).Where("user", "==", userRef).Where("timestamp", ">=", time.Now().AddDate(0, 0, -7))
	return getNotesForQuery(ctx, query)
}

// Get the user's notes (limited by count)
func GetNotes(ctx *context.Context, userID string, count int) ([]Note, error) {
	userRef := client.Collection(users.String()).Doc(userID)
	query := client.Collection(notes.String()).OrderBy("timestamp", firestore.Desc).Where("user", "==", userRef).Limit(count)
	return getNotesForQuery(ctx, query)
}

func getNotesForQuery(ctx *context.Context, query firestore.Query) ([]Note, error) {
	docs, err := query.Documents(*ctx).GetAll()
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

// Get notes count for the current user (aggregation query)
func NotesCount(ctx *context.Context, userID string) (int64, error) {
	userRef := client.Collection(users.String()).Doc(userID)
	query := client.Collection(notes.String()).Where("user", "==", userRef)
	aggregationQuery := query.NewAggregationQuery().WithCount("all")

	results, err := aggregationQuery.Get(*ctx)
	if err != nil {
		return 0, err
	}

	count, ok := results["all"]
	if !ok {
		return 0, errors.New("[Firestore]: couldn't get alias for COUNT from results")
	}

	countValue := count.(*firestorepb.Value)
	return countValue.GetIntegerValue(), nil
}
