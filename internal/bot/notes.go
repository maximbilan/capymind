package bot

import (
	"log"
	"time"

	"github.com/capymind/internal/firestore"
	"github.com/capymind/internal/translator"
)

// Start typing a note
func startNote(session *Session) {
	setOutputText("start_note", session)
	session.User.IsTyping = true
}

// Finish typing a note
func finishNote(session *Session) {
	saveNote(*session.Job.Input, session)
	setOutputText("finish_note", session)
	session.User.IsTyping = false
}

// Handle the note command
func handleLastNote(session *Session) {
	client, ctx := createClient()
	defer client.Close()

	userID := session.User.ID
	note, err := firestore.LastNote(ctx, client, userID)
	if err != nil {
		log.Printf("[Bot] Error getting last note from firestore, %s", err.Error())
	}

	if note != nil {
		var response string = translator.Translate(session.Locale(), "your_last_note") + note.Text
		setOutputText(response, session)
	} else {
		var button JobResultTextButton = JobResultTextButton{
			TextID:   "make_record_to_journal",
			Callback: string(Note),
		}
		setOutputTextWithButtons("no_notes", []JobResultTextButton{button}, session)
	}
}

// Save a note
func saveNote(text string, session *Session) {
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
	err := firestore.NewNote(ctx, client, *session.User, note)
	if err != nil {
		log.Printf("[Bot] Error saving note in firestore, %s", err.Error())
	}
}

// Get the user's notes
func getNotes(session *Session) []firestore.Note {
	client, ctx := createClient()
	defer client.Close()

	notes, err := firestore.GetNotes(ctx, client, session.User.ID)
	if err != nil {
		log.Printf("[Bot] Error getting notes from firestore, %s", err.Error())
	}
	return notes
}
