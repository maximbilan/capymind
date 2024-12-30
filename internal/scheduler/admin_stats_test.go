package scheduler

import (
	"context"
	"testing"

	"github.com/capymind/internal/mocks"
	"github.com/capymind/internal/translator"
)

func TestAdminStats(t *testing.T) {
	context := context.Background()
	locale := translator.EN
	adminStorage := mocks.AdminStorageMock{}
	feedbackStorage := mocks.FeedbackStorageMock{}

	response := prepareAdminStats(&context, locale, adminStorage, feedbackStorage)

	if *response != "The total number of users is 100\nThe total number of active users is 75\nThe total number of notes is 999\n\nFeedback from last week 📈\n\nJohn Doe:\nTest feedback\n\nJohn Doe:\nTest feedback 2\n\n" {
		t.Errorf("Expected valid response, got %s", *response)
	}
}
