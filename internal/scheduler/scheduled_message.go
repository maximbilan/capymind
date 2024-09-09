package scheduler

import "github.com/capymind/internal/translator"

type ScheduledMessage struct {
	ChatID int64             `json:"chatId"`
	Text   string            `json:"text"`
	Type   MessageType       `json:"type"`
	Locale translator.Locale `json:"locale"`
}
