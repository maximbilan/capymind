package app

import (
	"testing"

	"github.com/capymind/internal/database"
	"github.com/capymind/internal/mocks"
)

func TestTotalUserCountHandler(t *testing.T) {
	session := createSession(&Job{Command: "/total_user_count"}, &database.User{}, nil, nil)
	adminStorage := mocks.AdminStorageMock{}

	handleTotalUserCount(session, adminStorage)

	if session.Job.Output[0].TextID != "The total number of users is 100" {
		t.Errorf("Expected The total number of users is 100, got %s", session.Job.Output[0].TextID)
	}
}

func TestTotalActiveUserCountHandler(t *testing.T) {
	session := createSession(&Job{Command: "/total_active_user_count"}, &database.User{}, nil, nil)
	adminStorage := mocks.AdminStorageMock{}

	handleTotalActiveUserCount(session, adminStorage)

	if session.Job.Output[0].TextID != "The total number of active users is 75" {
		t.Errorf("Expected The total number of active users is 75, got %s", session.Job.Output[0].TextID)
	}
}

func TestTotalNoteCountHandler(t *testing.T) {
	session := createSession(&Job{Command: "/total_note_count"}, &database.User{}, nil, nil)
	adminStorage := mocks.AdminStorageMock{}

	handleTotalNoteCount(session, adminStorage)

	if session.Job.Output[0].TextID != "The total number of notes is 999" {
		t.Errorf("Expected The total number of notes is 999, got %s", session.Job.Output[0].TextID)
	}
}
