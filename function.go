package capymind

import (
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/capymind/internal/bot"
)

func init() {
	functions.HTTP("handler", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	bot.Parse(w, r)
}
