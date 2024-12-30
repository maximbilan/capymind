//coverage:ignore file

package mocks

import (
	"context"

	"github.com/capymind/internal/database"
)

type EmptyNoteStorageMock struct{}

func (storage EmptyNoteStorageMock) NewNote(ctx *context.Context, user database.User, note database.Note) error {
	return nil
}

func (storage EmptyNoteStorageMock) LastNote(ctx *context.Context, userID string) (*database.Note, error) {
	return nil, nil
}

func (storage EmptyNoteStorageMock) GetNotesForLastWeek(ctx *context.Context, userID string) ([]database.Note, error) {
	return []database.Note{}, nil
}

func (storage EmptyNoteStorageMock) GetNotes(ctx *context.Context, userID string, count int) ([]database.Note, error) {
	return []database.Note{}, nil
}

func (storage EmptyNoteStorageMock) GetAllNotes(ctx *context.Context, userID string) ([]database.Note, error) {
	return []database.Note{}, nil
}

func (storage EmptyNoteStorageMock) NotesCount(ctx *context.Context, userID string) (int64, error) {
	return 0, nil
}

func (storage EmptyNoteStorageMock) DeleteAllNotes(ctx *context.Context, userID string) error {
	return nil
}
