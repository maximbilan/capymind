//coverage:ignore file

package database

import (
	"context"
	"time"
)

type Feedback struct {
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

type UserFeedback struct {
	User     User
	Feedback Feedback
}

type FeedbackStorage interface {
	NewFeedback(ctx *context.Context, user User, feedback Feedback) error
	GetFeedbackForLastWeek(ctx *context.Context) ([]UserFeedback, error)
}
