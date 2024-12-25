package database

import (
	"context"
	"time"
)

type Note struct {
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

type NoteStorage interface {
	NewNote(ctx *context.Context, user User, note Note) error
	LastNote(ctx *context.Context, userID string) (*Note, error)
	GetNotesForLastWeek(ctx *context.Context, userID string) ([]Note, error)
	GetNotes(ctx *context.Context, userID string, count int) ([]Note, error)
	NotesCount(ctx *context.Context, userID string) (int64, error)
}
