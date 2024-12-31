//coverage:ignore file

package mocks

import (
	"context"

	"github.com/capymind/internal/database"
)

type EmptySettingsStorageMock struct{}

func (s EmptySettingsStorageMock) GetSettings(ctx *context.Context, userID string) (*database.Settings, error) {
	return &database.Settings{}, nil
}

func (s EmptySettingsStorageMock) SaveSettings(ctx *context.Context, userID string, settings database.Settings) error {
	return nil
}

func (s EmptySettingsStorageMock) DeleteSettings(ctx *context.Context, userID string) error {
	return nil
}
