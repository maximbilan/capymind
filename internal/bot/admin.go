package bot

import (
	"fmt"
	"log"

	"github.com/capymind/internal/firestore"
	"github.com/capymind/internal/translator"
)

func handleTotalUserCount(session *Session) {
	count, err := firestore.GetTotalUserCount(session.Context)
	if err != nil {
		log.Printf("[Admin] Error during fetching total user count: %v", err)
		return
	}

	message := fmt.Sprintf(translator.Translate(session.Locale(), "total_user_count"), count)
	setOutputText(message, session)
}

func handleTotalNoteCount(session *Session) {
	count, err := firestore.GetTotalNoteCount(session.Context)
	if err != nil {
		log.Printf("[Admin] Error during fetching total note count: %v", err)
		return
	}

	message := fmt.Sprintf(translator.Translate(session.Locale(), "total_note_count"), count)
	setOutputText(message, session)
}
