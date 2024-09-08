package telegram

type Message struct {
	MessageID int    `json:"message_id"`
	Text      string `json:"text,omitempty"`
	Chat      *Chat  `json:"chat"`
	From      *User  `json:"from,omitempty"`
	Date      int    `json:"date"`
}

type SendMessageRequest struct {
	ChatID      int64                 `json:"chat_id"`
	Text        string                `json:"text"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}
