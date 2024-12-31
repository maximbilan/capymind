package app

import (
	"testing"

	"github.com/capymind/internal/database"
	"github.com/capymind/internal/mocks"
)

func TestStatsHandler(t *testing.T) {
	session := createSession(&Job{Command: "/stats"}, &database.User{}, nil, nil)
	adminStorage := mocks.AdminStorageMock{}
	feedbackStorage := mocks.FeedbackStorageMock{}

	handleStats(session, adminStorage, feedbackStorage)

	if session.Job.Output[0].TextID != "The total number of users is 100\nThe total number of active users is 75\nThe total number of notes is 999\n\nFeedback from last week 📈\n\nJohn Doe:\nTest feedback\n\nJohn Doe:\nTest feedback 2\n\n" {
		t.Error("Wrong result, got", session.Job.Output[0].TextID)
	}
}
