package scheduler

import (
	"context"
	"testing"

	"github.com/capymind/internal/database"
	"github.com/capymind/internal/mocks"
	"github.com/capymind/internal/translator"
)

func TestWeeklyAnalysis(t *testing.T) {
	user := database.User{ID: "123"}
	context := context.Background()
	locale := translator.EN
	noteStorage := mocks.NoteStorageMock{}
	aiService := mocks.ValidAIServiceMock{}

	response := prepareWeeklyAnalysis(&user, &context, locale, noteStorage, aiService)
	if *response != "Weekly analysis 🧑\u200d⚕️\n\nvalid response" {
		t.Errorf("Expected valid response, got %s", *response)
	}
}

func TestWeeklyAnalysisNoNotes(t *testing.T) {
	user := database.User{ID: "123"}
	context := context.Background()
	locale := translator.EN
	noteStorage := mocks.EmptyNoteStorageMock{}
	aiService := mocks.ValidAIServiceMock{}

	response := prepareWeeklyAnalysis(&user, &context, locale, noteStorage, aiService)
	if response != nil {
		t.Errorf("Expected nil, got %s", *response)
	}
}

func TestWeeklyAnalysisNoAIService(t *testing.T) {
	user := database.User{ID: "123"}
	context := context.Background()
	locale := translator.EN
	noteStorage := mocks.NoteStorageMock{}
	aiService := mocks.InvalidAIServiceMock{}

	response := prepareWeeklyAnalysis(&user, &context, locale, noteStorage, aiService)
	if response != nil {
		t.Errorf("Expected nil, got %s", *response)
	}
}
