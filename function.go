package capymind

import (
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/capymind/internal/bot"
	"github.com/capymind/internal/scheduler"
)

func init() {
	log.Printf("Version: %s", AppVersion)

	functions.HTTP("handler", handler)
	functions.HTTP("schedule", schedule)
	functions.HTTP("sendMessage", sendMessage)
}

func handler(w http.ResponseWriter, r *http.Request) {
	bot.Parse(w, r)
}

func schedule(w http.ResponseWriter, r *http.Request) {
	scheduler.Schedule(w, r)
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	scheduler.SendMessage(w, r)
}
