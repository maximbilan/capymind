package bot

import (
	"context"
	"time"

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

func (session *Session) isAdmin() bool {
	return isAdmin(session.User)
}

// Handle the session
func handleSession(session *Session) {
	now := time.Now()
	command := session.Job.Command
	commandStr := string(command)
	session.User.LastCommand = &commandStr
	session.User.Timestamp = &now

	if command.IsAdmin() && !session.isAdmin() {
		handleHelp(session)
		return
	}

	switch command {
	case Start:
		handleStart(session)
	case Note:
		startNote(session)
	case Last:
		handleLastNote(session)
	case Analysis:
		handleAnalysis(session)
	case Settings:
		handleSettings(session)
	case Language:
		handleLanguage(session)
	case Timezone:
		handleTimezone(session)
	case Support:
		startFeedback(session)
	case Help:
		handleHelp(session)
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
	case FeedbackLastWeek:
		handleFeedbackLastWeek(session)
	case None:
		// Typing mode
		if session.User.IsTyping && session.Job.Input != nil {
			switch session.Job.LastCommand {
			case Note:
				finishNote(session)
			case Support:
				finishFeedback(session)
			default:
				// If the user is typing and the last command is not recognized
				handleHelp(session)
			}
		} else {
			// Otherwise show the help message
			handleHelp(session)
		}
	default:
		// Unknown command
		handleHelp(session)
	}
}

// Finish the session. Send the output to the user
func finishSession(session *Session) {
	// Save the user's data
	session.SaveUser()
	// Prepare the messages, localize and send it
	sendOutputMessages(session)
}
