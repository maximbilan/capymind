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
		Callback: translator.GetLocaleParameter(translator.EN),
	}
	var ukButton JobResultTextButton = JobResultTextButton{
		TextID:   translator.Ukrainian.String(),
		Callback: translator.GetLocaleParameter(translator.UK),
	}
	setOutputTextWithButtons("language_set", []JobResultTextButton{enButton, ukButton}, session)
}

func setupLanguage(session Session) {
	session.User.Locale = session.Job.Input
	setOutputText("language_set", session)
}
