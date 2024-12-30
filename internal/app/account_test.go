package app

import (
	"context"
	"os"
	"testing"

	"github.com/capymind/internal/database"
)

type NoteStorageMock struct{}

func (storage NoteStorageMock) NewNote(ctx *context.Context, user database.User, note database.Note) error {
	return nil
}

func (storage NoteStorageMock) LastNote(ctx *context.Context, userID string) (*database.Note, error) {
	note := database.Note{
		Text: "Test note",
	}
	return &note, nil
}

func (storage NoteStorageMock) GetNotesForLastWeek(ctx *context.Context, userID string) ([]database.Note, error) {
	note1 := database.Note{
		Text: "Test note",
	}
	note2 := database.Note{
		Text: "Test note 2",
	}
	notes := []database.Note{note1, note2}
	return notes, nil
}

func (storage NoteStorageMock) GetNotes(ctx *context.Context, userID string, count int) ([]database.Note, error) {
	note3 := database.Note{
		Text: "Test note 3",
	}
	note4 := database.Note{
		Text: "Test note 4",
	}
	notes := []database.Note{note3, note4}
	return notes, nil
}

func (storage NoteStorageMock) GetAllNotes(ctx *context.Context, userID string) ([]database.Note, error) {
	note5 := database.Note{
		Text: "Test note 5",
	}
	note6 := database.Note{
		Text: "Test note 6",
	}
	notes := []database.Note{note5, note6}
	return notes, nil
}

func (storage NoteStorageMock) NotesCount(ctx *context.Context, userID string) (int64, error) {
	return 10, nil
}

func (storage NoteStorageMock) DeleteAllNotes(ctx *context.Context, userID string) error {
	return nil
}

type UserStorageMock struct{}

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
