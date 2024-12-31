package app

import (
	"testing"

	"github.com/capymind/internal/database"
)

func TestAskForReminders(t *testing.T) {
	session := createSession(&Job{Command: "/reminders"}, &database.User{}, nil, nil)
	askForReminders(session)

	if session.Job.Output[0].TextID != "onboarding_reminders" {
		t.Error("Expected 'onboarding_reminders', got", session.Job.Output[0].TextID)
	}
	if session.Job.Output[0].Buttons[0].TextID != "reminders_enable_button" {
		t.Error("Expected 'reminders_enable_button', got", session.Job.Output[0].Buttons[0].TextID)
	}
	if session.Job.Output[0].Buttons[0].Callback != "/enable_all_reminders" {
		t.Error("Expected '/enable_all_reminders', got", session.Job.Output[0].Buttons[0].Callback)
	}
	if session.Job.Output[0].Buttons[1].TextID != "continue" {
		t.Error("Expected 'continue', got", session.Job.Output[0].Buttons[1].TextID)
	}
	if session.Job.Output[0].Buttons[1].Callback != "/skip_reminders" {
		t.Error("Expected '/skip_reminders', got", session.Job.Output[0].Buttons[1].Callback)
	}
}

func TestRemindersOnHandler(t *testing.T) {
	enabled := true
	session := createSession(&Job{Command: "/reminders"}, &database.User{}, &database.Settings{
		HasMorningReminder: &enabled,
		HasEveningReminder: &enabled,
	}, nil)
	handleReminders(session)

	if session.Job.Output[0].TextID != "reminders_button" {
		t.Error("Expected 'reminders_button', got", session.Job.Output[0].TextID)
	}
	if len(session.Job.Output[0].Buttons) != 4 {
		t.Errorf("Expected 4 buttons, got %d", len(session.Job.Output[0].Buttons))
	}
	if session.Job.Output[0].Buttons[0].TextID != "reminders_disable_button" {
		t.Error("Expected 'reminders_disable_button', got", session.Job.Output[0].Buttons[0].TextID)
	}
	if session.Job.Output[0].Buttons[0].Callback != "/disable_all_reminders" {
		t.Error("Expected '/disable_all_reminders', got", session.Job.Output[0].Buttons[0].Callback)
	}
	if session.Job.Output[0].Buttons[1].TextID != "timezone" {
		t.Error("Expected 'timezone', got", session.Job.Output[0].Buttons[1].TextID)
	}
	if session.Job.Output[0].Buttons[1].Callback != "/timezone" {
		t.Error("Expected '/timezone', got", session.Job.Output[0].Buttons[1].Callback)
	}
	if session.Job.Output[0].Buttons[2].TextID != "morning_reminder_button" {
		t.Error("Expected 'morning_reminder_button', got", session.Job.Output[0].Buttons[2].TextID)
	}
	if session.Job.Output[0].Buttons[2].Callback != "/morning_reminder" {
		t.Error("Expected '/morning_reminder', got", session.Job.Output[0].Buttons[2].Callback)
	}
	if session.Job.Output[0].Buttons[3].TextID != "evening_reminder_button" {
		t.Error("Expected 'evening_reminder_button', got", session.Job.Output[0].Buttons[3].TextID)
	}
	if session.Job.Output[0].Buttons[3].Callback != "/evening_reminder" {
		t.Error("Expected '/evening_reminder', got", session.Job.Output[0].Buttons[3].Callback)
	}
}

func TestRemindersOffHandler(t *testing.T) {
	enabled := false
	session := createSession(&Job{Command: "/reminders"}, &database.User{}, &database.Settings{
		HasMorningReminder: &enabled,
		HasEveningReminder: &enabled,
	}, nil)
	handleReminders(session)

	if session.Job.Output[0].TextID != "reminders_button" {
		t.Error("Expected 'reminders_button', got", session.Job.Output[0].TextID)
	}
	if len(session.Job.Output[0].Buttons) != 4 {
		t.Errorf("Expected 4 buttons, got %d", len(session.Job.Output[0].Buttons))
	}
	if session.Job.Output[0].Buttons[0].TextID != "reminders_enable_button" {
		t.Error("Expected 'reminders_enable_button', got", session.Job.Output[0].Buttons[0].TextID)
	}
	if session.Job.Output[0].Buttons[0].Callback != "/enable_all_reminders" {
		t.Error("Expected '/enable_all_reminders', got", session.Job.Output[0].Buttons[0].Callback)
	}
}
