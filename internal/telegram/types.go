package telegram

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
	From User   `json:"from"`
	Date int    `json:"date"`
}

type Chat struct {
	Id int `json:"id"`
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}
