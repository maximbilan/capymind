package bot

import (
	"context"
	"log"

	"github.com/capymind/internal/firestore"
	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
)

// Create a user from an update
func createUser(update telegram.Update) *firestore.User {
	var chatID int64
	var telegramUser *telegram.User

	// Check if the update is a callback query or a message
	callbackQuery := update.CallbackQuery
	if callbackQuery != nil && callbackQuery.Data != "" {
		chatID = callbackQuery.Message.Chat.ID
		telegramUser = callbackQuery.From
	} else {
		message := update.Message
		chatID = message.Chat.ID
		telegramUser = message.From
	}

	// Check if the user is valid
	if telegramUser == nil || telegramUser.ID == 0 {
		return nil
	}

	// Create a user from the telegram user
	user := firestore.User{
		ID:        telegramUser.StringID(),
		ChatID:    chatID,
		UserName:  &telegramUser.UserName,
		FirstName: &telegramUser.FirstName,
		LastName:  &telegramUser.LastName,
		Locale:    &telegramUser.LanguageCode,
	}

	return &user
}

// Update the user's data in the database if necessary
func updateUser(user *firestore.User, ctx *context.Context) *firestore.User {
	if user == nil {
		return nil
	}

	// Check if the user exists
	fetchedUser, err := firestore.GetUser(ctx, user.ID)
	if err != nil {
		log.Printf("[User] Error fetching user from firestore, %s", err.Error())

		// If the user doesn't exist, create a new user
		fetchedUser = &firestore.User{
			ID: user.ID,
		}
	}

	// Update the user's data
	fetchedUser.ChatID = user.ChatID
	fetchedUser.UserName = user.UserName
	fetchedUser.FirstName = user.FirstName
	fetchedUser.LastName = user.LastName

	// Update the user's locale from Telegram if it's valid
	if fetchedUser.Locale == nil && user.Locale != nil {
		userLocale := *user.Locale
		if userLocale == translator.EN.String() || userLocale == translator.UK.String() {
			fetchedUser.Locale = user.Locale
		}
	}

	return fetchedUser
}

// Save a user to the database
func saveUser(user *firestore.User, ctx *context.Context) {
	err := firestore.SaveUser(ctx, *user)
	if err != nil {
		log.Printf("[User] Error saving user to firestore, %s", err.Error())
	}
}
