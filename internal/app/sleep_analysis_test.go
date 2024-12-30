package app

import (
	"testing"

	"github.com/capymind/internal/database"
	"github.com/capymind/internal/mocks"
	"github.com/capymind/internal/translator"
)

func TestSleepAnalysisHandler(t *testing.T) {
	session := createSession(&Job{Command: "/sleep_analysis"}, &database.User{}, nil)
	noteStorage := mocks.NoteStorageMock{}
	aiService := mocks.ValidAIServiceMock{}

	handleSleepAnalysis(session, noteStorage, aiService)

	if session.Job.Output[0].TextID != "valid response" {
		t.Error("Expected 'valid response', got", session.Job.Output[0].TextID)
	}
}

func TestSleepAnalysisNoNotesHandler(t *testing.T) {
	session := createSession(&Job{Command: "/sleep_analysis"}, &database.User{}, nil)
	noteStorage := mocks.EmptyNoteStorageMock{}
	aiService := mocks.ValidAIServiceMock{}

	handleSleepAnalysis(session, noteStorage, aiService)

	if session.Job.Output[0].TextID != "no_notes" {
		t.Error("Expected 'no_notes', got", session.Job.Output[0].TextID)
	}
}

func TestNoSleepAnalysisHandler(t *testing.T) {
	session := createSession(&Job{Command: "/sleep_analysis"}, &database.User{}, nil)
	noteStorage := mocks.NoteStorageMock{}
	aiService := mocks.InvalidAIServiceMock{}

	handleSleepAnalysis(session, noteStorage, aiService)

	if session.Job.Output[0].TextID != "no_notes" {
		t.Error("Expected 'no_notes', got", session.Job.Output[0].TextID)
	}
}

func TestAskForSleepAnalysisHandler(t *testing.T) {
	session := createSession(&Job{Command: "/sleep_analysis"}, &database.User{}, nil)
	askForSleepAnalysis(session)

	if session.Job.Output[0].TextID != "do_you_want_sleep_analysis" {
		t.Error("Expected 'do_you_want_sleep_analysis', got", session.Job.Output[0].TextID)
	}
	if session.Job.Output[0].Buttons[0].TextID != "sleep_analysis" {
		t.Error("Expected 'sleep_analysis', got", session.Job.Output[0].Buttons[0].TextID)
	}
}

func TestCheckIfNoteADream(t *testing.T) {
	text1 := "I had a dream last night"
	locale1 := translator.Locale("en")
	if !checkIfNoteADream(text1, locale1) {
		t.Fatalf("Expected true, got false")
	}

	text2 := "Снилось щось цікаве"
	locale2 := translator.Locale("uk")
	if !checkIfNoteADream(text2, locale2) {
		t.Fatalf("Expected true, got false")
	}
}

func TestHandleWeeklyAnalysisHandler(t *testing.T) {
	user := &database.User{
		ID: "test",
	}
	job := createJob("/weekly_analysis", user)
	session := createSession(job, user, nil)
	noteStorage := mocks.NoteStorageMock{}
	aiService := mocks.ValidAIServiceMock{}

	handleWeeklyAnalysis(session, noteStorage, aiService)

	if session.Job.Output[0].TextID != "Weekly analysis 🧑‍⚕️\n\nvalid response" {
		t.Error("Expected 'analysis_waiting', got", session.Job.Output[0].TextID)
	}
}

func TestHandleWeeklyAnalysisEmptyNotesHandler(t *testing.T) {
	user := &database.User{
		ID: "test",
	}
	job := createJob("/weekly_analysis", user)
	session := createSession(job, user, nil)
	noteStorage := mocks.EmptyNoteStorageMock{}
	aiService := mocks.ValidAIServiceMock{}

	handleWeeklyAnalysis(session, noteStorage, aiService)

	if session.Job.Output[0].TextID != "no_notes" {
		t.Error("Expected 'no_notes', got", session.Job.Output[0].TextID)
	}
}

func TestHandleWeeklyAnalysisEmptyAnalysisHandler(t *testing.T) {
	user := &database.User{
		ID: "test",
	}
	job := createJob("/weekly_analysis", user)
	session := createSession(job, user, nil)
	noteStorage := mocks.NoteStorageMock{}
	aiService := mocks.InvalidAIServiceMock{}

	handleWeeklyAnalysis(session, noteStorage, aiService)

	if session.Job.Output[0].TextID != "no_notes" {
		t.Error("Expected 'no_notes', got", session.Job.Output[0].TextID)
	}
}
