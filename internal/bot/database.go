package bot

import (
	"context"
	"log"

	google "cloud.google.com/go/firestore"
	"github.com/capymind/internal/firestore"
	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
)

func convertTelegramUser(message *telegram.Message) *firestore.User {
	if message == nil {
		return nil
	}

	user := firestore.User{
		ID:        message.UserID(),
		ChatID:    message.ChatID(),
		UserName:  &message.From.UserName,
		FirstName: &message.From.FirstName,
		LastName:  &message.From.LastName,
	}
	return &user
}

func createClient() (*google.Client, context.Context) {
	ctx := context.Background()
	var client, err = firestore.NewClient(ctx)
	if err != nil {
		log.Printf("[Database] Error creating firestore client, %s", err.Error())
	}
	return client, ctx
}

func createOrUpdateUser(user firestore.User) {
	client, ctx := createClient()
	defer client.Close()

	err := firestore.NewUser(ctx, client, user)
	if err != nil {
		log.Printf("[Database] Error creating user in firestore, %s", err.Error())
	}
}

func userExists(userID string) bool {
	client, ctx := createClient()
	defer client.Close()

	exists, err := firestore.UserExists(ctx, client, userID)
	if err != nil {
		log.Printf("[Database] Error checking if user exists in firestore, %s", err.Error())
	}
	return exists
}

func setupLocale(userID string, locale string) {
	client, ctx := createClient()
	defer client.Close()

	err := firestore.UpdateUserLocale(ctx, client, userID, locale)
	if err != nil {
		log.Printf("[Database] Error updating user locale in firestore, %s", err.Error())
	}
}

func getUserLocale(message telegram.Message) *translator.Locale {
	userID := message.UserID()
	return getUserLocaleByUserID(userID)
}

func getUserLocaleByUserID(userID string) *translator.Locale {
	client, ctx := createClient()
	defer client.Close()

	localeStr, err := firestore.UserLocale(ctx, client, userID)
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

// func saveNote(message telegram.Message) {
// 	client, ctx := createClient()
// 	defer client.Close()

// 	var user = *convertTelegramUser(&message)

// 	timestamp := time.Now()
// 	var note = firestore.Note{
// 		ID:        user.ID,
// 		Text:      message.Text,
// 		Timestamp: timestamp,
// 	}

// 	err := firestore.NewRecord(ctx, client, user, note)
// 	if err != nil {
// 		log.Printf("[Database] Error saving note to firestore, %s", err.Error())
// 	}
// }

// func getLastNote(message telegram.Message) *firestore.Note {
// 	client, ctx := createClient()
// 	defer client.Close()

// 	userID := message.UserID()
// 	note, err := firestore.LastNote(ctx, client, userID)
// 	if err != nil {
// 		log.Printf("[Database] Error getting last note from firestore, %s", err.Error())
// 	}
// 	return note
// }

// func setupTimezone(userID string, secondsFromUTC int) {
// 	client, ctx := createClient()
// 	defer client.Close()

// 	err := firestore.UpdateUserTimezone(ctx, client, userID, secondsFromUTC)
// 	if err != nil {
// 		log.Printf("[Database] Error updating user timezone in firestore, %s", err.Error())
// 	}
// }

func getTimeZone(userID string) *int {
	client, ctx := createClient()
	defer client.Close()

	secondsFromUTC, err := firestore.UserTimezone(ctx, client, userID)
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

func startWritingMode(userID string) {
	client, ctx := createClient()
	defer client.Close()

	err := firestore.StartWriting(ctx, client, userID)
	if err != nil {
		log.Printf("[Database] Error starting writing mode in firestore, %s", err.Error())
	}
}

func stopWritingMode(userID string) {
	client, ctx := createClient()
	defer client.Close()

	err := firestore.StopWriting(ctx, client, userID)
	if err != nil {
		log.Printf("[Database] Error stopping writing mode in firestore, %s", err.Error())
	}
}

func isWriting(userID string) bool {
	client, ctx := createClient()
	defer client.Close()

	isWriting, err := firestore.UserWritingStatus(ctx, client, userID)
	if err != nil {
		log.Printf("[Database] Error getting writing mode from firestore, %s", err.Error())
	}
	return isWriting
}
