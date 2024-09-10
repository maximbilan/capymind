package bot

import (
	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
)

// Set the text of the output
func setOutputText(textID string, session *Session) {
	session.Job.Output = &JobResult{
		TextID: textID,
	}
}

// Set the text of the output with buttons
func setOutputTextWithButtons(textID string, buttons []JobResultTextButton, session *Session) {
	session.Job.Output = &JobResult{
		TextID:  textID,
		Buttons: buttons,
	}
}

// Send the output message
func sendOutputMessage(session *Session) {
	locale := session.Locale()
	chatID := session.User.ChatID

	// Prepare the reply markup
	var replyMarkup *telegram.InlineKeyboardMarkup
	if session.Job.Output != nil && len(session.Job.Output.Buttons) > 0 {
		var inlineKeyboard [][]telegram.InlineKeyboardButton
		for _, button := range session.Job.Output.Buttons {
			inlineKeyboard = append(inlineKeyboard, []telegram.InlineKeyboardButton{
				{Text: translator.Translate(locale, button.TextID), CallbackData: &button.Callback},
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

// Send a message to the user
func sendMessage(textID string, session *Session) {
	locale := session.Locale()
	chatID := session.User.ChatID

	// Localize the message
	text := translator.Translate(locale, textID)
	// Send the message
	telegram.SendMessage(chatID, text, nil)
}
