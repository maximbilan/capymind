package firestore

import (
	"context"

	"github.com/capymind/internal/database"
)

type UpdateStorage struct{}

func (storage UpdateStorage) NewUpdate(ctx *context.Context, update database.Update) error {
	_, _, err := client.Collection(database.Updates.String()).Add(*ctx, map[string]interface{}{
		"text":         update.Text,
		"hasDelivered": update.HasDelivered,
	})
	return err
}
