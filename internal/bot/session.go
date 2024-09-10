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

	session.User.LastCommand = string(command)

	switch command {
	case Start:
		// handleStart(message, locale)
	case Note:
		// handleNote(message, locale)
	case Last:
		// handleLast(message, locale)
	case Analysis:
		handleAnalysis(session)
	case Language:
		var enButton JobResultTextButton = JobResultTextButton{
			TextID:   translator.English.String(),
			Callback: translator.GetLocaleParameter(translator.EN),
		}
		var ukButton JobResultTextButton = JobResultTextButton{
			TextID:   translator.Ukrainian.String(),
			Callback: translator.GetLocaleParameter(translator.UK),
		}
		setOutputTextWithButtons("language_set", []JobResultTextButton{enButton, ukButton}, session)
	case Timezone:
		handleTimezone(session)
	case Help:
		setOutputText("commands_hint", session)
	case None:
		// Typing mode
		if session.User.IsTyping && session.Job.Input != nil {
			finishNote(session)
		} else {
			// Otherwise show the help message
			setOutputText("commands_hint", session)
		}
	default:
		// Unknown command
		setOutputText("commands_hint", session)
	}
}

// Finish the session. Send the output to the user
func finishSession(session Session) {
	// Save the user's data
	session.SaveUser()
	// Prepare the message, localize and send it
	sendOutputMessage(session)
}