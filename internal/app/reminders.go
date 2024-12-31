package app

import (
	"fmt"

	"github.com/capymind/internal/botservice"
)

func handleReminders(session *Session) {
	user := *session.User

	var switchButton botservice.BotResultTextButton
	if user.AreRemindersEnabled() {
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
	user := *session.User

	var buttons []botservice.BotResultTextButton

	var switchButton botservice.BotResultTextButton
	if user.IsMorningReminderEnabled() {
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

	setOutputTextWithButtons("morning_reminder_button", buttons, session)
}

func handleEveningReminder(session *Session) {
	user := *session.User

	var buttons []botservice.BotResultTextButton

	var switchButton botservice.BotResultTextButton
	if user.IsEveningReminderEnabled() {
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
			Callback: string(SetEveningReminderTime) + " " + fmt.Sprintf("%d", offset),
		})
	}

	setOutputTextWithButtons("evening_reminder_button", buttons, session)
}
