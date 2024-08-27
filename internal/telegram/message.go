package telegram

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
	From User   `json:"from"`
	Date int    `json:"date"`
}

type SendMessageRequest struct {
	ChatID      int                   `json:"chat_id"`
	Text        string                `json:"text"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}
