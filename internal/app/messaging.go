package app

import (
	"github.com/capymind/internal/botservice"
	"github.com/capymind/internal/translator"
)

func appendJobResult(jobResult botservice.BotResult, session *Session) {
	output := session.Job.Output
	session.Job.Output = append(output, jobResult)
}

// Set the text of the output
func setOutputText(textID string, session *Session) {
	jobResult := botservice.BotResult{
		TextID: textID,
		Locale: session.Locale(),
	}
	appendJobResult(jobResult, session)
}

// Set raw text (no translation key), useful for backend-generated answers
func setOutputRawText(text string, session *Session) {
	jobResult := botservice.BotResult{
		TextID: text,
		Locale: session.Locale(),
	}
	appendJobResult(jobResult, session)
}

// Set the text of the output with buttons
func setOutputTextWithButtons(textID string, buttons []botservice.BotResultTextButton, session *Session) {
	jobResult := botservice.BotResult{
		TextID:  textID,
		Locale:  session.Locale(),
		Buttons: buttons,
	}
	appendJobResult(jobResult, session)
}

// Send the output messages
func sendOutputMessages(session *Session) {
	//coverage:ignore
	if len(session.Job.Output) == 0 {
		return
	}

	for _, jobResult := range session.Job.Output {
		sendJobResult(jobResult, session)
	}
}

// Send the output messages
func sendJobResult(jobResult botservice.BotResult, session *Session) {
	//coverage:ignore
	bot.SendResult(session.User.ChatID, jobResult)
}

// Send a message to the user (Immediately)
func sendMessage(textID string, session *Session) {
	//coverage:ignore
	locale := session.Locale()
	chatID := session.User.ChatID

	// Localize the message
	text := translator.Translate(locale, textID)
	// Send the message
	bot.SendMessage(chatID, text)
}
