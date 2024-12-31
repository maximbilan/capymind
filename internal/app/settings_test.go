package app

import (
	"context"
	"testing"

	"github.com/capymind/internal/database"
	"github.com/capymind/internal/mocks"
)

func TestSettingsHandler(t *testing.T) {
	session := createSession(&Job{Command: "/settings"}, &database.User{}, nil, nil)
	handleSettings(session)

	if session.Job.Output[0].TextID != "settings_descr" {
		t.Errorf("Expected output text to be 'settings_descr', got %s", session.Job.Output[0].TextID)
	}
	if len(session.Job.Output[0].Buttons) != 4 {
		t.Errorf("Expected 4 buttons, got %d", len(session.Job.Output[0].Buttons))
	}
	if session.Job.Output[0].Buttons[0].TextID != "language" {
		t.Errorf("Expected button text to be 'language', got %s", session.Job.Output[0].Buttons[0].TextID)
	}
	if session.Job.Output[0].Buttons[1].TextID != "reminders_button" {
		t.Errorf("Expected button text to be 'reminders_button', got %s", session.Job.Output[0].Buttons[1].TextID)
	}
	if session.Job.Output[0].Buttons[2].TextID != "download_all_notes" {
		t.Errorf("Expected button text to be 'download_all_notes', got %s", session.Job.Output[0].Buttons[2].TextID)
	}
	if session.Job.Output[0].Buttons[3].TextID != "delete_account" {
		t.Errorf("Expected button text to be 'delete_account', got %s", session.Job.Output[0].Buttons[3].TextID)
	}
}

func TestGetSettings(t *testing.T) {
	ctx := context.Background()
	userId := "123"
	settingsStorage := mocks.SettingsStorageMock{}

	settings := getSettings(&ctx, userId, settingsStorage)

	if settings == nil {
		t.Error("Expected settings, got nil")
	}
	if *settings.SecondsFromUTC != 7200 {
		t.Errorf("Expected 7200, got %d", settings.SecondsFromUTC)
	}
	if *settings.MorningReminderOffset != 7200 {
		t.Errorf("Expected 7200, got %d", *settings.MorningReminderOffset)
	}
	if *settings.EveningReminderOffset != 7200 {
		t.Errorf("Expected 7200, got %d", *settings.EveningReminderOffset)
	}
	if *settings.HasMorningReminder != true {
		t.Errorf("Expected true, got %t", *settings.HasMorningReminder)
	}
	if *settings.HasEveningReminder != true {
		t.Errorf("Expected true, got %t", *settings.HasEveningReminder)
	}
}
