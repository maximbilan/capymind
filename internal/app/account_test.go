package app

import (
	"os"
	"testing"

	"github.com/capymind/internal/database"
)

func TestCreateZipFile(t *testing.T) {
	userID := "123"

	note1 := database.Note{
		Text: "Test note 1",
	}
	note2 := database.Note{
		Text: "Test note 2",
	}
	notes := []database.Note{note1, note2}

	zipFile, err := createZipFile(userID, notes)

	if err != nil {
		t.Error("Expected nil error, got", err)
	}

	if zipFile == nil {
		t.Error("Expected zip file, got nil")
	}
	// Check if the file name starts with notes_123
	if zipFile.Name()[:9] != "notes_123" {
		t.Error("Expected file name notes_123, got", zipFile.Name())
	}

	os.Remove(zipFile.Name())
	zipFile.Close()
}
