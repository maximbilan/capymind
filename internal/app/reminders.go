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
		Callback: string(Timezone),
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
	user := *session.User
	settings := *session.Settings

	settings.HasMorningReminder = new(bool)
	*settings.HasMorningReminder = true
	settings.HasEveningReminder = new(bool)
	*settings.HasEveningReminder = true

	saveSettings(session.Context, session.User.ID, settings, settingsStorage)
	setOutputText("reminders_enabled", session)

	if user.SecondsFromUTC == nil || settings.SecondsFromUTC == nil {
		requestTimezone(session)
	}
}

func disableAllReminders(session *Session, settingsStorage database.SettingsStorage) {
	settings := *session.Settings

	settings.HasMorningReminder = new(bool)
	*settings.HasMorningReminder = false
	settings.HasEveningReminder = new(bool)
	*settings.HasEveningReminder = false

	saveSettings(session.Context, session.User.ID, settings, settingsStorage)
	setOutputText("reminders_disabled", session)
}

func enableMorningReminder(session *Session, settingsStorage database.SettingsStorage) {
	user := *session.User
	settings := *session.Settings

	settings.HasMorningReminder = new(bool)
	*settings.HasMorningReminder = true

	saveSettings(session.Context, session.User.ID, settings, settingsStorage)
	setOutputText("reminder_set", session)

	if user.SecondsFromUTC == nil || settings.SecondsFromUTC == nil {
		requestTimezone(session)
	}
}

func disableMorningReminder(session *Session, settingsStorage database.SettingsStorage) {
	settings := *session.Settings

	settings.HasMorningReminder = new(bool)
	*settings.HasMorningReminder = false

	saveSettings(session.Context, session.User.ID, settings, settingsStorage)
	setOutputText("reminder_unset", session)
}

func enableEveningReminder(session *Session, settingsStorage database.SettingsStorage) {
	user := *session.User
	settings := *session.Settings

	settings.HasEveningReminder = new(bool)
	*settings.HasEveningReminder = true

	saveSettings(session.Context, session.User.ID, settings, settingsStorage)
	setOutputText("reminder_set", session)

	if user.SecondsFromUTC == nil || settings.SecondsFromUTC == nil {
		requestTimezone(session)
	}
}

func disableEveningReminder(session *Session, settingsStorage database.SettingsStorage) {
	settings := *session.Settings

	settings.HasEveningReminder = new(bool)
	*settings.HasEveningReminder = false

	saveSettings(session.Context, session.User.ID, settings, settingsStorage)
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

	settings := *session.Settings
	settings.MorningReminderOffset = new(int)
	*settings.MorningReminderOffset = *offset
	settings.HasMorningReminder = new(bool)
	*settings.HasMorningReminder = true

	saveSettings(session.Context, session.User.ID, settings, settingsStorage)
	setOutputText("reminder_set", session)

	user := *session.User
	if user.SecondsFromUTC == nil || settings.SecondsFromUTC == nil {
		requestTimezone(session)
	}
}

func setEveningReminderOffset(session *Session, settingsStorage database.SettingsStorage) {
	offset := parseReminderTime(session.Job.Parameters[0])
	if offset == nil {
		return
	}

	settings := *session.Settings
	settings.EveningReminderOffset = new(int)
	*settings.EveningReminderOffset = *offset
	settings.HasEveningReminder = new(bool)
	*settings.HasEveningReminder = true

	saveSettings(session.Context, session.User.ID, settings, settingsStorage)
	setOutputText("reminder_set", session)

	user := *session.User
	if user.SecondsFromUTC == nil || settings.SecondsFromUTC == nil {
		requestTimezone(session)
	}
}

func skipReminders(session *Session, settingsStorage database.SettingsStorage) {
	session.User.IsOnboarded = true

	settings := *session.Settings
	settings.HasMorningReminder = new(bool)
	*settings.HasMorningReminder = false
	settings.HasEveningReminder = new(bool)
	*settings.HasEveningReminder = false

	saveSettings(session.Context, session.User.ID, settings, settingsStorage)
	sendWelcome(session)
}
