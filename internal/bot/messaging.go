package bot

import (
	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
)

func SendMessage(chatId int, text string) {
	telegram.SendMessage(chatId, text, nil)
}

func SendMessageWithReply(chatId int, text string, replyMarkup *telegram.InlineKeyboardMarkup) {
	telegram.SendMessage(chatId, text, replyMarkup)
}

func LocalizeAndSendMessage(chatId int, locale translator.Locale, text string) {
	localizedMessage := translator.Translate(locale, text)
	SendMessage(chatId, localizedMessage)
}

func LocalizeAndSendMessageWithReply(chatId int, locale translator.Locale, text string, replyMarkup *telegram.InlineKeyboardMarkup) {
	localizedMessage := translator.Translate(locale, text)
	SendMessageWithReply(chatId, localizedMessage, replyMarkup)
}
