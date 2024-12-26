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
	var downloadDataButton botservice.BotResultTextButton = botservice.BotResultTextButton{
		TextID:   "download_all_notes",
		Locale:   session.Locale(),
		Callback: string(DownloadData),
	}
	var deleteAccountButton botservice.BotResultTextButton = botservice.BotResultTextButton{
		TextID:   "delete_account",
		Locale:   session.Locale(),
		Callback: string(DeleteAccount),
	}

	setOutputTextWithButtons("settings_descr", []botservice.BotResultTextButton{languageButton, timezoneButton, downloadDataButton, deleteAccountButton}, session)
}
