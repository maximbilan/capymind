package app

import (
	"testing"

	"github.com/capymind/internal/botservice"
)

func TestCreateUserFromMessage(t *testing.T) {
	message := botservice.BotMessage{
		UserID:       "456",
		ChatID:       456,
		UserName:     "test",
		FirstName:    "Test",
		LastName:     "User",
		LanguageCode: "en",
		Text:         "/language en",
	}

	user := createUser(message)
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
	message := botservice.BotMessage{
		UserID:       "456",
		ChatID:       456,
		UserName:     "test",
		FirstName:    "Test",
		LastName:     "User",
		LanguageCode: "en",
		Text:         "/timezone 25200",
	}

	user := createUser(message)
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
