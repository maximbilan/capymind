package analysis

import (
	"context"
	"testing"

	"github.com/capymind/internal/mocks"
	"github.com/capymind/internal/translator"
)

func TestValidAnalysis(t *testing.T) {
	service := mocks.ValidAIServiceMock{}
	notes := []string{"note1", "note2", "note3"}
	ctx := context.Background()
	response := AnalyzeQuickly(service, notes, translator.EN, &ctx)
	if *response != "valid response" {
		t.Error("Expected valid response, got nil")
	}
}

func TestInvalidAnalysis(t *testing.T) {
	service := mocks.InvalidAIServiceMock{}
	notes := []string{"note1", "note2", "note3"}
	ctx := context.Background()
	response := AnalyzeQuickly(service, notes, translator.EN, &ctx)
	if response != nil {
		t.Error("Expected nil response, got valid response")
	}
}

func TestAnalysisWithHeader(t *testing.T) {
	service := mocks.ValidAIServiceMock{}
	notes := []string{"note1", "note2", "note3"}
	ctx := context.Background()
	response := AnalyzeLastWeek(service, notes, translator.EN, &ctx)
	if *response != "Weekly analysis 🧑‍⚕️\n\nvalid response" {
		t.Error("Expected Weekly analysis 🧑‍⚕️\n\nvalid response, got nil")
	}
}

func TestAnalyzeSleep(t *testing.T) {
	service := mocks.ValidAIServiceMock{}
	text := "I slept well last night"
	ctx := context.Background()
	response := AnalyzeSleep(service, text, translator.EN, &ctx)
	if *response != "valid response" {
		t.Error("Expected valid response, got nil")
	}
}
