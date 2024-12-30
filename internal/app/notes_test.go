package app

import (
	"testing"

	"github.com/capymind/internal/database"
	"github.com/capymind/internal/mocks"
)

func TestStartNoteHandler(t *testing.T) {
	session := createSession(&Job{Command: "/note"}, &database.User{
		IsTyping: false,
	}, nil)
	startNote(session)

	if session.User.IsTyping != true {
		t.Error("Expected 'true', got", false)
	}
}

func TestFinishNoteHandler(t *testing.T) {
	session := createSession(&Job{Command: "/note"}, &database.User{
		IsTyping: true,
	}, nil)
	noteStorage := mocks.NoteStorageMock{}

	finishNote("dream", session, noteStorage)

	if session.Job.Output[0].TextID != "finish_note" {
		t.Error("Expected 'finish_note', got", session.Job.Output[0].TextID)
	}
	if session.Job.Output[1].TextID != "do_you_want_sleep_analysis" {
		t.Error("Expected 'do_you_want_sleep_analysis', got", session.Job.Output[1].TextID)
	}
	if session.Job.Output[1].Buttons[0].TextID != "sleep_analysis" {
		t.Error("Expected 'sleep_analysis', got", session.Job.Output[1].Buttons[0].TextID)
	}
	if session.User.IsTyping != false {
		t.Error("Expected 'false', got", true)
	}
}
