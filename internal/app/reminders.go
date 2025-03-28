package app

import (
	"fmt"
	"log"
	"strconv"

	"github.com/capymind/internal/botservice"
	"github.com/capymind/internal/database"
)

func askForReminders(session *Session) {
	enableButton := botservice.BotResultTextButton{
		TextID:   "reminders_enable_button",
		Locale:   session.Locale(),
		Callback: string(EnableAllReminders),
	}
	continueButton := botservice.BotResultTextButton{
		TextID:   "continue",
		Locale:   session.Locale(),
		Callback: string(SkipReminders),
	}
	setOutputTextWithButtons("onboarding_reminders", []botservice.BotResultTextButton{enableButton, continueButton}, session)
}

func handleReminders(session *Session) {
	settings := *session.Settings

	var switchButton botservice.BotResultTextButton
	if settings.AreRemindersEnabled() {
		switchButton = botservice.BotResultTextButton{
			TextID:   "reminders_disable_button",
			Locale:   session.Locale(),
			Callback: string(DisableAllReminders),
		}
	} else {
		switchButton = botservice.BotResultTextButton{
			TextID:   "reminders_enable_button",
			Locale:   session.Locale(),
			Callback: string(EnableAllReminders),
		}
	}

	var timezoneButton botservice.BotResultTextButton = botservice.BotResultTextButton{
		TextID:   "timezone",
		Locale:   session.Locale(),
		Callback: string(AskForCity),
	}
	var morningReminderButton botservice.BotResultTextButton = botservice.BotResultTextButton{
		TextID:   "morning_reminder_button",
		Locale:   session.Locale(),
		Callback: string(MorningReminder),
	}
	var eveningReminderButton botservice.BotResultTextButton = botservice.BotResultTextButton{
		TextID:   "evening_reminder_button",
		Locale:   session.Locale(),
		Callback: string(EveningReminder),
	}

	setOutputTextWithButtons("reminders_button", []botservice.BotResultTextButton{switchButton, timezoneButton, morningReminderButton, eveningReminderButton}, session)
}

func handleMorningReminder(session *Session) {
	settings := *session.Settings

	var buttons []botservice.BotResultTextButton

	var switchButton botservice.BotResultTextButton
	if settings.IsMorningReminderEnabled() {
		switchButton = botservice.BotResultTextButton{
			TextID:   "reminder_disable",
			Locale:   session.Locale(),
			Callback: string(DisableMorningReminder),
		}
	} else {
		switchButton = botservice.BotResultTextButton{
			TextID:   "reminder_enable",
			Locale:   session.Locale(),
			Callback: string(EnableMorningReminder),
		}
	}
	buttons = append(buttons, switchButton)

	var timeOffsets []int
	for i := 6; i < 11; i++ {
		timeOffsets = append(timeOffsets, i)
	}

	for _, offset := range timeOffsets {
		time := fmt.Sprintf("%d:00", offset)
		buttons = append(buttons, botservice.BotResultTextButton{
			TextID:   time,
			Locale:   session.Locale(),
			Callback: string(SetMorningReminderTime) + " " + fmt.Sprintf("%d", offset),
		})
	}

	setOutputTextWithButtons("morning_reminder_descr", buttons, session)
}

func handleEveningReminder(session *Session) {
	settings := *session.Settings

	var buttons []botservice.BotResultTextButton

	var switchButton botservice.BotResultTextButton
	if settings.IsEveningReminderEnabled() {
		switchButton = botservice.BotResultTextButton{
			TextID:   "reminder_disable",
			Locale:   session.Locale(),
			Callback: string(DisableEveningReminder),
		}
	} else {
		switchButton = botservice.BotResultTextButton{
			TextID:   "reminder_enable",
			Locale:   session.Locale(),
			Callback: string(EnableEveningReminder),
		}
	}
	buttons = append(buttons, switchButton)

	var timeOffsets []int
	for i := 19; i < 24; i++ {
		timeOffsets = append(timeOffsets, i)
	}

	for _, offset := range timeOffsets {
		time := fmt.Sprintf("%d:00", offset)
		buttons = append(buttons, botservice.BotResultTextButton{
			TextID:   time,
			Locale:   session.Locale(),
			Callback: string(SetEveningReminderTime) + " " + fmt.Sprintf("%d", offset-12),
		})
	}

	setOutputTextWithButtons("evening_reminder_descr", buttons, session)
}

