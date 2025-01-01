package database

import "context"

type Settings struct {
	SecondsFromUTC        *int    `json:"secondsFromUTC"`
	HasMorningReminder    *bool   `json:"hasMorningReminder"`
	HasEveningReminder    *bool   `json:"hasEveningReminder"`
	MorningReminderOffset *int    `json:"morningReminderOffset"`
	EveningReminderOffset *int    `json:"eveningReminderOffset"`
	Location              *string `json:"location"`
}

type SettingsStorage interface {
	GetSettings(ctx *context.Context, userID string) (*Settings, error)
	SaveSettings(ctx *context.Context, userID string, settings Settings) error
	DeleteSettings(ctx *context.Context, userID string) error
}

func (s Settings) IsMorningReminderEnabled() bool {
	if s.HasMorningReminder == nil {
		return true
	}
	return *s.HasMorningReminder
}

func (s Settings) IsEveningReminderEnabled() bool {
	if s.HasEveningReminder == nil {
		return true
	}
	return *s.HasEveningReminder
}

func (s Settings) AreRemindersEnabled() bool {
	return s.IsMorningReminderEnabled() && s.IsEveningReminderEnabled()
}
