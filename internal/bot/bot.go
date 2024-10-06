package bot

import (
	"context"
	"log"
	"net/http"

	"github.com/capymind/internal/firestore"
	"github.com/capymind/internal/telegram"
)

// Entry point
func Parse(w http.ResponseWriter, r *http.Request) {
	update := telegram.Parse(r)
	if update == nil {
		log.Printf("[Bot] No update to process")
		return
	}

	// Create a context
	ctx := context.Background()

	// Creat a database connection
	firestore.CreateClient(&ctx)

	// Create a user
	user := createUser(*update)
	if user == nil {
		log.Printf("[Bot] No user to process: message_id=%d", update.ID)
		return
	}

	// Update the user's data in the database if necessary
	updatedUser := updateUser(user, &ctx)

	// Create a job
	job := createJob(*update, updatedUser)
	if job == nil {
		log.Printf("[Bot] No job to process: update_id=%d", update.ID)
		return
	}

	// Create and start a session
	session := createSession(job, updatedUser, &ctx)
	// Execute the job
	handleSession(session)
	// Send the response
	finishSession(session)
	// Close the database connection
	firestore.CloseClient()
}
