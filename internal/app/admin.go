package app

import "github.com/capymind/internal/helpers"

func handleTotalUserCount(session *Session) {
	message := helpers.GetTotalUserCount(session.Context, session.Locale())
	if message != nil {
		setOutputText(*message, session)
	}
}

func handleTotalActiveUserCount(session *Session) {
	message := helpers.GetTotalActiveUserCount(session.Context, session.Locale())
	if message != nil {
		setOutputText(*message, session)
	}
}

func handleTotalNoteCount(session *Session) {
	message := helpers.GetTotalNoteCount(session.Context, session.Locale())
	if message != nil {
		setOutputText(*message, session)
	}
}
