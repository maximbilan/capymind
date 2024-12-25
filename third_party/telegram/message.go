package telegram

type Message struct {
	ID   int64  `json:"message_id"`
	Text string `json:"text,omitempty"`
	Chat *Chat  `json:"chat"`
	From *User  `json:"from,omitempty"`
	Date int    `json:"date"`
}
