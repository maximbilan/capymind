package app

import (
	"context"
	"log"

	"github.com/capymind/internal/botservice"
	"github.com/capymind/internal/database"
	"github.com/capymind/internal/translator"
)

func createUser(message botservice.BotMessage) *database.User {
	user := database.User{
		ID:        message.UserID,
		ChatID:    message.ChatID,
		UserName:  &message.UserName,
		FirstName: &message.FirstName,
		LastName:  &message.LastName,
		Locale:    &message.LanguageCode,
	}
	return &user
}

// Update the user's data in the database if necessary
func updateUser(user *database.User, ctx *context.Context, userStorage database.UserStorage) *database.User {
	if user == nil {
		return nil
	}

	// Check if the user exists
	fetchedUser, err := userStorage.GetUser(ctx, user.ID)
	if err != nil {
		log.Printf("[User] Error fetching user from firestore, %s", err.Error())

		// If the user doesn't exist, create a new user
		fetchedUser = &database.User{
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

//coverage:ignore
func saveUser(user *database.User, ctx *context.Context, userStorage database.UserStorage) {
	err := userStorage.SaveUser(ctx, *user)
	if err != nil {
		log.Printf("[User] Error saving user to firestore, %s", err.Error())
	}
}

// Check if the user is an admin
func isAdmin(user *database.User) bool {
	if user == nil {
		return false
	}
	role := user.Role
	return database.IsAdmin(role)
}
