package database

import (
	"testing"
	"time"
)

func TestIsActive(t *testing.T) {
	yesterday := time.Now().AddDate(0, 0, -1)

	user := User{
		Timestamp: &yesterday,
	}
	if !user.IsActive() {
		t.Error("Expected user to be active, got inactive")
	}

	tenDaysAgo := time.Now().AddDate(0, 0, -10)
	user.Timestamp = &tenDaysAgo

	if user.IsActive() {
		t.Error("Expected user to be inactive, got active")
	}
}

func TestIsNonActive(t *testing.T) {
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)

	user := User{
		Timestamp: &sevenDaysAgo,
	}
	if user.IsNonActive() {
		t.Error("Expected user to be active, got inactive")
	}

	oneMonthAgo := time.Now().AddDate(0, 0, -30)
	user.Timestamp = &oneMonthAgo

	if !user.IsNonActive() {
		t.Error("Expected user to be inactive, got active")
	}
}

func TestReminders(t *testing.T) {
	user := User{
		HasMorningReminder: nil,
	}

	if !user.IsMorningReminderEnabled() {
		t.Error("Expected morning reminder to be enabled, got disabled")
	}

	user.HasMorningReminder = new(bool)
	*user.HasMorningReminder = false

	if user.IsMorningReminderEnabled() {
		t.Error("Expected morning reminder to be disabled, got enabled")
	}

	user.HasEveningReminder = nil

	if !user.IsEveningReminderEnabled() {
		t.Error("Expected evening reminder to be enabled, got disabled")
	}

	if user.AreRemindersEnabled() {
		t.Error("Expected reminders to be enabled, got disabled")
	}

	user.HasEveningReminder = new(bool)
	*user.HasEveningReminder = false

	if user.IsEveningReminderEnabled() {
		t.Error("Expected evening reminder to be disabled, got enabled")
	}
}
