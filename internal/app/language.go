package app

import (
	"github.com/capymind/internal/botservice"
	"github.com/capymind/internal/translator"
)

// Handle the language command
func handleLanguage(session *Session) {
	if len(session.Job.Parameters) == 0 {
		requestLanguage(session)
	} else {
		setupLanguage(session)
	}
}

// Request the language select
func requestLanguage(session *Session) {
	var enButton botservice.BotResultTextButton = botservice.BotResultTextButton{
		TextID:   translator.English.String(),
		Callback: string(Language) + " " + translator.EN.String(),
	}
	var ukButton botservice.BotResultTextButton = botservice.BotResultTextButton{
		TextID:   translator.Ukrainian.String(),
		Callback: string(Language) + " " + translator.UK.String(),
	}
	setOutputTextWithButtons("language_select", []botservice.BotResultTextButton{enButton, ukButton}, session)
}

// Set the language
func setupLanguage(session *Session) {
	session.User.Locale = &session.Job.Parameters[0]

	if session.User.SecondsFromUTC == nil {
		setOutputText("locale_set", session)
		requestTimezone(session)
	} else {
		setOutputText("locale_set", session)
	}
}
