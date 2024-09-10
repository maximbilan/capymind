package bot

import (
	"github.com/capymind/internal/firestore"
	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
)

type Session struct {
	Job  Job
	User firestore.User
}

// Return the locale of the current user
func (session Session) Locale() translator.Locale {
	if session.User.Locale != nil {
		return translator.Locale(*session.User.Locale)
	}
	return translator.EN
}

// Save the user's data
func (session Session) SaveUser() {
	saveUser(session.User)
}

// Create a session
func createSession(job Job, user firestore.User) Session {
	session := Session{
		Job:  job,
		User: user,
	}
	return session
}

// Handle the session
func handleSession(session Session) {
	command := session.Job.Command
	parameters := session.Job.Parameters

	switch command {
	case Start:
		// handleStart(message, locale)
	case Note:
		// handleNote(message, locale)
	case Last:
		// handleLast(message, locale)
	case Analysis:
		// handleAnalysis(message, locale)
	case Language:
		// handleLanguage(message, locale)
	case Timezone:
		// handleTimezone(message, locale)
	case Help:
		setText(session, "commands_hint")
	case None:
		// Typing mode
	default:
		// Unknown command
		// handleUnknownState(message, locale)
	}
}

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

// Finish the session. Send the output to the user
func finishSession(session Session) {
	locale := session.Locale()
	chatID := session.User.ChatID

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

	text := translator.Translate(locale, session.Job.Output.TextID)
	telegram.SendMessage(chatID, text, replyMarkup)
}
