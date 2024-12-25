package app

import "github.com/capymind/internal/botservice"

func handleSettings(session *Session) {
	var languageButton botservice.BotResultTextButton = botservice.BotResultTextButton{
		TextID:   "language",
		Callback: string(Language),
	}
	var timezoneButton botservice.BotResultTextButton = botservice.BotResultTextButton{
		TextID:   "timezone",
		Callback: string(Timezone),
	}
	setOutputTextWithButtons("settings_descr", []botservice.BotResultTextButton{languageButton, timezoneButton}, session)
}
