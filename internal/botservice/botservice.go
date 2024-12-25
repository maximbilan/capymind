package botservice

import "net/http"

type BotService interface {
	Parse(r *http.Request) *string
}
