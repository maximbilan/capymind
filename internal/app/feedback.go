package app

import (
	"github.com/capymind/internal/database"
	"github.com/capymind/internal/helpers"
)

func handleFeedbackLastWeek(session *Session, feedbackStorage database.FeedbackStorage) {
	array := helpers.PrepareFeedback(session.Context, session.Locale(), feedbackStorage)
	for _, item := range array {
		setOutputText(item, session)
	}
}
