package capymind

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/capymind/firestore"
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

	ctx := context.Background()

	fmt.Println("Saving note to firestore...")
	var firestoreClient, errFirestore = firestore.NewClient(ctx)
	if errFirestore != nil {
		log.Printf("error creating firestore client, %s", errFirestore.Error())
	}
	defer firestoreClient.Close()

	timestamp := time.Now()
	var note = firestore.Note{
		ID:        fmt.Sprintf("%d", update.Message.Chat.Id),
		Text:      update.Message.Text,
		Timestamp: timestamp,
	}

	userId := fmt.Sprintf("%d", update.Message.From.ID)
	firestore.SaveNote(ctx, firestoreClient, userId, note)

	var telegramResponseBody, errTelegram = telegram.SendMessage(update.Message.Chat.Id, "Now I can respond to you!")
	if errTelegram != nil {
		log.Printf("got error %s from telegram, reponse body is %s", errTelegram.Error(), telegramResponseBody)
	}
}
