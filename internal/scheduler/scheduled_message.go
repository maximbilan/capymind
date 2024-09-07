package scheduler

import "github.com/capymind/internal/translator"

type ScheduledMessage struct {
	ChatId int               `json:"chatId"`
	Text   string            `json:"text"`
	Type   MessageType       `json:"type"`
	Locale translator.Locale `json:"locale"`
}
