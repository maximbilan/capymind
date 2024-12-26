package app

import (
	"archive/zip"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/capymind/internal/database"
)

func handleDownloadData(session *Session) {
	sendMessage("download_all_notes_waiting", session)

	userID := session.User.ID
	notes, err := noteStorage.GetAllNotes(session.Context, userID)
	if err != nil {
		log.Printf("[Bot] Error getting all notes from firestore, %s", err.Error())
		setOutputText("download_all_notes_error", session)
		return
	}

	if len(notes) == 0 {
		setOutputText("download_all_notes_empty", session)
		return
	}

	// create a ZIP file with all notes
	zipFile, err := createZipFile(notes)
	if err != nil {
		log.Printf("[Bot] Error creating ZIP file, %s", err.Error())
		setOutputText("download_all_notes_error", session)
		return
	}
	if zipFile != nil {
		// Upload the ZIP file to Google Drive
	} else {
		setOutputText("download_all_notes_error", session)
	}
}

func createZipFile(notes []database.Note) (*os.File, error) {
	zipFileName := fmt.Sprintf("notes_%s.zip", time.Now().Format("2006-01-02_15-04-05"))
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return nil, err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("notes.txt")
	if err != nil {
		return nil, err
	}

	for _, note := range notes {
		writer.Write([]byte(note.Timestamp.Format("2006-01-02 15:04:05") + "\n"))
		writer.Write([]byte(note.Text + "\n\n"))
	}
	return zipFile, nil
}
