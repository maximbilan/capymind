//coverage:ignore file

package firestore

import (
	"context"

	"github.com/capymind/internal/database"
)

type SettingsStorage struct{}

func (s SettingsStorage) GetSettings(ctx *context.Context, userID string) (*database.Settings, error) {
	doc, err := client.Collection(database.SettingsCollection.String()).Doc(userID).Get(*ctx)
	if err != nil {
		return nil, err
	}

	var settings database.Settings
	doc.DataTo(&settings)
	return &settings, nil
}

func (s SettingsStorage) SaveSettings(ctx *context.Context, userID string, settings database.Settings) error {
	_, err := client.Collection(database.SettingsCollection.String()).Doc(userID).Set(*ctx, settings)
	return err
}

func (s SettingsStorage) DeleteSettings(ctx *context.Context, userID string) error {
	_, err := client.Collection(database.SettingsCollection.String()).Doc(userID).Delete(*ctx)
	return err
}
