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

	if len(session.Job.Output) != 11 {
		t.Errorf("Expected 11 feedback items, got %d", len(session.Job.Output))
	}
	if session.Job.Output[6].TextID != "\nTest feedback\n" {
		t.Errorf("Expected Test feedback, got %s", session.Job.Output[0].TextID)
	}
	if session.Job.Output[10].TextID != "\nTest feedback 2\n" {
		t.Errorf("Expected Test feedback 2, got %s", session.Job.Output[0].TextID)
	}
}
