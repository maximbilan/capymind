package app

import "github.com/capymind/internal/helpers"

func handleFeedbackLastWeek(session *Session) {
	array := helpers.PrepareFeedback(session.Context, session.Locale(), feedbackStorage)
	for _, item := range array {
		setOutputText(item, session)
	}
}
