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

func TestMorningReminderHandler(t *testing.T) {
	enabled := false
	session := createSession(&Job{Command: "/morning_reminder"}, &database.User{}, &database.Settings{
		HasMorningReminder: &enabled,
	}, nil)
	handleMorningReminder(session)

	if session.Job.Output[0].TextID != "morning_reminder_descr" {
		t.Error("Expected 'morning_reminder_descr', got", session.Job.Output[0].TextID)
	}
	if len(session.Job.Output[0].Buttons) != 6 {
		t.Errorf("Expected 6 buttons, got %d", len(session.Job.Output[0].Buttons))
	}
	if session.Job.Output[0].Buttons[0].TextID != "reminder_enable" {
		t.Error("Expected 'reminder_enable', got", session.Job.Output[0].Buttons[0].TextID)
	}
	if session.Job.Output[0].Buttons[0].Callback != "/enable_morning_reminder" {
		t.Error("Expected '/enable_morning_reminder', got", session.Job.Output[0].Buttons[0].Callback)
	}
	if session.Job.Output[0].Buttons[1].TextID != "6:00" {
		t.Error("Expected '6:00', got", session.Job.Output[0].Buttons[1].TextID)
	}
	if session.Job.Output[0].Buttons[1].Callback != "/set_morning_reminder_time 6" {
		t.Error("Expected '/set_morning_reminder_time 6', got", session.Job.Output[0].Buttons[1].Callback)
	}
	if session.Job.Output[0].Buttons[2].TextID != "7:00" {
		t.Error("Expected '7:00', got", session.Job.Output[0].Buttons[2].TextID)
	}
	if session.Job.Output[0].Buttons[2].Callback != "/set_morning_reminder_time 7" {
		t.Error("Expected '/set_morning_reminder_time 7', got", session.Job.Output[0].Buttons[2].Callback)
	}
	if session.Job.Output[0].Buttons[3].TextID != "8:00" {
		t.Error("Expected '8:00', got", session.Job.Output[0].Buttons[3].TextID)
	}
	if session.Job.Output[0].Buttons[3].Callback != "/set_morning_reminder_time 8" {
		t.Error("Expected '/set_morning_reminder_time 8', got", session.Job.Output[0].Buttons[3].Callback)
	}
	if session.Job.Output[0].Buttons[4].TextID != "9:00" {
		t.Error("Expected '9:00', got", session.Job.Output[0].Buttons[4].TextID)
	}
	if session.Job.Output[0].Buttons[4].Callback != "/set_morning_reminder_time 9" {
		t.Error("Expected '/set_morning_reminder_time 9', got", session.Job.Output[0].Buttons[4].Callback)
	}
	if session.Job.Output[0].Buttons[5].TextID != "10:00" {
		t.Error("Expected '10:00', got", session.Job.Output[0].Buttons[5].TextID)
	}
	if session.Job.Output[0].Buttons[5].Callback != "/set_morning_reminder_time 10" {
		t.Error("Expected '/set_morning_reminder_time 10', got", session.Job.Output[0].Buttons[5].Callback)
	}
}

func TestMorningReminderActiveHandler(t *testing.T) {
	enabled := true
	session := createSession(&Job{Command: "/morning_reminder"}, &database.User{}, &database.Settings{
		HasMorningReminder: &enabled,
	}, nil)
	handleMorningReminder(session)

	if len(session.Job.Output[0].Buttons) != 6 {
		t.Errorf("Expected 6 buttons, got %d", len(session.Job.Output[0].Buttons))
	}
	if session.Job.Output[0].Buttons[0].TextID != "reminder_disable" {
		t.Error("Expected 'reminder_disable', got", session.Job.Output[0].Buttons[0].TextID)
	}
}

func TestEveningReminderHandler(t *testing.T) {
	enabled := false
	session := createSession(&Job{Command: "/evening_reminder"}, &database.User{}, &database.Settings{
		HasEveningReminder: &enabled,
	}, nil)
	handleEveningReminder(session)

	if session.Job.Output[0].TextID != "evening_reminder_descr" {
		t.Error("Expected 'evening_reminder_descr', got", session.Job.Output[0].TextID)
	}
	if len(session.Job.Output[0].Buttons) != 6 {
		t.Errorf("Expected 6 buttons, got %d", len(session.Job.Output[0].Buttons))
	}
	if session.Job.Output[0].Buttons[0].TextID != "reminder_enable" {
		t.Error("Expected 'reminder_enable', got", session.Job.Output[0].Buttons[0].TextID)
	}
	if session.Job.Output[0].Buttons[0].Callback != "/enable_evening_reminder" {
		t.Error("Expected '/enable_evening_reminder', got", session.Job.Output[0].Buttons[0].Callback)
	}
	if session.Job.Output[0].Buttons[1].TextID != "19:00" {
		t.Error("Expected '19:00', got", session.Job.Output[0].Buttons[1].TextID)
	}
	if session.Job.Output[0].Buttons[1].Callback != "/set_evening_reminder_time 7" {
		t.Error("Expected '/set_evening_reminder_time 7', got", session.Job.Output[0].Buttons[1].Callback)
	}
	if session.Job.Output[0].Buttons[2].TextID != "20:00" {
		t.Error("Expected '20:00', got", session.Job.Output[0].Buttons[2].TextID)
	}
	if session.Job.Output[0].Buttons[2].Callback != "/set_evening_reminder_time 8" {
		t.Error("Expected '/set_evening_reminder_time 8', got", session.Job.Output[0].Buttons[2].Callback)
	}
	if session.Job.Output[0].Buttons[3].TextID != "21:00" {
		t.Error("Expected '21:00', got", session.Job.Output[0].Buttons[3].TextID)
	}
	if session.Job.Output[0].Buttons[3].Callback != "/set_evening_reminder_time 9" {
		t.Error("Expected '/set_evening_reminder_time 9', got", session.Job.Output[0].Buttons[3].Callback)
	}
	if session.Job.Output[0].Buttons[4].TextID != "22:00" {
		t.Error("Expected '22:00', got", session.Job.Output[0].Buttons[4].TextID)
	}
	if session.Job.Output[0].Buttons[4].Callback != "/set_evening_reminder_time 10" {
		t.Error("Expected '/set_evening_reminder_time 10', got", session.Job.Output[0].Buttons[4].Callback)
	}
	if session.Job.Output[0].Buttons[5].TextID != "23:00" {
		t.Error("Expected '23:00', got", session.Job.Output[0].Buttons[5].TextID)
	}
	if session.Job.Output[0].Buttons[5].Callback != "/set_evening_reminder_time 11" {
		t.Error("Expected '/set_evening_reminder_time 11', got", session.Job.Output[0].Buttons[5].Callback)
	}
}

func TestEveningReminderActiveHandler(t *testing.T) {
	enabled := true
	session := createSession(&Job{Command: "/evening_reminder"}, &database.User{}, &database.Settings{
		HasEveningReminder: &enabled,
	}, nil)
	handleEveningReminder(session)

	if len(session.Job.Output[0].Buttons) != 6 {
		t.Errorf("Expected 6 buttons, got %d", len(session.Job.Output[0].Buttons))
	}
	if session.Job.Output[0].Buttons[0].TextID != "reminder_disable" {
		t.Error("Expected 'reminder_disable', got", session.Job.Output[0].Buttons[0].TextID)
	}
}
