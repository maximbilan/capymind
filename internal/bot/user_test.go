package bot

import (
	"testing"

	"github.com/capymind/internal/telegram"
)

func TestCreateUserFromMessage(t *testing.T) {
	update := telegram.Update{
		ID: 789,
		Message: &telegram.Message{
			ID:   101,
			Text: "/language en",
			Chat: &telegram.Chat{
				ID: 456,
			},
			From: &telegram.User{
				ID:           456,
				UserName:     "test",
				FirstName:    "Test",
				LastName:     "User",
				LanguageCode: "en",
			},
		},
	}

	user := createUser(update)
	if user == nil {
		t.Fatalf("User is nil")
	}
	if user.ID != "456" {
		t.Fatalf("ID is not 456")
	}
	if user.ChatID != 456 {
		t.Fatalf("ChatID is not 456")
	}
	if *user.UserName != "test" {
		t.Fatalf("UserName is not test")
	}
	if *user.FirstName != "Test" {
		t.Fatalf("FirstName is not Test")
	}
	if *user.LastName != "User" {
		t.Fatalf("LastName is not User")
	}
	if *user.Locale != "en" {
		t.Fatalf("Locale is not en")
	}
}

func TestUserFromCallback(t *testing.T) {
	update := telegram.Update{
		ID: 789,
		CallbackQuery: &telegram.CallbackQuery{
			ID:   "123",
			Data: "/timezone 25200",
			From: &telegram.User{
				ID:           456,
				UserName:     "test",
				FirstName:    "Test",
				LastName:     "User",
				LanguageCode: "en",
			},
			Message: &telegram.Message{
				ID: 101,
				Chat: &telegram.Chat{
					ID: 456,
				},
			},
		},
	}

	user := createUser(update)
	if user == nil {
		t.Fatalf("User is nil")
	}
	if user.ID != "456" {
		t.Fatalf("ID is not 456")
	}
	if user.ChatID != 456 {
		t.Fatalf("ChatID is not 456")
	}
	if *user.UserName != "test" {
		t.Fatalf("UserName is not test")
	}
	if *user.FirstName != "Test" {
		t.Fatalf("FirstName is not Test")
	}
	if *user.LastName != "User" {
		t.Fatalf("LastName is not User")
	}
	if *user.Locale != "en" {
		t.Fatalf("Locale is not en")
	}
}
