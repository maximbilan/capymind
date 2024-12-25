package botservice

import (
	"net/http"
)

type BotService interface {
	Parse(r *http.Request) *BotMessage
	SendMessage(chatID int64, text string)
	SendResult(chatID int64, result BotResult)
}
