package app

import "github.com/capymind/internal/botservice"

func handleSettings(session *Session) {
	var languageButton botservice.BotResultTextButton = botservice.BotResultTextButton{
		TextID:   "language",
		Locale:   session.Locale(),
		Callback: string(Language),
	}
	var timezoneButton botservice.BotResultTextButton = botservice.BotResultTextButton{
		TextID:   "timezone",
		Locale:   session.Locale(),
		Callback: string(Timezone),
	}
	setOutputTextWithButtons("settings_descr", []botservice.BotResultTextButton{languageButton, timezoneButton}, session)
}
