//coverage:ignore file

package mocks

import (
	"context"

	"github.com/capymind/internal/database"
)

type SettingsStorageMock struct{}

func (s SettingsStorageMock) GetSettings(ctx *context.Context, userID string) (*database.Settings, error) {
	enabler := true
	offset := 7200
	settings := database.Settings{
		SecondsFromUTC:        &offset,
		HasMorningReminder:    &enabler,
		HasEveningReminder:    &enabler,
		MorningReminderOffset: &offset,
		EveningReminderOffset: &offset,
	}
	return &settings, nil
}

func (s SettingsStorageMock) SaveSettings(ctx *context.Context, userID string, settings database.Settings) error {
	return nil
}

func (s SettingsStorageMock) DeleteSettings(ctx *context.Context, userID string) error {
	return nil
}
