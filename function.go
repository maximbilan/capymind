package capymind

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/capymind/telegram"
)

func init() {
	functions.HTTP("handler", Handler)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var update, err = telegram.Parse(r)
	if err != nil {
		log.Printf("error parsing update, %s", err.Error())
		return
	}

	fmt.Println("Server received message: ", update.Message.Text)

	var telegramResponseBody, errTelegram = telegram.SendMessage(update.Message.Chat.Id, "Now I can respond to you!")
	if errTelegram != nil {
		log.Printf("got error %s from telegram, reponse body is %s", errTelegram.Error(), telegramResponseBody)
	}
}
