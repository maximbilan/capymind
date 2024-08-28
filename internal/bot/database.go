package bot

import (
	"context"
	"fmt"
	"log"
	"time"

	google "cloud.google.com/go/firestore"
	"github.com/capymind/internal/firestore"
	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
)

func createClient() (*google.Client, context.Context) {
	ctx := context.Background()
	var client, err = firestore.NewClient(ctx)
	if err != nil {
		log.Printf("[Database] Error creating firestore client, %s", err.Error())
	}
	return client, ctx
}

func CreateOrUpdateUser(message telegram.Message) {
	client, ctx := createClient()
	defer client.Close()

	var user = firestore.User{
		ID:   fmt.Sprintf("%d", message.Chat.Id),
		Name: message.From.Username,
	}

	err := firestore.NewUser(ctx, client, user)
	if err != nil {
		log.Printf("[Database] Error creating user in firestore, %s", err.Error())
	}
}

func SetupLocale(userId string, locale string) {
	client, ctx := createClient()
	defer client.Close()

	err := firestore.UpdateUserLocale(ctx, client, userId, locale)
	if err != nil {
		log.Printf("[Database] Error updating user locale in firestore, %s", err.Error())
	}
}

func GetUserLocale(message telegram.Message) *translator.Locale {
	client, ctx := createClient()
	defer client.Close()

	var userId = fmt.Sprintf("%d", message.From.ID)
	localeStr, err := firestore.UserLocale(ctx, client, userId)
	if err != nil {
		log.Printf("[Database] Error getting user locale from firestore, %s", err.Error())
		return nil
	}
	if localeStr == nil || *localeStr == "" {
		log.Printf("[Database] User locale is nil")
		return nil
	}

	var locale = translator.Locale(*localeStr)
	return &locale
}

func SaveNote(message telegram.Message) {
	client, ctx := createClient()
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

	err := firestore.NewRecord(ctx, client, user, note)
	if err != nil {
		log.Printf("[Database] Error saving note to firestore, %s", err.Error())
	}
}

func GetLastNote(message telegram.Message) *firestore.Note {
	client, ctx := createClient()
	defer client.Close()

	var userId = fmt.Sprintf("%d", message.From.ID)
	note, err := firestore.LastNote(ctx, client, userId)
	if err != nil {
		log.Printf("[Database] Error getting last note from firestore, %s", err.Error())
	}
	return note
}

func SetupTimezone(userId string, secondsFromUTC int) {
	client, ctx := createClient()
	defer client.Close()

	err := firestore.UpdateUserTimezone(ctx, client, userId, secondsFromUTC)
	if err != nil {
		log.Printf("[Database] Error updating user timezone in firestore, %s", err.Error())
	}
}
