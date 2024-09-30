package bot

import (
	"context"

	"github.com/capymind/internal/firestore"
	"github.com/capymind/internal/translator"
)

type Session struct {
	Job     *Job
	User    *firestore.User
	Context *context.Context
}

// Return the locale of the current user
func (session *Session) Locale() translator.Locale {
	if session.User.Locale != nil {
		return translator.Locale(*session.User.Locale)
	}
	return translator.EN
}

// Save the user's data
func (session *Session) SaveUser() {
	saveUser(session.User, session.Context)
}

// Create a session
func createSession(job *Job, user *firestore.User, context *context.Context) *Session {
	session := Session{
		Job:     job,
		User:    user,
		Context: context,
	}
	return &session
}

// Handle the session
func handleSession(session *Session) {
	command := session.Job.Command
	session.User.LastCommand = string(command)

	switch command {
	case Start:
		handleStart(session)
	case Note:
		startNote(session)
	case Last:
		handleLastNote(session)
	case Analysis:
		handleAnalysis(session)
	case Language:
		handleLanguage(session)
	case Timezone:
		handleTimezone(session)
	case Help:
		setOutputText("commands_hint", session)
	case SleepAnalysis:
		handleSleepAnalysis(session)
	case WeeklyAnalysis:
		handleWeeklyAnalysis(session)
	case NoteCount:
		handleNoteCount(session)
	case TotalUserCount:
		handleTotalUserCount(session)
	case TotalNoteCount:
		handleTotalNoteCount(session)
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
func finishSession(session *Session) {
	// Save the user's data
	session.SaveUser()
	// Prepare the messages, localize and send it
	sendOutputMessages(session)
}
