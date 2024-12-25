package database

import "context"

type AdminStorage interface {
	GetTotalUserCount(ctx *context.Context) (int64, error)
	GetActiveUserCount(ctx *context.Context) (int64, error)
	GetTotalNoteCount(ctx *context.Context) (int64, error)
}
