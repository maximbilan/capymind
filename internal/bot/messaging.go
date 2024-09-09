package bot

import (
	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
)

func sendMessage(chatID int64, userID string, text string) {
	saveLastChatID(chatID, userID)
	telegram.SendMessage(chatID, text, nil)
}

func sendMessageWithReply(chatID int64, userID string, text string, replyMarkup *telegram.InlineKeyboardMarkup) {
	saveLastChatID(chatID, userID)
	telegram.SendMessage(chatID, text, replyMarkup)
}

func localizeAndSendMessage(chatID int64, userID string, locale translator.Locale, text string) {
	localizedMessage := translator.Translate(locale, text)
	sendMessage(chatID, userID, localizedMessage)
}

func localizeAndSendMessageWithReply(chatID int64, userID string, locale translator.Locale, text string, replyMarkup *telegram.InlineKeyboardMarkup) {
	localizedMessage := translator.Translate(locale, text)
	sendMessageWithReply(chatID, userID, localizedMessage, replyMarkup)
}
