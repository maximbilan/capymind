package botservice

import (
	"net/http"

	"github.com/capymind/internal/translator"
)

type BotService interface {
	Parse(r *http.Request) *BotMessage
	SendMessage(chatID int64, text string)
	SendResult(chatID int64, locale translator.Locale, result BotResult)
}
