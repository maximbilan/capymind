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

func NewFeedback(ctx *context.Context, user User, feedback Feedback) error {
	userRef := client.Collection(users.String()).Doc(user.ID)
	_, _, err := client.Collection(feedbacks.String()).Add(*ctx, map[string]interface{}{
		"text":      feedback.Text,
		"timestamp": feedback.Timestamp,
		"user":      userRef,
	})
	return err
}

func GetFeedbackForLastWeek(ctx *context.Context, userID string) ([]Feedback, error) {
	userRef := client.Collection(users.String()).Doc(userID)
	query := client.Collection(feedbacks.String()).OrderBy("timestamp", firestore.Desc).Where("user", "==", userRef).Where("timestamp", ">=", time.Now().AddDate(0, 0, -7))
	return getFeedbackForQuery(ctx, query)
}

func getFeedbackForQuery(ctx *context.Context, query firestore.Query) ([]Feedback, error) {
	docs, err := query.Documents(*ctx).GetAll()
	if err != nil {
		return nil, err
	}

	var feedbacks []Feedback
	for _, doc := range docs {
		var feedback Feedback
		doc.DataTo(&feedback)
		feedbacks = append(feedbacks, feedback)
	}

	return feedbacks, nil
}
