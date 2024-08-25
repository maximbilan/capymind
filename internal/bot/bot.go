package bot

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/capymind/internal/firestore"
	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
	"github.com/capymind/internal/utils"
)

var userIds *utils.ThreadSafeArray[int64]

func init() {
	userIds = utils.NewThreadSafeArray[int64]()
}

func Parse(w http.ResponseWriter, r *http.Request) {
	var update, err = telegram.Parse(r)
	if err != nil {
		log.Printf("Error parsing update, %s", err.Error())
		return
	}

	message := update.Message
	text := message.Text
	command := Command(text)
	locale := translator.EN

	fmt.Printf("Received message text: %v\n", text)

	switch command {
	case Start:
		handleStart(message, locale)
	case Note:
		handleNote(message, locale)
	case Last:
		handleLast(message, locale)
	case Locale:
		handleLocale(message)
	case Info:
		handleInfo(message, locale)
	case Help:
		handleHelp(message, locale)
	default:
		handleUnknownState(message, locale)
	}
}

func handleUser(message telegram.Message, locale translator.Locale) {
	if message.Text == "" {
		return
	}
	ctx := context.Background()
	createOrUpdateUser(ctx, message)
}

func handleStart(message telegram.Message, locale translator.Locale) {
	sendMessage(message.Chat.Id, locale, "welcome")
	handleUser(message, locale)
}

func handleNote(message telegram.Message, locale translator.Locale) {
	sendMessage(message.Chat.Id, locale, "start_note")

	userId := message.From.ID
	userIds.Append(userId)
}

func handleLast(message telegram.Message, locale translator.Locale) {
	ctx := context.Background()
	note := getLastNote(ctx, message)
	if note != nil {
		var response string = translator.Translate(locale, "your_last_note") + note.Text
		sendLocalizedMessage(message.Chat.Id, locale, response)
	} else {
		sendMessage(message.Chat.Id, locale, "no_notes")
	}
}

func handleLocale(message telegram.Message) {
	replyMarkup := telegram.InlineKeyboardMarkup{
		InlineKeyboard: [][]telegram.InlineKeyboardButton{
			{
				{Text: "English", CallbackData: "en"},
				{Text: "Ukrainian", CallbackData: "uk"},
			},
		},
	}

	err := telegram.SendMessageWithButtons(message.Chat.Id, "Choose an option:", replyMarkup)
	if err != nil {
		fmt.Println("Error sending message:", err)
	} else {
		fmt.Println("Message sent successfully")
	}
}

// func setupLocale(message telegram.Message) {
// 	ctx := context.Background()
// 	var client, err = firestore.NewClient(ctx)
// 	if err != nil {
// 		log.Printf("Error creating firestore client, %s", err.Error())
// 	}
// 	defer client.Close()

// 	var userId = fmt.Sprintf("%d", message.From.ID)
// }

func handleUnknownState(message telegram.Message, locale translator.Locale) {
	userId := message.From.ID
	if userIds.Contains(userId) {
		ctx := context.Background()
		saveNote(ctx, message)
		sendMessage(message.Chat.Id, locale, "finish_note")
		userIds.Remove(userId)
	} else {
		handleHelp(message, locale)
	}
}

func handleInfo(message telegram.Message, locale translator.Locale) {
	sendMessage(message.Chat.Id, locale, "info")
}

func handleHelp(message telegram.Message, locale translator.Locale) {
	sendMessage(message.Chat.Id, locale, "commands_hint")
}

func sendMessage(chatId int, locale translator.Locale, text string) {
	localizedMessage := translator.Translate(locale, text)
	sendLocalizedMessage(chatId, locale, localizedMessage)
}

func sendLocalizedMessage(chatId int, locale translator.Locale, text string) {
	body, err := telegram.SendMessage(chatId, text)
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

func getLastNote(ctx context.Context, message telegram.Message) *firestore.Note {
	var client, err = firestore.NewClient(ctx)
	if err != nil {
		log.Printf("Error creating firestore client, %s", err.Error())
	}
	defer client.Close()

	var userId = fmt.Sprintf("%d", message.From.ID)
	note, err := firestore.LastNote(ctx, client, userId)
	if err != nil {
		log.Printf("Error getting last note from firestore, %s", err.Error())
	}
	return note
}

func createOrUpdateUser(ctx context.Context, message telegram.Message) {
	var client, err = firestore.NewClient(ctx)
	if err != nil {
		log.Printf("Error creating firestore client, %s", err.Error())
	}
	defer client.Close()

	var user = firestore.User{
		ID:   fmt.Sprintf("%d", message.Chat.Id),
		Name: message.From.Username,
	}

	err = firestore.NewUser(ctx, client, user)
	if err != nil {
		log.Printf("Error creating user in firestore, %s", err.Error())
	}
}
