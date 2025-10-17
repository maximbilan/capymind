package app

import (
    "fmt"
    "log"
    "time"

    "github.com/capymind/internal/botservice"
    "github.com/capymind/internal/database"
    "github.com/capymind/internal/translator"
)

// Start typing a note
func startNote(session *Session) {
	setOutputText("start_note", session)
	session.User.IsTyping = true
}

// Finish typing a note
func finishNote(text string, session *Session, noteStorage database.NoteStorage) {
	saveNote(text, session, noteStorage)
	setOutputText("finish_note", session)

	isDream := checkIfNoteADream(text, session.Locale())
	if isDream {
		askForSleepAnalysis(session)
	}

	session.User.IsTyping = false
}

func handleMissingNote(session *Session, noteStorage database.NoteStorage) {
	text := *session.Job.Input
	// remove the command from the text at the beginning
	text = text[len(MissingNote):]

	finishNote(text, session, noteStorage)
}

func handleInputWithoutCommand(session *Session) {
	var saveButton botservice.BotResultTextButton = botservice.BotResultTextButton{
		TextID:   "missing_note_confirm",
		Locale:   session.Locale(),
		Callback: fmt.Sprintf("%s %s", string(MissingNote), *session.Job.Input),
	}
	setOutputTextWithButtons("missing_note", []botservice.BotResultTextButton{saveButton}, session)
}

// Handle the note command
func handleLastNote(session *Session, noteStorage database.NoteStorage) {
	userID := session.User.ID
	note, err := noteStorage.LastNote(session.Context, userID)
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
func saveNote(text string, session *Session, noteStorage database.NoteStorage) {
	// Note data
	timestamp := time.Now()
	var note = database.Note{
		Text:      text,
		Timestamp: timestamp,
	}

	// Save the note
	err := noteStorage.NewNote(session.Context, *session.User, note)
	if err != nil {
		log.Printf("[Bot] Error saving note in firestore, %s", err.Error())
	}
}

// Get the user's notes
func getNotes(session *Session, noteStorage database.NoteStorage, count int) []database.Note {
	notes, err := noteStorage.GetNotes(session.Context, session.User.ID, count)
	if err != nil {
		log.Printf("[Bot] Error getting notes from firestore, %s", err.Error())
	}
	return notes
}

// moved therapy-related functions to therapy.go

// Send a message that says there are no notes
func sendNoNotes(session *Session) {
	var button botservice.BotResultTextButton = botservice.BotResultTextButton{
		TextID:   "make_record_to_journal",
		Locale:   session.Locale(),
		Callback: string(Note),
	}
	setOutputTextWithButtons("no_notes", []botservice.BotResultTextButton{button}, session)
}

// Handles the note count request
func handleNoteCount(session *Session, noteStorage database.NoteStorage) {
	count, err := noteStorage.NotesCount(session.Context, session.User.ID)
	if err != nil {
		log.Printf("[Bot] Error getting notes count from firestore, %s", err.Error())
	} else {
		message := fmt.Sprintf(translator.Translate(session.Locale(), "user_progress_message"), count)
		setOutputText(message, session)
	}
}
