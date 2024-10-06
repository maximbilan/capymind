package firestore

import "context"

type Update struct {
	Text         string
	HasDelivered bool
}

func NewUpdate(ctx *context.Context, update Update) error {
	_, _, err := client.Collection(updates.String()).Add(*ctx, map[string]interface{}{
		"text":         update.Text,
		"hasDelivered": update.HasDelivered,
	})
	return err
}
