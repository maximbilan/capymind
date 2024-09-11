package bot

// handleStart is the entry point for the bot. It checks if the user has a locale and timezone set and sends a welcome message
func handleStart(session *Session) {
	if !session.User.IsOnboarded {
		// Go onboarding
		sendMessage("welcome_onboarding", session)
		handleLanguage(session)
	} else if session.User.SecondsFromUTC == nil {
		// Go onboarding
		handleTimezone(session)
	} else {
		sendWelcome(session)
	}
}

// Welcome message to the user
func sendWelcome(session *Session) {
	var noteButton JobResultTextButton = JobResultTextButton{
		TextID:   "make_record_to_journal",
		Callback: string(Note),
	}
	var helpButton JobResultTextButton = JobResultTextButton{
		TextID:   "how_to_use",
		Callback: string(Help),
	}
	setOutputTextWithButtons("welcome", []JobResultTextButton{noteButton, helpButton}, session)
}
