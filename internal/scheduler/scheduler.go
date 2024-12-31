//coverage:ignore file

package scheduler

import (
	"context"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/capymind/internal/database"
	"github.com/capymind/internal/taskservice"
)

var wg sync.WaitGroup

// Schedule a message for all users
func Schedule(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	log.Println("Schedule capymind...")

	// Parse the URL parameters
	typeStr, offset := parse(r.URL)
	if typeStr == nil {
		log.Println("Missing type parameter")
		return
	}
	// Get the message type
	messageType := taskservice.MessageType(*typeStr)
	// Get the message
	message := getTextMessage(messageType, time.Now().Weekday())
	if message == nil {
		log.Println("Missing message type parameter")
		return
	}

	// Create a context
	ctx := context.Background()
	// Create a Firestore client
	db.CreateClient(&ctx)
	// Create a Tasks client
	tasks.Connect(&ctx)

	var isCloud = false
	if os.Getenv("CLOUD") == "true" {
		isCloud = true
	}

	userStorage.ForEachUser(&ctx, func(users []database.User) error {
		for _, user := range users {
			// Skip users without locale or timezone
			if user.Locale == nil || user.SecondsFromUTC == nil {
				continue
			}
			// For debugging locally
			if !isCloud && !database.IsAdmin(user.Role) {
				continue
			}

			// Prepare message concurrently
			wg.Add(1)
			go prepareMessage(&user, &ctx, offset, messageType, *message, isCloud)
		}
		return nil
	})

	wg.Wait()

	// Close Firestore client
	tasks.Close()
	// Close Tasks client
	db.CloseClient()

	// Calculate how seconds this function takes to execute
	elapsed := time.Since(start)
	log.Printf("[Scheduler] Execution time for %s: %s", messageType, elapsed)
}
