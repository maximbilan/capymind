package bot

import "github.com/capymind/internal/translator"

func handleLanguage(session Session) {
	if session.Job.Input == nil {
		requestLanguage(session)
	} else {
		setupLanguage(session)
	}
}

func requestLanguage(session Session) {
	var enButton JobResultTextButton = JobResultTextButton{
		TextID:   translator.English.String(),
		Callback: string(Language) + " " + translator.EN.String(),
	}
	var ukButton JobResultTextButton = JobResultTextButton{
		TextID:   translator.Ukrainian.String(),
		Callback: string(Language) + " " + translator.UK.String(),
	}
	setOutputTextWithButtons("language_set", []JobResultTextButton{enButton, ukButton}, session)
}

func setupLanguage(session Session) {
	session.User.Locale = session.Job.Input
	setOutputText("language_set", session)
}
