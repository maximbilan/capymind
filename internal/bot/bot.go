package bot

import (
	"log"
	"net/http"

	"github.com/capymind/internal/telegram"
)

func Parse(w http.ResponseWriter, r *http.Request) {
	update := telegram.Parse(r)
	if update == nil {
		log.Printf("[Bot] No update to process")
		return
	}

	// Create a user
	user := createUser(*update)
	if user == nil {
		log.Printf("[Bot] No user to process: message_id=%d", update.Message.ID)
		return
	}

	// Update the user's data in the database if necessary
	updatedUser := updateUser(user)

	// Create a job
	job := createJob(*update)
	if job == nil {
		log.Printf("[Bot] No job to process: update_id=%d", update.ID)
		return
	}

	// Create and start a session
	session := createSession(job, updatedUser)
	// Execute the job
	handleSession(session)
	// Send the response
	finishSession(session)
}
