package app

import (
	"testing"

	"github.com/capymind/internal/database"
)

func TestSettingsHandler(t *testing.T) {
	session := createSession(&Job{Command: "/settings"}, &database.User{}, nil)
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
	if session.Job.Output[0].Buttons[1].TextID != "timezone" {
		t.Errorf("Expected button text to be 'timezone', got %s", session.Job.Output[0].Buttons[1].TextID)
	}
	if session.Job.Output[0].Buttons[2].TextID != "download_all_notes" {
		t.Errorf("Expected button text to be 'download_all_notes', got %s", session.Job.Output[0].Buttons[2].TextID)
	}
	if session.Job.Output[0].Buttons[3].TextID != "delete_account" {
		t.Errorf("Expected button text to be 'delete_account', got %s", session.Job.Output[0].Buttons[3].TextID)
	}
}
