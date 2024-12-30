package scheduler

import (
	"context"
	"testing"

	"github.com/capymind/internal/database"
	"github.com/capymind/internal/mocks"
	"github.com/capymind/internal/translator"
)

func TestUserStats(t *testing.T) {
	user := database.User{ID: "123"}
	context := context.Background()
	locale := translator.EN
	noteStorage := mocks.NoteStorageMock{}

	response := prepareUserStats(&user, &context, locale, noteStorage)

	if *response != "You have made a total of 10 entries in your journal.\nKeep up the great work! 🚀" {
		t.Error("Expected valid response, got nil")
	}
}
