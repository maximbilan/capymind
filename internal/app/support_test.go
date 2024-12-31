package app

import (
	"testing"

	"github.com/capymind/internal/database"
	"github.com/capymind/internal/mocks"
)

func TestSupportHandler(t *testing.T) {
	user := &database.User{
		IsTyping: false,
	}
	job := createJob("/support", user)
	session := createSession(job, user, nil, nil)

	startFeedback(session)

	if session.User.IsTyping != true {
		t.Error("Expected 'true', got", false)
	}
}

func TestFinishFeedbackHandler(t *testing.T) {
	user := &database.User{
		IsTyping: true,
	}
	job := createJob("/support", user)
	session := createSession(job, user, nil, nil)
	feedbackStorage := mocks.FeedbackStorageMock{}

	finishFeedback(session, feedbackStorage)

	if session.Job.Output[0].TextID != "finish_feedback" {
		t.Error("Expected 'finish_feedback', got", session.Job.Output[0].TextID)
	}
	if session.User.IsTyping != false {
		t.Error("Expected 'false', got", true)
	}
}
