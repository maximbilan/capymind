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
	update := telegram.Parse(r)
	if update == nil {
		return
	}

	callbackQuery := update.CallbackQuery
	if callbackQuery != nil && callbackQuery.Data != "" {
		log.Printf("[Bot] Received callback data: %s", callbackQuery.Data)
		locale, ok := translator.ParseLocale(callbackQuery.Data)
		if ok && locale != nil {
			userId := fmt.Sprintf("%d", callbackQuery.From.ID)
			setupLocale(userId, *locale)
			LocalizeAndSendMessage(callbackQuery.Message.Chat.Id, translator.Locale(*locale), "locale_set")
		}
		return
	}

	message := update.Message

	var locale translator.Locale
	userLocale := getUserLocale(message)
	if userLocale != nil {
		locale = *userLocale
	} else {
		locale = translator.EN
	}

	text := message.Text
	command := Command(text)

	log.Printf("[Bot] Received message text: %s", text)

	switch command {
	case Start:
		handleStart(message, locale)
	case Note:
		handleNote(message, locale)
	case Last:
		handleLast(message, locale)
	case Locale:
		handleLocale(message, locale)
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
	LocalizeAndSendMessage(message.Chat.Id, locale, "welcome")
	handleUser(message, locale)
}

func handleNote(message telegram.Message, locale translator.Locale) {
	LocalizeAndSendMessage(message.Chat.Id, locale, "start_note")

	userId := message.From.ID
	userIds.Append(userId)
}

func handleLast(message telegram.Message, locale translator.Locale) {
	ctx := context.Background()
	note := getLastNote(ctx, message)
	if note != nil {
		var response string = translator.Translate(locale, "your_last_note") + note.Text
		LocalizeAndSendMessage(message.Chat.Id, locale, response)
	} else {
		LocalizeAndSendMessage(message.Chat.Id, locale, "no_notes")
	}
}

func handleLocale(message telegram.Message, locale translator.Locale) {
	replyMarkup := telegram.InlineKeyboardMarkup{
		InlineKeyboard: [][]telegram.InlineKeyboardButton{
			{
				{Text: translator.English.String(), CallbackData: translator.GetLocaleParameter(translator.EN)},
				{Text: translator.Ukrainian.String(), CallbackData: translator.GetLocaleParameter(translator.UK)},
			},
		},
	}
	LocalizeAndSendMessageWithReply(message.Chat.Id, locale, "language_select", &replyMarkup)
}

func setupLocale(userId string, locale string) {
	ctx := context.Background()
	var client, err = firestore.NewClient(ctx)
	if err != nil {
		log.Printf("Error creating firestore client, %s", err.Error())
	}
	defer client.Close()

	err = firestore.UpdateUserLocale(ctx, client, userId, locale)
	if err != nil {
		log.Printf("Error updating user locale in firestore, %s", err.Error())
	}
}

func handleUnknownState(message telegram.Message, locale translator.Locale) {
	userId := message.From.ID
	if userIds.Contains(userId) {
		ctx := context.Background()
		saveNote(ctx, message)
		LocalizeAndSendMessage(message.Chat.Id, locale, "finish_note")
		userIds.Remove(userId)
	} else {
		handleHelp(message, locale)
	}
}

func handleInfo(message telegram.Message, locale translator.Locale) {
	LocalizeAndSendMessage(message.Chat.Id, locale, "info")
}

func handleHelp(message telegram.Message, locale translator.Locale) {
	LocalizeAndSendMessage(message.Chat.Id, locale, "commands_hint")
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

func getUserLocale(message telegram.Message) *translator.Locale {
	ctx := context.Background()
	var client, err = firestore.NewClient(ctx)
	if err != nil {
		log.Printf("Error creating firestore client, %s", err.Error())
	}
	defer client.Close()

	var userId = fmt.Sprintf("%d", message.From.ID)
	localeStr, err := firestore.UserLocale(ctx, client, userId)
	if err != nil {
		log.Printf("Error getting user locale from firestore, %s", err.Error())
		return nil
	}
	if localeStr == nil || *localeStr == "" {
		log.Printf("User locale is nil")
		return nil
	}

	var locale = translator.Locale(*localeStr)
	return &locale
}
