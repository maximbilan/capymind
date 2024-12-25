package app

import (
	"fmt"
	"log"

	"github.com/capymind/internal/translator"
)

func getTotalUserCount(session *Session) *string {
	count, err := adminStorage.GetTotalUserCount(session.Context)
	if err != nil {
		log.Printf("[Admin] Error during fetching total user count: %v", err)
		return nil
	}
	message := fmt.Sprintf(translator.Translate(session.Locale(), "total_user_count"), count)
	return &message
}

func handleTotalUserCount(session *Session) {
	message := getTotalUserCount(session)
	if message != nil {
		setOutputText(*message, session)
	}
}

func getTotalActiveUserCount(session *Session) *string {
	count, err := adminStorage.GetActiveUserCount(session.Context)
	if err != nil {
		log.Printf("[Admin] Error during fetching active user count: %v", err)
		return nil
	}
	message := fmt.Sprintf(translator.Translate(session.Locale(), "total_active_user_count"), count)
	return &message
}

func handleTotalActiveUserCount(session *Session) {
	message := getTotalActiveUserCount(session)
	if message != nil {
		setOutputText(*message, session)
	}
}

func getTotalNoteCount(session *Session) *string {
	count, err := adminStorage.GetTotalNoteCount(session.Context)
	if err != nil {
		log.Printf("[Admin] Error during fetching total note count: %v", err)
		return nil
	}
	message := fmt.Sprintf(translator.Translate(session.Locale(), "total_note_count"), count)
	return &message
}

func handleTotalNoteCount(session *Session) {
	message := getTotalNoteCount(session)
	if message != nil {
		setOutputText(*message, session)
	}
}
