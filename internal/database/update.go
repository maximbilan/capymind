package database

import "context"

type Update struct {
	Text         string
	HasDelivered bool
}

type UpdateStorage interface {
	NewUpdate(ctx *context.Context, update Update) error
}
