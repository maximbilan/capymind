package scheduler

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/capymind/internal/taskservice"
)

// Returns the type and offset parameters from the URL
func parse(url *url.URL) (*string, int) {
	typeStr := url.Query().Get("type")
	offsetStr := url.Query().Get("offset") // hours (from UTC 0)
	var offset int = 0
	if offsetStr != "" {
		_, err := fmt.Sscanf(offsetStr, "%d", &offset)
		if err != nil {
			log.Printf("[Scheduler] Error getting offset parameter, %s", err.Error())
		}
	}
	return &typeStr, offset
}

func getTextMessage(messageType taskservice.MessageType) *string {
	var message string
	switch messageType {
	case taskservice.Morning, taskservice.Evening:
		message = taskservice.GetMessage(messageType, time.Now().Weekday())
	case taskservice.Feedback:
		message = "ask_write_review_about_bot"
	case taskservice.WeeklyAnalysis, taskservice.UserStats, taskservice.AdminStats:
		// Personalized for each user
		message = ""
	default:
		log.Println("Missing message type parameter")
		return nil
	}
	return &message
}
