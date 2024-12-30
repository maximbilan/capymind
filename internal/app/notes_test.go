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

func TestHandleMissingNoteHandler(t *testing.T) {
	user := &database.User{
		IsTyping: false,
	}
	job := createJob("/missing_note potential message", user)
	session := createSession(job, user, nil)
	noteStorage := mocks.NoteStorageMock{}

	handleMissingNote(session, noteStorage)

	if session.Job.Output[0].TextID != "finish_note" {
		t.Error("Expected 'finish_note', got", session.Job.Output[0].TextID)
	}
	if session.User.IsTyping != false {
		t.Error("Expected 'false', got", true)
	}
}

func TestHandleInputWithoutCommandHandler(t *testing.T) {
	user := &database.User{
		IsTyping: false,
	}
	job := createJob("potential message", user)
	session := createSession(job, user, nil)

	handleInputWithoutCommand(session)

	if session.Job.Output[0].TextID != "missing_note" {
		t.Error("Expected 'missing_note', got", session.Job.Output[0].TextID)
	}
	if session.Job.Output[0].Buttons[0].TextID != "missing_note_confirm" {
		t.Error("Expected 'missing_note_confirm', got", session.Job.Output[0].Buttons[0].TextID)
	}
}

func TestHandleLastNoteHandler(t *testing.T) {
	user := &database.User{
		IsTyping: false,
	}
	job := createJob("/last_note", user)
	session := createSession(job, user, nil)
	noteStorage := mocks.NoteStorageMock{}

	handleLastNote(session, noteStorage)

	if session.Job.Output[0].TextID != "Here’s your most recent note 👇\n\nTest note ... dream" {
		t.Error("Expected 'your_last_note', got", session.Job.Output[0].TextID)
	}
	if session.Job.Output[1].TextID != "do_you_want_sleep_analysis" {
		t.Error("Expected 'do_you_want_sleep_analysis', got", session.Job.Output[1].TextID)
	}
	if session.Job.Output[1].Buttons[0].TextID != "sleep_analysis" {
		t.Error("Expected 'sleep_analysis', got", session.Job.Output[1].Buttons[0].TextID)
	}
}

func TestHandleLastEmptyNoteHandler(t *testing.T) {
	user := &database.User{
		IsTyping: false,
	}
	job := createJob("/last_note", user)
	session := createSession(job, user, nil)
	noteStorage := mocks.EmptyNoteStorageMock{}

	handleLastNote(session, noteStorage)

	if session.Job.Output[0].TextID != "no_notes" {
		t.Error("Expected 'no_notes', got", session.Job.Output[0].TextID)
	}
}

func TestNoteCountHandler(t *testing.T) {
	user := &database.User{}
	job := createJob("/note_count", user)
	session := createSession(job, user, nil)
	noteStorage := mocks.NoteStorageMock{}

	handleNoteCount(session, noteStorage)

	if session.Job.Output[0].TextID != "You have made a total of 10 entries in your journal.\nKeep up the great work! 🚀" {
		t.Error("Expected 'You have made a total of 10 entries in your journal.\nKeep up the great work! 🚀', got", session.Job.Output[0].TextID)
	}
}
