package telegram

import "fmt"

type CallbackQuery struct {
	ID      string   `json:"id"`
	From    *User    `json:"from"`
	Message *Message `json:"message,omitempty"`
	Data    string   `json:"data,omitempty"`
}

func (query CallbackQuery) UserID() *string {
	if query.From != nil {
		id := fmt.Sprintf("%d", query.From.ID)
		return &id
	}
	return nil
}

func (query CallbackQuery) ChatID() *int64 {
	if query.Message != nil && query.Message.Chat != nil {
		return &query.Message.Chat.ID
	}
	return nil
}
