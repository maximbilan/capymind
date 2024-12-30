package app

import (
	"testing"

	"github.com/capymind/internal/database"
	"github.com/capymind/internal/mocks"
)

func TestFeedbackHandler(t *testing.T) {
	session := createSession(&Job{Command: "/feedback"}, &database.User{}, nil)
	feedbackStorage := mocks.FeedbackStorageMock{}

	handleFeedbackLastWeek(session, feedbackStorage)

	if len(session.Job.Output) != 5 {
		t.Errorf("Expected 5 feedback items, got %d", len(session.Job.Output))
	}
	if session.Job.Output[3].TextID != "John Doe:\nTest feedback\n" {
		t.Errorf("Expected Test feedback, got %s", session.Job.Output[3].TextID)
	}
	if session.Job.Output[4].TextID != "John Doe:\nTest feedback 2\n" {
		t.Errorf("Expected Test feedback 2, got %s", session.Job.Output[4].TextID)
	}
}
