package telegram

type Update struct {
	UpdateId      int            `json:"update_id"`
	Message       Message        `json:"message"`
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
}

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
	From User   `json:"from"`
	Date int    `json:"date"`
}

type SendMessageRequest struct {
	ChatID      int                  `json:"chat_id"`
	Text        string               `json:"text"`
	ReplyMarkup InlineKeyboardMarkup `json:"reply_markup"`
}

type Chat struct {
	Id int `json:"id"`
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data,omitempty"`
}

type CallbackQuery struct {
	ID      string   `json:"id"`
	From    *User    `json:"from"`
	Message *Message `json:"message,omitempty"`
	Data    string   `json:"data,omitempty"`
}
