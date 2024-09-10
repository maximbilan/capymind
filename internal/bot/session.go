package bot

import (
	"github.com/capymind/internal/firestore"
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

	session.User.LastCommand = string(command)

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
		if session.User.IsTyping && session.Job.Input != nil {
			saveNote(*session.Job.Input, session)
			setText(session, "finish_note")
			session.User.IsTyping = false
		} else {
			// Otherwise show the help message
			setText(session, "commands_hint")
		}
	default:
		// Unknown command
		setText(session, "commands_hint")
	}
}

// Finish the session. Send the output to the user
func finishSession(session Session) {
	// Save the user's data
	session.SaveUser()
	// Prepare the message, localize and send it
	sendMessage(session)
}
