package bot

func handleStart(session *Session) {
	if session.User.Locale == nil {
		// Go onboarding
		handleLanguage(session)
	} else {
		sendWelcome(session)
	}
}

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
