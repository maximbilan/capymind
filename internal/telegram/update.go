package telegram

type Update struct {
	UpdateId      int            `json:"update_id"`
	Message       Message        `json:"message"`
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
}
