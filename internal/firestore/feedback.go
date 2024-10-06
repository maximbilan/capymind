package firestore

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

type Feedback struct {
	Text      string                 `firestore:"text"`
	Timestamp time.Time              `firestore:"timestamp"`
	User      *firestore.DocumentRef `firestore:"user"`
}

type UserFeedback struct {
	User     User
	Feedback Feedback
}

func NewFeedback(ctx *context.Context, user User, feedback Feedback) error {
	userRef := client.Collection(users.String()).Doc(user.ID)
	_, _, err := client.Collection(feedbacks.String()).Add(*ctx, map[string]interface{}{
		"text":      feedback.Text,
		"timestamp": feedback.Timestamp,
		"user":      userRef,
	})
	return err
}

func GetFeedbackForLastWeek(ctx *context.Context) ([]UserFeedback, error) {
	query := client.Collection(feedbacks.String()).OrderBy("timestamp", firestore.Desc).Where("timestamp", ">=", time.Now().AddDate(0, 0, -7))
	docs, err := query.Documents(*ctx).GetAll()
	if err != nil {
		return nil, err
	}

	var userFeedbacks []UserFeedback
	for _, doc := range docs {
		var feedback Feedback
		doc.DataTo(&feedback)

		userRef := feedback.User
		userDoc, err := userRef.Get(*ctx)
		if err != nil {
			return nil, err
		}

		var user User
		userDoc.DataTo(&user)

		userFeedbacks = append(userFeedbacks, UserFeedback{
			User:     user,
			Feedback: feedback,
		})
	}

	return userFeedbacks, nil
}
