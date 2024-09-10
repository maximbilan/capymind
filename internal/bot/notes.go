package bot

import (
	"log"
	"time"

	"github.com/capymind/internal/firestore"
)

func saveNote(text string, session Session) {
	// Setup the database connection
	client, ctx := createClient()
	defer client.Close()

	// Note data
	timestamp := time.Now()
	var note = firestore.Note{
		ID:        session.User.ID,
		Text:      text,
		Timestamp: timestamp,
	}

	// Save the note
	err := firestore.NewNote(ctx, client, session.User, note)
	if err != nil {
		log.Printf("[Note] Error saving note in firestore, %s", err.Error())
	}
}

func getNotes(session Session) []firestore.Note {
	client, ctx := createClient()
	defer client.Close()

	notes, err := firestore.GetNotes(ctx, client, session.User.ID)
	if err != nil {
		log.Printf("[Database] Error getting notes from firestore, %s", err.Error())
	}
	return notes
}
