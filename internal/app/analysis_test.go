package app

import (
	"testing"

	"github.com/capymind/internal/database"
	"github.com/capymind/internal/mocks"
)

func TestAnalysisHandler(t *testing.T) {
	session := createSession(&Job{Command: "/analysis"}, &database.User{}, nil)
	noteStorage := mocks.NoteStorageMock{}
	aiService := mocks.ValidAIServiceMock{}

	handleAnalysis(session, noteStorage, aiService)

	if session.Job.Output[0].TextID != "valid response" {
		t.Errorf("Expected valid response, got %s", session.Job.Output[0].TextID)
	}
}
