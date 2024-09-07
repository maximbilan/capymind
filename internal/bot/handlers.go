package bot

import (
	"fmt"

	"github.com/capymind/internal/analysis"
	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
	"github.com/capymind/internal/utils"
)

func handleUser(message telegram.Message) {
	if message.Text == "" {
		return
	}
	createOrUpdateUser(message)
}

func handleStart(message telegram.Message, locale translator.Locale) {
	userId := fmt.Sprintf("%d", message.From.ID)
	localizeAndSendMessage(message.Chat.Id, userId, locale, "welcome")
	handleUser(message)
}

func handleNote(message telegram.Message, locale translator.Locale) {
	userId := message.From.ID
	userIdStr := fmt.Sprintf("%d", userId)
	localizeAndSendMessage(message.Chat.Id, userIdStr, locale, "start_note")
	userIds.Append(userId)
}

func handleLast(message telegram.Message, locale translator.Locale) {
	userId := fmt.Sprintf("%d", message.From.ID)
	note := getLastNote(message)
	if note != nil {
		var response string = translator.Translate(locale, "your_last_note") + note.Text
		sendMessage(message.Chat.Id, userId, response)
	} else {
		replyMarkup := telegram.InlineKeyboardMarkup{
			InlineKeyboard: [][]telegram.InlineKeyboardButton{
				{
					{Text: translator.Translate(locale, "make_record_to_journal"), CallbackData: "note_make"},
				},
			},
		}
		localizeAndSendMessageWithReply(message.Chat.Id, userId, locale, "no_notes", &replyMarkup)
	}
}

func handleAnalysis(message telegram.Message, locale translator.Locale) {
	userId := fmt.Sprintf("%d", message.From.ID)
	notes := getNotes(message)
	if len(notes) > 0 {
		var strings []string
		for _, note := range notes {
			if note.Text != "" {
				strings = append(strings, note.Text)
			}
		}

		localizeAndSendMessage(message.Chat.Id, userId, locale, "analysis_waiting")
		analysis := analysis.Request(strings, locale)
		if analysis != nil {
			sendMessage(message.Chat.Id, userId, *analysis)
			return
		}
	}

	replyMarkup := telegram.InlineKeyboardMarkup{
		InlineKeyboard: [][]telegram.InlineKeyboardButton{
			{
				{Text: translator.Translate(locale, "make_record_to_journal"), CallbackData: "note_make"},
			},
		},
	}
	localizeAndSendMessageWithReply(message.Chat.Id, userId, locale, "no_analysis", &replyMarkup)
}

func handleLocale(message telegram.Message, locale translator.Locale) {
	userId := fmt.Sprintf("%d", message.From.ID)
	replyMarkup := telegram.InlineKeyboardMarkup{
		InlineKeyboard: [][]telegram.InlineKeyboardButton{
			{
				{Text: translator.English.String(), CallbackData: translator.GetLocaleParameter(translator.EN)},
				{Text: translator.Ukrainian.String(), CallbackData: translator.GetLocaleParameter(translator.UK)},
			},
		},
	}
	localizeAndSendMessageWithReply(message.Chat.Id, userId, locale, "language_select", &replyMarkup)
}

func handleTimezone(message telegram.Message, locale translator.Locale) {
	userId := fmt.Sprintf("%d", message.From.ID)
	timeZones := utils.GetTimeZones()

	var inlineKeyboard [][]telegram.InlineKeyboardButton
	for _, tz := range timeZones {
		inlineKeyboard = append(inlineKeyboard, []telegram.InlineKeyboardButton{
			{Text: tz.String(), CallbackData: utils.GetTimezoneParameter(tz)},
		})
	}

	replyMarkup := telegram.InlineKeyboardMarkup{
		InlineKeyboard: inlineKeyboard,
	}

	localizeAndSendMessageWithReply(message.Chat.Id, userId, locale, "timezone_select", &replyMarkup)
}

func handleUnknownState(message telegram.Message, locale translator.Locale) {
	userId := message.From.ID
	userIdStr := fmt.Sprintf("%d", userId)
	if userIds.Contains(userId) {
		saveNote(message)
		localizeAndSendMessage(message.Chat.Id, userIdStr, locale, "finish_note")
		userIds.Remove(userId)
	} else {
		handleHelp(message, locale)
	}
}

func handleHelp(message telegram.Message, locale translator.Locale) {
	userId := fmt.Sprintf("%d", message.From.ID)
	localizeAndSendMessage(message.Chat.Id, userId, locale, "commands_hint")
}
