package database

import "testing"

func TestReminders(t *testing.T) {
	settings := Settings{
		HasMorningReminder: nil,
	}

	if !settings.IsMorningReminderEnabled() {
		t.Error("Expected morning reminder to be enabled, got disabled")
	}

	settings.HasMorningReminder = new(bool)
	*settings.HasMorningReminder = false

	if settings.IsMorningReminderEnabled() {
		t.Error("Expected morning reminder to be disabled, got enabled")
	}

	settings.HasEveningReminder = nil

	if !settings.IsEveningReminderEnabled() {
		t.Error("Expected evening reminder to be enabled, got disabled")
	}

	if settings.AreRemindersEnabled() {
		t.Error("Expected reminders to be enabled, got disabled")
	}

	settings.HasEveningReminder = new(bool)
	*settings.HasEveningReminder = false

	if settings.IsEveningReminderEnabled() {
		t.Error("Expected evening reminder to be disabled, got enabled")
	}
}
