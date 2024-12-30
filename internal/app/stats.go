package app

import (
	"github.com/capymind/internal/database"
	"github.com/capymind/internal/helpers"
)

func handleStats(session *Session, adminStorage database.AdminStorage, feedbackStorage database.FeedbackStorage) {
	stats := helpers.GetStats(session.Context, session.Locale(), adminStorage, feedbackStorage)

	var finalString string
	for _, stat := range stats {
		finalString += stat + "\n"
	}
	setOutputText(finalString, session)
}
