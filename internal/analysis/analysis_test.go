package analysis

import (
	"context"
	"testing"

	"github.com/capymind/internal/translator"
)

type validServiceMock struct{}

func (service validServiceMock) Request(name string, description string, systemPrompt string, userPrompt string, ctx *context.Context) *string {
	response := "valid response"
	return &response
}

type invalidServiceMock struct{}

func (service invalidServiceMock) Request(name string, description string, systemPrompt string, userPrompt string, ctx *context.Context) *string {
	return nil
}

func TestValidAnalysis(t *testing.T) {
	service := validServiceMock{}
	notes := []string{"note1", "note2", "note3"}
	ctx := context.Background()
	response := AnalyzeQuickly(service, notes, translator.EN, &ctx)
	if *response != "valid response" {
		t.Error("Expected valid response, got nil")
	}
}

func TestInvalidAnalysis(t *testing.T) {
	service := invalidServiceMock{}
	notes := []string{"note1", "note2", "note3"}
	ctx := context.Background()
	response := AnalyzeQuickly(service, notes, translator.EN, &ctx)
	if response != nil {
		t.Error("Expected nil response, got valid response")
	}
}

func TestAnalysisWithHeader(t *testing.T) {
	service := validServiceMock{}
	notes := []string{"note1", "note2", "note3"}
	ctx := context.Background()
	response := AnalyzeLastWeek(service, notes, translator.EN, &ctx)
	if *response != "Weekly analysis 🧑‍⚕️\n\nvalid response" {
		t.Error("Expected Weekly analysis 🧑‍⚕️\n\nvalid response, got nil")
	}
}

func TestAnalyzeSleep(t *testing.T) {
	service := validServiceMock{}
	text := "I slept well last night"
	ctx := context.Background()
	response := AnalyzeSleep(service, text, translator.EN, &ctx)
	if *response != "valid response" {
		t.Error("Expected valid response, got nil")
	}
}
