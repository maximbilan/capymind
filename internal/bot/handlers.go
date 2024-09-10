package bot

import (
	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
	"github.com/capymind/internal/utils"
)

func handleStart(message telegram.Message, locale translator.Locale) {
	userID := message.UserID()
	chatID := message.ChatID()

	userLocale := getUserLocaleByUserID(userID)
	if userLocale == nil {
		handleLanguage(message, locale)
	} else {
		sendStartMessage(chatID, userID, &message.From.UserName, *userLocale)
	}
}

func sendStartMessage(chatID int64, userID string, name *string, locale translator.Locale) {
	noteCallbackData := "note_make"
	helpCallbackData := "help"
	replyMarkup := telegram.InlineKeyboardMarkup{
		InlineKeyboard: [][]telegram.InlineKeyboardButton{
			{
				{Text: translator.Translate(locale, "make_record_to_journal_short"), CallbackData: &noteCallbackData},
				{Text: translator.Translate(locale, "how_to_use"), CallbackData: &helpCallbackData},
			},
		},
	}
	localizeAndSendMessageWithReply(chatID, userID, locale, "welcome", &replyMarkup)
	createOrUpdateUser(chatID, userID, name)
}

func handleNote(message telegram.Message, locale translator.Locale) {
	userID := message.UserID()
	chatID := message.ChatID()
	localizeAndSendMessage(chatID, userID, locale, "start_note")
	startWritingMode(userID)
}

func handleLast(message telegram.Message, locale translator.Locale) {
	userID := message.UserID()
	chatID := message.ChatID()
	note := getLastNote(message)
	if note != nil {
		var response string = translator.Translate(locale, "your_last_note") + note.Text
		sendMessage(chatID, userID, response)
	} else {
		noteCallbackData := "note_make"
		replyMarkup := telegram.InlineKeyboardMarkup{
			InlineKeyboard: [][]telegram.InlineKeyboardButton{
				{
					{Text: translator.Translate(locale, "make_record_to_journal"), CallbackData: &noteCallbackData},
				},
			},
		}
		localizeAndSendMessageWithReply(chatID, userID, locale, "no_notes", &replyMarkup)
	}
}

// func handleAnalysis(message telegram.Message, locale translator.Locale) {
// 	userID := message.UserID()
// 	chatID := message.ChatID()
// 	notes := getNotes(message)
// 	if len(notes) > 0 {
// 		var strings []string
// 		for _, note := range notes {
// 			if note.Text != "" {
// 				strings = append(strings, note.Text)
// 			}
// 		}

// 		localizeAndSendMessage(message.Chat.ID, userID, locale, "analysis_waiting")
// 		analysis := analysis.Request(strings, locale)
// 		if analysis != nil {
// 			sendMessage(chatID, userID, *analysis)
// 			return
// 		}
// 	}

// 	noteCallbackData := "note_make"
// 	replyMarkup := telegram.InlineKeyboardMarkup{
// 		InlineKeyboard: [][]telegram.InlineKeyboardButton{
// 			{
// 				{Text: translator.Translate(locale, "make_record_to_journal"), CallbackData: &noteCallbackData},
// 			},
// 		},
// 	}
// 	localizeAndSendMessageWithReply(message.Chat.ID, userID, locale, "no_analysis", &replyMarkup)
// }

func handleLanguage(message telegram.Message, locale translator.Locale) {
	userID := message.UserID()
	chatID := message.ChatID()
	sendLanguageSetMessage(chatID, userID, locale)
}

func sendLanguageSetMessage(chatID int64, userID string, locale translator.Locale) {
	enCallbackData := translator.GetLocaleParameter(translator.EN)
	ukCallbackData := translator.GetLocaleParameter(translator.UK)
	replyMarkup := telegram.InlineKeyboardMarkup{
		InlineKeyboard: [][]telegram.InlineKeyboardButton{
			{
				{Text: translator.English.String(), CallbackData: &enCallbackData},
				{Text: translator.Ukrainian.String(), CallbackData: &ukCallbackData},
			},
		},
	}
	localizeAndSendMessageWithReply(chatID, userID, locale, "language_select", &replyMarkup)
}

// func handleTimezone(message telegram.Message, locale translator.Locale) {
// 	userID := message.UserID()
// 	chatID := message.ChatID()
// 	sendTimezoneSetMessage(chatID, userID, locale)
// }

func sendTimezoneSetMessage(chatID int64, userID string, locale translator.Locale) {
	timeZones := utils.GetTimeZones()

	var inlineKeyboard [][]telegram.InlineKeyboardButton
	for _, tz := range timeZones {
		callbackData := utils.GetTimezoneParameter(tz)
		inlineKeyboard = append(inlineKeyboard, []telegram.InlineKeyboardButton{
			{Text: tz.String(), CallbackData: &callbackData},
		})
	}

	replyMarkup := telegram.InlineKeyboardMarkup{
		InlineKeyboard: inlineKeyboard,
	}

	localizeAndSendMessageWithReply(chatID, userID, locale, "timezone_select", &replyMarkup)
}

func handleUnknownState(message telegram.Message, locale translator.Locale) {
	userID := message.UserID()
	chatID := message.ChatID()

	if isWriting(userID) {
		saveNote(message)
		localizeAndSendMessage(chatID, userID, locale, "finish_note")
		stopWritingMode(userID)
	} else {
		handleHelp(message, locale)
	}
}

func handleHelp(message telegram.Message, locale translator.Locale) {
	userID := message.UserID()
	chatID := message.ChatID()
	sendHelpMessage(chatID, userID, locale)
}

func sendHelpMessage(chatID int64, userId string, locale translator.Locale) {
	localizeAndSendMessage(chatID, userId, locale, "commands_hint")
}
