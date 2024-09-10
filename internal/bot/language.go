package bot

import "github.com/capymind/internal/translator"

func handleLanguage(session *Session) {
	if len(session.Job.Parameters) == 0 {
		requestLanguage(session)
	} else {
		setupLanguage(session)
	}
}

func requestLanguage(session *Session) {
	var enButton JobResultTextButton = JobResultTextButton{
		TextID:   translator.English.String(),
		Callback: string(Language) + " " + translator.EN.String(),
	}
	var ukButton JobResultTextButton = JobResultTextButton{
		TextID:   translator.Ukrainian.String(),
		Callback: string(Language) + " " + translator.UK.String(),
	}
	setOutputTextWithButtons("language_select", []JobResultTextButton{enButton, ukButton}, session)
}

func setupLanguage(session *Session) {
	session.User.Locale = &session.Job.Parameters[0]
	setOutputText("locale_set", session)
}