func enableAllReminders(session *Session, settingsStorage database.SettingsStorage) {
	session.Settings.HasMorningReminder = new(bool)
	*session.Settings.HasMorningReminder = true
	session.Settings.HasEveningReminder = new(bool)
	*session.Settings.HasEveningReminder = true

	saveSettings(session.Context, session.User.ID, *session.Settings, settingsStorage)
	setOutputText("reminders_enabled", session)

	if session.User.SecondsFromUTC == nil || session.Settings.SecondsFromUTC == nil {
		requestTimezone(session)
	}
}

func disableAllReminders(session *Session, settingsStorage database.SettingsStorage) {
	session.Settings.HasMorningReminder = new(bool)
	*session.Settings.HasMorningReminder = false
	session.Settings.HasEveningReminder = new(bool)
	*session.Settings.HasEveningReminder = false

	saveSettings(session.Context, session.User.ID, *session.Settings, settingsStorage)
	setOutputText("reminders_disabled", session)
}

func enableMorningReminder(session *Session, settingsStorage database.SettingsStorage) {
	session.Settings.HasMorningReminder = new(bool)
	*session.Settings.HasMorningReminder = true

	saveSettings(session.Context, session.User.ID, *session.Settings, settingsStorage)
	setOutputText("reminder_set", session)

	if session.User.SecondsFromUTC == nil || session.Settings.SecondsFromUTC == nil {
		requestTimezone(session)
	}
}

func disableMorningReminder(session *Session, settingsStorage database.SettingsStorage) {
	session.Settings.HasMorningReminder = new(bool)
	*session.Settings.HasMorningReminder = false

	saveSettings(session.Context, session.User.ID, *session.Settings, settingsStorage)
	setOutputText("reminder_unset", session)
}

func enableEveningReminder(session *Session, settingsStorage database.SettingsStorage) {
	session.Settings.HasEveningReminder = new(bool)
	*session.Settings.HasEveningReminder = true

	saveSettings(session.Context, session.User.ID, *session.Settings, settingsStorage)
	setOutputText("reminder_set", session)

	if session.User.SecondsFromUTC == nil || session.Settings.SecondsFromUTC == nil {
		requestTimezone(session)
	}
}

func disableEveningReminder(session *Session, settingsStorage database.SettingsStorage) {
	session.Settings.HasEveningReminder = new(bool)
	*session.Settings.HasEveningReminder = false

	saveSettings(session.Context, session.User.ID, *session.Settings, settingsStorage)
	setOutputText("reminder_unset", session)
}

func parseReminderTime(param string) *int {
	offset, err := strconv.Atoi(param)
	if err != nil {
		log.Printf("[Bot] Error parsing timezone: %v", err)
		return nil
	}
	return &offset
}

func setMorningReminderOffset(session *Session, settingsStorage database.SettingsStorage) {
	offset := parseReminderTime(session.Job.Parameters[0])
	if offset == nil {
		return
	}

	session.Settings.MorningReminderOffset = new(int)
	*session.Settings.MorningReminderOffset = *offset
	session.Settings.HasMorningReminder = new(bool)
	*session.Settings.HasMorningReminder = true

	saveSettings(session.Context, session.User.ID, *session.Settings, settingsStorage)
	setOutputText("reminder_set", session)

	if session.User.SecondsFromUTC == nil || session.Settings.SecondsFromUTC == nil {
		requestTimezone(session)
	}
}

func setEveningReminderOffset(session *Session, settingsStorage database.SettingsStorage) {
	offset := parseReminderTime(session.Job.Parameters[0])
	if offset == nil {
		return
	}

	session.Settings.EveningReminderOffset = new(int)
	*session.Settings.EveningReminderOffset = *offset
	session.Settings.HasEveningReminder = new(bool)
	*session.Settings.HasEveningReminder = true

	saveSettings(session.Context, session.User.ID, *session.Settings, settingsStorage)
	setOutputText("reminder_set", session)

	if session.User.SecondsFromUTC == nil || session.Settings.SecondsFromUTC == nil {
		requestTimezone(session)
	}
}

func skipReminders(session *Session, settingsStorage database.SettingsStorage) {
	session.User.IsOnboarded = true

	session.Settings.HasMorningReminder = new(bool)
	*session.Settings.HasMorningReminder = false
	session.Settings.HasEveningReminder = new(bool)
	*session.Settings.HasEveningReminder = false

	saveSettings(session.Context, session.User.ID, *session.Settings, settingsStorage)
	sendWelcome(session)
}
