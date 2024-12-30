//coverage:ignore file

package mocks

import (
	"context"

	"github.com/capymind/internal/database"
)

type FeedbackStorageMock struct{}

func (storage FeedbackStorageMock) GetFeedbackForLastWeek(ctx *context.Context) ([]database.UserFeedback, error) {
	var array []database.UserFeedback

	firstName := "John"
	lastName := "Doe"
	user := database.User{
		ID:        "1",
		FirstName: &firstName,
		LastName:  &lastName,
	}
	feedback1 := database.Feedback{
		Text: "Test feedback",
	}
	feedback2 := database.Feedback{
		Text: "Test feedback 2",
	}

	userFeedback1 := database.UserFeedback{
		User:     user,
		Feedback: feedback1,
	}
	userFeedback2 := database.UserFeedback{
		User:     user,
		Feedback: feedback2,
	}

	array = append(array, userFeedback1)
	array = append(array, userFeedback2)

	return array, nil
}

func (storage FeedbackStorageMock) NewFeedback(ctx *context.Context, user database.User, feedback database.Feedback) error {
	return nil
}
