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

// Create a session
func createSession(job Job, user firestore.User) Session {
	session := Session{
		Job:  job,
		User: user,
	}
	return session
}

// Start the session
func (session Session) Start() {
	session.Job.execute()
}
