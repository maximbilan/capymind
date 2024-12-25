package bot

import (
	"log"
	"time"

	"github.com/capymind/internal/database"
)

func startFeedback(session *Session) {
	setOutputText("start_feedback", session)
	session.User.IsTyping = true
}

func finishFeedback(session *Session) {
	text := *session.Job.Input
	saveFeedback(text, session)
	setOutputText("finish_feedback", session)
	session.User.IsTyping = false
}

func saveFeedback(text string, session *Session) {
	timestamp := time.Now()
	var feedback = database.Feedback{
		Text:      text,
		Timestamp: timestamp,
	}

	// Save the note
	err := feedbackStorage.NewFeedback(session.Context, *session.User, feedback)
	if err != nil {
		log.Printf("[Bot] Error saving feedback in firestore, %s", err.Error())
	}
}
