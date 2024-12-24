package bot

import (
	"fmt"
	"log"
	"sync"

	"github.com/capymind/internal/firestore"
	"github.com/capymind/internal/translator"
)

type statFunc func(session *Session) *string

var wg sync.WaitGroup

func getTotalUserCount(session *Session) *string {
	count, err := firestore.GetTotalUserCount(session.Context)
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
	count, err := firestore.GetActiveUserCount(session.Context)
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
	count, err := firestore.GetTotalNoteCount(session.Context)
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

func handleStats(session *Session) {
	totalUserCount := waitForStatFunction(getTotalUserCount, session)
	totalActiveUserCount := waitForStatFunction(getTotalActiveUserCount, session)
	totalNoteCount := waitForStatFunction(getTotalNoteCount, session)

	wg.Wait()

	if totalUserCount != nil {
		setOutputText(*totalUserCount, session)
	}
	if totalActiveUserCount != nil {
		setOutputText(*totalActiveUserCount, session)
	}
	if totalNoteCount != nil {
		setOutputText(*totalNoteCount, session)
	}
}

func waitForStatFunction(statFunc statFunc, session *Session) *string {
	wg.Add(1)
	ch := make(chan *string)
	go func() {
		defer wg.Done()
		result := statFunc(session)
		ch <- result
	}()
	result := <-ch
	return result
}
