package app

import (
	"github.com/capymind/internal/database"
	"github.com/capymind/internal/helpers"
)

func handleTotalUserCount(session *Session, adminStorage database.AdminStorage) {
	message := helpers.GetTotalUserCount(session.Context, session.Locale(), adminStorage)
	if message != nil {
		setOutputText(*message, session)
	}
}

func handleTotalActiveUserCount(session *Session, adminStorage database.AdminStorage) {
	message := helpers.GetTotalActiveUserCount(session.Context, session.Locale(), adminStorage)
	if message != nil {
		setOutputText(*message, session)
	}
}

func handleTotalNoteCount(session *Session, adminStorage database.AdminStorage) {
	message := helpers.GetTotalNoteCount(session.Context, session.Locale(), adminStorage)
	if message != nil {
		setOutputText(*message, session)
	}
}
