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

func createOrUpdateUser(chatId int, userId string, name *string) {
	client, ctx := createClient()
	defer client.Close()

	var user = firestore.User{
		ID:   fmt.Sprintf("%d", chatId),
		Name: name,
	}

	err := firestore.NewUser(ctx, client, user)
	if err != nil {
		log.Printf("[Database] Error creating user in firestore, %s", err.Error())
	}
}

func userExists(userId string) bool {
	client, ctx := createClient()
	defer client.Close()

	exists, err := firestore.UserExists(ctx, client, userId)
	if err != nil {
		log.Printf("[Database] Error checking if user exists in firestore, %s", err.Error())
	}
	return exists
}

func setupLocale(userId string, locale string) {
	client, ctx := createClient()
	defer client.Close()

	err := firestore.UpdateUserLocale(ctx, client, userId, locale)
	if err != nil {
		log.Printf("[Database] Error updating user locale in firestore, %s", err.Error())
	}
}

func getUserLocale(message telegram.Message) *translator.Locale {
	var userId = fmt.Sprintf("%d", message.From.ID)
	return getUserLocaleByUserId(userId)
}

func getUserLocaleByUserId(userId string) *translator.Locale {
	client, ctx := createClient()
	defer client.Close()

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

func saveNote(message telegram.Message) {
	client, ctx := createClient()
	defer client.Close()

	var user = firestore.User{
		ID:   fmt.Sprintf("%d", message.From.ID),
		Name: &message.From.UserName,
	}

	timestamp := time.Now()
	var note = firestore.Note{
		ID:        fmt.Sprintf("%d", message.Chat.ID),
		Text:      message.Text,
		Timestamp: timestamp,
	}

	err := firestore.NewRecord(ctx, client, user, note)
	if err != nil {
		log.Printf("[Database] Error saving note to firestore, %s", err.Error())
	}
}

func getLastNote(message telegram.Message) *firestore.Note {
	client, ctx := createClient()
	defer client.Close()

	var userId = fmt.Sprintf("%d", message.From.ID)
	note, err := firestore.LastNote(ctx, client, userId)
	if err != nil {
		log.Printf("[Database] Error getting last note from firestore, %s", err.Error())
	}
	return note
}

func getNotes(message telegram.Message) []firestore.Note {
	client, ctx := createClient()
	defer client.Close()

	var userId = fmt.Sprintf("%d", message.From.ID)
	notes, err := firestore.GetNotes(ctx, client, userId)
	if err != nil {
		log.Printf("[Database] Error getting notes from firestore, %s", err.Error())
	}
	return notes
}

func setupTimezone(userId string, secondsFromUTC int) {
	client, ctx := createClient()
	defer client.Close()

	err := firestore.UpdateUserTimezone(ctx, client, userId, secondsFromUTC)
	if err != nil {
		log.Printf("[Database] Error updating user timezone in firestore, %s", err.Error())
	}
}

func getTimeZone(userId string) *int64 {
	client, ctx := createClient()
	defer client.Close()

	secondsFromUTC, err := firestore.UserTimezone(ctx, client, userId)
	if err != nil {
		log.Printf("[Database] Error getting user timezone from firestore, %s", err.Error())
	}
	return secondsFromUTC
}

func saveLastChatID(chatID int64, userID string) {
	client, ctx := createClient()
	defer client.Close()

	err := firestore.SaveLastChatID(ctx, client, userID, chatID)
	if err != nil {
		log.Printf("[Database] Error saving last chat id to firestore, %s", err.Error())
	}
}

func startWritingMode(userId string) {
	client, ctx := createClient()
	defer client.Close()

	err := firestore.StartWriting(ctx, client, userId)
	if err != nil {
		log.Printf("[Database] Error starting writing mode in firestore, %s", err.Error())
	}
}

func stopWritingMode(userId string) {
	client, ctx := createClient()
	defer client.Close()

	err := firestore.StopWriting(ctx, client, userId)
	if err != nil {
		log.Printf("[Database] Error stopping writing mode in firestore, %s", err.Error())
	}
}

func isWriting(userId string) bool {
	client, ctx := createClient()
	defer client.Close()

	isWriting, err := firestore.UserWritingStatus(ctx, client, userId)
	if err != nil {
		log.Printf("[Database] Error getting writing mode from firestore, %s", err.Error())
	}
	return isWriting
}
