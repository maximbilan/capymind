package app

import (
	"github.com/capymind/internal/botservice"
	"github.com/capymind/internal/database"
)

// handleStart is the entry point for the bot. It checks if the user has a locale and timezone set and sends a welcome message
func handleStart(session *Session, settingsStorage database.SettingsStorage) {
	if !session.User.IsOnboarded {
		// Go onboarding
		setOutputText("welcome_onboarding", session)
		handleLanguage(session)
	} else if session.User.SecondsFromUTC == nil {
		// Go onboarding
		handleTimezone(session, settingsStorage)
	} else {
		sendWelcome(session)
	}
}

// Welcome message to the user
func sendWelcome(session *Session) {
	var noteButton botservice.BotResultTextButton = botservice.BotResultTextButton{
		TextID:   "make_record_to_journal",
		Locale:   session.Locale(),
		Callback: string(Note),
	}
	var helpButton botservice.BotResultTextButton = botservice.BotResultTextButton{
		TextID:   "how_to_use",
		Locale:   session.Locale(),
		Callback: string(Help),
	}
	setOutputTextWithButtons("welcome", []botservice.BotResultTextButton{noteButton, helpButton}, session)
}
