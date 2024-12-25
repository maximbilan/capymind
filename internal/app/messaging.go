package app

import (
	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
)

func appendJobResult(jobResult JobResult, session *Session) {
	output := session.Job.Output
	session.Job.Output = append(output, jobResult)
}

// Set the text of the output
func setOutputText(textID string, session *Session) {
	jobResult := JobResult{
		TextID: textID,
	}
	appendJobResult(jobResult, session)
}

// Set the text of the output with buttons
func setOutputTextWithButtons(textID string, buttons []JobResultTextButton, session *Session) {
	jobResult := JobResult{
		TextID:  textID,
		Buttons: buttons,
	}
	appendJobResult(jobResult, session)
}

// Send the output messages
func sendOutputMessages(session *Session) {
	if len(session.Job.Output) == 0 {
		return
	}

	for _, jobResult := range session.Job.Output {
		sendJobResult(jobResult, session)
	}
}

// Send the output messages
func sendJobResult(jobResult JobResult, session *Session) {
	locale := session.Locale()
	chatID := session.User.ChatID

	// Prepare the reply markup
	var replyMarkup *telegram.InlineKeyboardMarkup
	if len(jobResult.Buttons) > 0 {
		var inlineKeyboard [][]telegram.InlineKeyboardButton
		for _, button := range jobResult.Buttons {
			inlineKeyboard = append(inlineKeyboard, []telegram.InlineKeyboardButton{
				{Text: translator.Translate(locale, button.TextID), CallbackData: &button.Callback},
			})
		}

		replyMarkup = &telegram.InlineKeyboardMarkup{
			InlineKeyboard: inlineKeyboard,
		}
	}

	// Localize the message
	text := translator.Translate(locale, jobResult.TextID)
	// Send the message
	telegram.SendMessage(chatID, text, replyMarkup)
}

// Send a message to the user (Immediately)
func sendMessage(textID string, session *Session) {
	locale := session.Locale()
	chatID := session.User.ChatID

	// Localize the message
	text := translator.Translate(locale, textID)
	// Send the message
	telegram.SendMessage(chatID, text, nil)
}
