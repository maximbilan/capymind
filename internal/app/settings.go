package app

import (
	"context"
	"log"

	"github.com/capymind/internal/botservice"
	"github.com/capymind/internal/database"
)

func getSettings(ctx *context.Context, userId string, settingsStorage database.SettingsStorage) *database.Settings {
	settings, err := settingsStorage.GetSettings(ctx, userId)
	if err != nil {
		// First time user settings
		settings = &database.Settings{}
	}
	return settings
}

func saveSettings(ctx *context.Context, userId string, settings database.Settings, settingsStorage database.SettingsStorage) {
	//coverage:ignore
	err := settingsStorage.SaveSettings(ctx, userId, settings)
	if err != nil {
		log.Printf("[Settings] Error saving settings to firestore, %s", err.Error())
	}
}

func handleSettings(session *Session) {
	var languageButton botservice.BotResultTextButton = botservice.BotResultTextButton{
		TextID:   "language",
		Locale:   session.Locale(),
		Callback: string(Language),
	}
	var remindersButton botservice.BotResultTextButton = botservice.BotResultTextButton{
		TextID:   "reminders_button",
		Locale:   session.Locale(),
		Callback: string(Reminders),
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

	setOutputTextWithButtons("settings_descr", []botservice.BotResultTextButton{languageButton, remindersButton, downloadDataButton, deleteAccountButton}, session)
}
