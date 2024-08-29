package bot

import (
	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
)

func sendMessage(chatId int, userId string, text string) {
	saveLastChatId(chatId, userId)
	telegram.SendMessage(chatId, text, nil, nil)
}

func sendMessageWithReply(chatId int, userId string, text string, replyMarkup *telegram.InlineKeyboardMarkup) {
	saveLastChatId(chatId, userId)
	telegram.SendMessage(chatId, text, replyMarkup, nil)
}

func localizeAndSendMessage(chatId int, userId string, locale translator.Locale, text string) {
	localizedMessage := translator.Translate(locale, text)
	sendMessage(chatId, userId, localizedMessage)
}

func localizeAndSendMessageWithReply(chatId int, userId string, locale translator.Locale, text string, replyMarkup *telegram.InlineKeyboardMarkup) {
	localizedMessage := translator.Translate(locale, text)
	sendMessageWithReply(chatId, userId, localizedMessage, replyMarkup)
}
