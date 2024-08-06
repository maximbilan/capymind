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
	Start Command = "/start"
	Note  Command = "/note"
	Info  Command = "/info"
)

func init() {
	functions.HTTP("handler", Handler)

	err := localizer.Load("./localizer/translations.json")
	if err != nil {
		log.Fatalf("Failed to load translations: %v", err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var update, err = telegram.Parse(r)
	if err != nil {
		log.Printf("Error parsing update, %s", err.Error())
		return
	}

	message := update.Message
	text := message.Text
	command := Command(text)
	locale := localizer.EN

	fmt.Printf("Received message text: %v\n", text)

	switch command {
	case Start:
		handleStart(message, locale)
	case Note:
		handleNote(message, locale)
	case Info:
		handleInfo(message, locale)
	default:
		handleUnknownState(message, locale)
	}
	ctx := context.Background()
	saveNote(ctx, update.Message)

}

func handleStart(message telegram.Message, locale localizer.Locale) {
	sendMessage(message.Chat.Id, locale, "welcome")
}

func handleNote(message telegram.Message, locale localizer.Locale) {
	sendMessage(message.Chat.Id, locale, "start_note")
}

func handleUnknownState(message telegram.Message, locale localizer.Locale) {
	// logic to handle unknown state
}

func handleInfo(message telegram.Message, locale localizer.Locale) {
	sendMessage(message.Chat.Id, locale, "info")
}

func sendMessage(chatId int, locale localizer.Locale, text string) {
	body, err := telegram.SendMessage(chatId, localizer.Localize(locale, text))
	if err != nil {
		log.Printf("Got error %s from telegram, reponse body is %s", err.Error(), body)
	}
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
