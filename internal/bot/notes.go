package bot

import (
	"fmt"
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
	text := *session.Job.Input
	saveNote(text, session)
	setOutputText("finish_note", session)

	isDream := checkIfNoteADream(text, session.Locale())
	if isDream {
		askForSleepAnalysis(session)
	}

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

		isDream := checkIfNoteADream(note.Text, session.Locale())
		if isDream {
			askForSleepAnalysis(session)
		}
	} else {
		sendNoNotes(session)
	}
}

// Save a note
func saveNote(text string, session *Session) {
	// Note data
	timestamp := time.Now()
	var note = firestore.Note{
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
func getNotes(session *Session, count int) []firestore.Note {
	notes, err := firestore.GetNotes(session.Context, session.User.ID, count)
	if err != nil {
		log.Printf("[Bot] Error getting notes from firestore, %s", err.Error())
	}
	return notes
}

// Send a message that says there are no notes
func sendNoNotes(session *Session) {
	var button JobResultTextButton = JobResultTextButton{
		TextID:   "make_record_to_journal",
		Callback: string(Note),
	}
	setOutputTextWithButtons("no_notes", []JobResultTextButton{button}, session)
}

// Handles the note count request
func handleNoteCount(session *Session) {
	count, err := firestore.NotesCount(session.Context, session.User.ID)
	if err != nil {
		log.Printf("[Bot] Error getting notes count from firestore, %s", err.Error())
	} else {
		message := fmt.Sprintf(translator.Translate(session.Locale(), "user_progress_message"), count)
		setOutputText(message, session)
	}
}
