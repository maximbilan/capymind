package telegram

import "fmt"

type Message struct {
	ID   int64  `json:"message_id"`
	Text string `json:"text,omitempty"`
	Chat *Chat  `json:"chat"`
	From *User  `json:"from,omitempty"`
	Date int    `json:"date"`
}

func (message *Message) UserID() string {
	return fmt.Sprintf("%d", message.From.ID)
}

func (message *Message) ChatID() int64 {
	return message.Chat.ID
}
