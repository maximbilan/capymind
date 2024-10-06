package bot

import (
	"log"

	"github.com/capymind/internal/firestore"
)

func handleFeedbackLastWeek(session *Session) {
	setOutputText("feedback_last_week", session)
	setOutputText("\n\n", session)

	feedback, err := firestore.GetFeedbackForLastWeek(session.Context)
	if err != nil {
		log.Printf("[Bot] Error getting feedbacks from firestore, %s", err.Error())
	}

	if len(feedback) == 0 {
		setOutputText("no_feedback", session)
		return
	}

	for _, f := range feedback {
		setOutputText(*f.User.FirstName+" "+*f.User.LastName+":"+"\n"+f.Feedback.Text+"\n\n", session)
	}
}
