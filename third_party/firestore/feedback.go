package firestore

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/capymind/internal/database"
)

type FeedbackStorage struct{}

type FRFeedback struct {
	Text      string                 `firestore:"text"`
	Timestamp time.Time              `firestore:"timestamp"`
	User      *firestore.DocumentRef `firestore:"user"`
}

func (storage FeedbackStorage) NewFeedback(ctx *context.Context, user database.User, feedback database.Feedback) error {
	userRef := client.Collection(database.Users.String()).Doc(user.ID)
	_, _, err := client.Collection(database.Feedbacks.String()).Add(*ctx, map[string]interface{}{
		"text":      feedback.Text,
		"timestamp": feedback.Timestamp,
		"user":      userRef,
	})
	return err
}

func (storage FeedbackStorage) GetFeedbackForLastWeek(ctx *context.Context) ([]database.UserFeedback, error) {
	query := client.Collection(database.Feedbacks.String()).OrderBy("timestamp", firestore.Desc).Where("timestamp", ">=", time.Now().AddDate(0, 0, -7))
	docs, err := query.Documents(*ctx).GetAll()
	if err != nil {
		return nil, err
	}

	var userFeedbacks []database.UserFeedback
	for _, doc := range docs {
		var feedback FRFeedback
		doc.DataTo(&feedback)

		userRef := feedback.User
		userDoc, err := userRef.Get(*ctx)
		if err != nil {
			return nil, err
		}

		var user database.User
		userDoc.DataTo(&user)

		userFeedbacks = append(userFeedbacks, database.UserFeedback{
			User:     user,
			Feedback: database.Feedback{Text: feedback.Text, Timestamp: feedback.Timestamp},
		})
	}

	return userFeedbacks, nil
}
