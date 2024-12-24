package bot

import (
	"log"

	"github.com/capymind/internal/firestore"
)

func prepareFeedback(session *Session) []string {
	var array []string
	array = append(array, "feedback_last_week")
	array = append(array, "\n\n")

	feedback, err := firestore.GetFeedbackForLastWeek(session.Context)
	if err != nil {
		log.Printf("[Bot] Error getting feedbacks from firestore, %s", err.Error())
	}

	if len(feedback) == 0 {
		array = append(array, "no_feedback")
		return array
	}

	for _, f := range feedback {
		array = append(array, *f.User.FirstName+" "+*f.User.LastName+":"+"\n"+f.Feedback.Text+"\n\n")
	}

	return array
}

func handleFeedbackLastWeek(session *Session) {
	array := prepareFeedback(session)
	for _, item := range array {
		setOutputText(item, session)
	}
}
