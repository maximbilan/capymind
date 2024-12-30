//coverage:ignore file

package app

import (
	"context"
	"log"
	"net/http"
)

// Entry point
func Parse(w http.ResponseWriter, r *http.Request) {
	update := bot.Parse(r.Body)
	if update == nil {
		log.Printf("[Bot] No update to process")
		return
	}

	// Create a context
	ctx := context.Background()

	// Creat a database connection
	db.CreateClient(&ctx)

	// Create a user
	user := createUser(*update)

	// Update the user's data in the database if necessary
	updatedUser := updateUser(user, &ctx, userStorage)

	// Create a job
	job := createJob(update.Text, updatedUser)

	// Create and start a session
	session := createSession(job, updatedUser, &ctx)
	// Execute the job
	handleSession(session)
	// Send the response
	finishSession(session)
	// Close the database connection
	db.CloseClient()
}
