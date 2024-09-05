package scheduler

type ScheduledMessage struct {
	ChatId int    `json:"chatId"`
	Text   string `json:"text"`
}
