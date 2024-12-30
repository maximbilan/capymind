package app

import (
	"os"
	"testing"

	"github.com/capymind/internal/database"
	"github.com/capymind/internal/mocks"
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

func TestDownloadDataHandler(t *testing.T) {
	session := createSession(&Job{Command: "/download"}, &database.User{}, nil)
	noteStorage := mocks.NoteStorageMock{}
	fileStorage := mocks.ValidFileStorageMock{}

	handleDownloadData(session, noteStorage, fileStorage)

	if session.Job.Output[0].TextID != "link" {
		t.Error("Expected link, got nil")
	}
}

func TestDeleteAccountHandler(t *testing.T) {
	session := createSession(&Job{Command: "/delete"}, &database.User{}, nil)
	handleDeleteAccount(session)

	if session.Job.Output[0].TextID != "delete_account_are_you_sure" {
		t.Error("Expected delete_account_confirm, got nil")
	}
	if session.Job.Output[0].Buttons[0].TextID != "delete_account_confirm" {
		t.Error("Expected delete_account_are_you_sure, got nil")
	}
}

func TestForceDeleteAccountHandler(t *testing.T) {
	session := createSession(&Job{Command: "/force_delete"}, &database.User{}, nil)
	userStorage := mocks.UserStorageMock{}
	noteStorage := mocks.NoteStorageMock{}

	handleForceDeleteAccount(session, noteStorage, userStorage)

	if session.Job.Output[0].TextID != "delete_account_success" {
		t.Error("Expected delete_account_success, got nil")
	}
	if session.Job.Output[1].TextID != "delete_account_telegram_tip" {
		t.Error("Expected delete_account_telegram_tip, got nil")
	}
}
