package app

import (
	"archive/zip"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/capymind/internal/database"
	"github.com/capymind/third_party/googledrive"
)

var fileStorage googledrive.GoogleDrive

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
	zipFile, err := createZipFile(userID, notes)
	if err != nil {
		log.Printf("[Bot] Error creating ZIP file, %s", err.Error())
		setOutputText("download_all_notes_error", session)
		return
	}
	if zipFile != nil {
		// Upload the ZIP file to Google Drive
		title := fmt.Sprintf("Notes %s", userID)
		link := fileStorage.Upload(title, zipFile.Name(), time.Now().Add(7*24*time.Hour))
		if link != nil {
			setOutputText(*link, session)
		} else {
			setOutputText("download_all_notes_error", session)
		}
		// Remove the ZIP file after upload and close it
		os.Remove(zipFile.Name())
		zipFile.Close()
	} else {
		setOutputText("download_all_notes_error", session)
	}
}

// Create a ZIP file with all notes
// Attention: zipFile must be closed after use
func createZipFile(userID string, notes []database.Note) (*os.File, error) {
	zipFileName := fmt.Sprintf("notes_%s_%s.zip", userID, time.Now().Format("2006-01-02_15-04-05"))
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return nil, err
	}

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