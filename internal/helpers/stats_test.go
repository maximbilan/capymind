package helpers

import (
	"context"
	"testing"

	"github.com/capymind/internal/database"
	"github.com/capymind/internal/translator"
)

type adminStatsMock struct{}

func (storage adminStatsMock) GetTotalUserCount(ctx *context.Context) (int64, error) {
	return 100, nil
}

func (storage adminStatsMock) GetActiveUserCount(ctx *context.Context) (int64, error) {
	return 75, nil
}

func (storage adminStatsMock) GetTotalNoteCount(ctx *context.Context) (int64, error) {
	return 999, nil
}

type feedbackStatsMock struct{}

func (storage feedbackStatsMock) GetFeedbackForLastWeek(ctx *context.Context) ([]database.UserFeedback, error) {
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

func (storage feedbackStatsMock) NewFeedback(ctx *context.Context, user database.User, feedback database.Feedback) error {
	return nil
}

func TestGetStats(t *testing.T) {
	adminStorage := adminStatsMock{}
	feedbackStorage := feedbackStatsMock{}

	context := context.Background()
	stats := GetStats(&context, translator.EN, adminStorage, feedbackStorage)

	if len(stats) != 14 {
		t.Error("Expected 14 stats, got", len(stats))
	}

	if stats[0] != "The total number of users is 100" {
		t.Error("Expected The total number of users is 100, got", stats[0])
	}

	if stats[1] != "The total number of active users is 75" {
		t.Error("Expected The total number of active users is 75, got", stats[1])
	}

	if stats[2] != "The total number of notes is 999" {
		t.Error("Expected The total number of notes is 999, got", stats[2])
	}

	if stats[9] != "\nTest feedback\n" {
		t.Error("Expected Test feedback, got", stats[9])
	}

	if stats[13] != "\nTest feedback 2\n" {
		t.Error("Expected Test feedback 2, got", stats[9])
	}
}
