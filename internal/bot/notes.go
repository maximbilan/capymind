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
	userID := session.User.ID
	note, err := firestore.LastNote(session.Context, userID)
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
	// Note data
	timestamp := time.Now()
	var note = firestore.Note{
		ID:        session.User.ID,
		Text:      text,
		Timestamp: timestamp,
	}

	// Save the note
	err := firestore.NewNote(session.Context, *session.User, note)
	if err != nil {
		log.Printf("[Bot] Error saving note in firestore, %s", err.Error())
	}
}

// Get the user's notes
func getNotes(session *Session) []firestore.Note {
	notes, err := firestore.GetNotes(session.Context, session.User.ID)
	if err != nil {
		log.Printf("[Bot] Error getting notes from firestore, %s", err.Error())
	}
	return notes
}
