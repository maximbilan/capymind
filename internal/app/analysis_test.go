package app

import (
	"testing"

	"github.com/capymind/internal/database"
	"github.com/capymind/internal/mocks"
)

func TestAnalysisHandler(t *testing.T) {
	session := createSession(&Job{Command: "/analysis"}, &database.User{}, nil, nil)
	noteStorage := mocks.NoteStorageMock{}
	aiService := mocks.ValidAIServiceMock{}

	handleAnalysis(session, noteStorage, aiService)

	if session.Job.Output[0].TextID != "valid response" {
		t.Errorf("Expected valid response, got %s", session.Job.Output[0].TextID)
	}
}

func TestAnalysisHandlerNoNotes(t *testing.T) {
	session := createSession(&Job{Command: "/analysis"}, &database.User{}, nil, nil)
	noteStorage := mocks.EmptyNoteStorageMock{}
	aiService := mocks.ValidAIServiceMock{}

	handleAnalysis(session, noteStorage, aiService)

	if session.Job.Output[0].TextID != "no_analysis" {
		t.Errorf("Expected no_analysis, got %s", session.Job.Output[0].TextID)
	}
}

func TestAnalysisHandlerNoAIService(t *testing.T) {
	session := createSession(&Job{Command: "/analysis"}, &database.User{}, nil, nil)
	noteStorage := mocks.NoteStorageMock{}
	aiService := mocks.InvalidAIServiceMock{}

	handleAnalysis(session, noteStorage, aiService)

	if session.Job.Output[0].TextID != "no_analysis" {
		t.Errorf("Expected no_analysis, got %s", session.Job.Output[0].TextID)
	}
}
