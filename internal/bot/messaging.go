package bot

import (
	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
)

// func sendMessage(chatID int64, userID string, text string) {
// 	saveLastChatID(chatID, userID)
// 	telegram.SendMessage(chatID, text, nil)
// }

// func sendMessageWithReply(chatID int64, userID string, text string, replyMarkup *telegram.InlineKeyboardMarkup) {
// 	saveLastChatID(chatID, userID)
// 	telegram.SendMessage(chatID, text, replyMarkup)
// }

// func localizeAndSendMessage(chatID int64, userID string, locale translator.Locale, text string) {
// 	localizedMessage := translator.Translate(locale, text)
// 	sendMessage(chatID, userID, localizedMessage)
// }

// func localizeAndSendMessageWithReply(chatID int64, userID string, locale translator.Locale, text string, replyMarkup *telegram.InlineKeyboardMarkup) {
// 	localizedMessage := translator.Translate(locale, text)
// 	sendMessageWithReply(chatID, userID, localizedMessage, replyMarkup)
// }

// Set the text of the output
func setText(session Session, textID string) {
	session.Job.Output = &JobResult{
		TextID: textID,
	}
}

// Set the text of the output with buttons
func setTextWithButtons(session Session, textID string, buttons []JobResultTextButton) {
	session.Job.Output = &JobResult{
		TextID:  textID,
		Buttons: buttons,
	}
}

func sendMessage(session Session) {
	locale := session.Locale()
	chatID := session.User.ChatID

	// Prepare the reply markup
	var replyMarkup *telegram.InlineKeyboardMarkup
	if len(session.Job.Output.Buttons) > 0 {
		var inlineKeyboard [][]telegram.InlineKeyboardButton

		for _, button := range session.Job.Output.Buttons {
			callbackData := button.Callback
			inlineKeyboard = append(inlineKeyboard, []telegram.InlineKeyboardButton{
				{Text: translator.Translate(locale, button.TextID), CallbackData: &callbackData},
			})
		}

		replyMarkup = &telegram.InlineKeyboardMarkup{
			InlineKeyboard: inlineKeyboard,
		}
	}

	// Localize the message
	text := translator.Translate(locale, session.Job.Output.TextID)
	// Send the message
	telegram.SendMessage(chatID, text, replyMarkup)
}
