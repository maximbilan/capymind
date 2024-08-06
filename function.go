package capymind

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/capymind/firestore"
	"github.com/capymind/localizer"
	"github.com/capymind/telegram"
)

type Command string

const (
	Start Command = "start"
	Note  Command = "note"
	Info  Command = "info"
)

func init() {
	functions.HTTP("handler", Handler)

	err := localizer.Load("./localizer/translations.json")
	if err != nil {
		log.Fatalf("Failed to load translations: %v", err)
	}

	fmt.Println(localizer.Translate("en", "welcome"))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var update, err = telegram.Parse(r)
	if err != nil {
		log.Printf("Error parsing update, %s", err.Error())
		return
	}

	text := update.Message.Text
	command := Command(text)
	message := update.Message

	switch command {
	case Start:
		handleStart(message)
	case Note:
		handleNote(message)
	case Info:
		handleInfo(message)
	default:
		handleUnknownState(message)
	}
	ctx := context.Background()
	saveNote(ctx, update.Message)

	var telegramResponseBody, errTelegram = telegram.SendMessage(update.Message.Chat.Id, "Now I can respond to you!")
	if errTelegram != nil {
		log.Printf("got error %s from telegram, reponse body is %s", errTelegram.Error(), telegramResponseBody)
	}
}

func handleStart(message telegram.Message) {
}

func handleNote(telegram.Message) {
}

func handleUnknownState(message telegram.Message) {
}

func handleInfo(message telegram.Message) {

}

func saveNote(ctx context.Context, message telegram.Message) {
	fmt.Println("Saving note to firestore...")

	var client, err = firestore.NewClient(ctx)
	if err != nil {
		log.Printf("Error creating firestore client, %s", err.Error())
	}
	defer client.Close()

	var user = firestore.User{
		ID:   fmt.Sprintf("%d", message.Chat.Id),
		Name: message.From.Username,
	}

	timestamp := time.Now()
	var note = firestore.Note{
		ID:        fmt.Sprintf("%d", message.Chat.Id),
		Text:      message.Text,
		Timestamp: timestamp,
	}

	err = firestore.NewRecord(ctx, client, user, note)
	if err != nil {
		log.Printf("Error saving note to firestore, %s", err.Error())
	}
}
