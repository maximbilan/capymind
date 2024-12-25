package app

func handleSettings(session *Session) {
	var languageButton JobResultTextButton = JobResultTextButton{
		TextID:   "language",
		Callback: string(Language),
	}
	var timezoneButton JobResultTextButton = JobResultTextButton{
		TextID:   "timezone",
		Callback: string(Timezone),
	}
	setOutputTextWithButtons("settings_descr", []JobResultTextButton{languageButton, timezoneButton}, session)
}
