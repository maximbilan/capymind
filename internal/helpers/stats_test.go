package helpers

import (
	"context"
	"testing"

	"github.com/capymind/internal/mocks"
	"github.com/capymind/internal/translator"
)

func TestGetStats(t *testing.T) {
	adminStorage := mocks.AdminStorageMock{}
	feedbackStorage := mocks.FeedbackStorageMock{}

	context := context.Background()
	stats := GetStats(&context, translator.EN, adminStorage, feedbackStorage)

	if len(stats) != 14 {
		t.Error("Expected 14 stats, got", len(stats))
	}

	if stats[0] != "The total number of users is 100" {
		t.Error("Expected The total number of users is 100, got", stats[0])
	}

	if stats[1] != "The total number of active users is 75" {
		t.Error("Expected The total number of active users is 75, got", stats[1])
	}

	if stats[2] != "The total number of notes is 999" {
		t.Error("Expected The total number of notes is 999, got", stats[2])
	}

	if stats[9] != "\nTest feedback\n" {
		t.Error("Expected Test feedback, got", stats[9])
	}

	if stats[13] != "\nTest feedback 2\n" {
		t.Error("Expected Test feedback 2, got", stats[9])
	}
}
