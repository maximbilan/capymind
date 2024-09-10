package bot

import (
	"log"

	"github.com/capymind/internal/firestore"
	"github.com/capymind/internal/telegram"
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
	}

	return &user
}

// Update the user's data in the database if necessary
func updateUser(user *firestore.User) {
	if user == nil {
		return
	}

	// Setup the database connection
	client, ctx := createClient()
	defer client.Close()

	// Check if the user exists
	fetchedUser, err := firestore.GetUser(ctx, client, user.ID)
	if err != nil {
		log.Printf("[User] Error fetching user from firestore, %s", err.Error())
	} else if fetchedUser == nil {
		fetchedUser.ID = user.ID
		// Save the first user data to firestore
		err := firestore.SaveUser(ctx, client, *user)
		if err != nil {
			log.Printf("[User] Error saving the frist user data to firestore, %s", err.Error())
		}
	}

	// Update the user's data
	fetchedUser.ChatID = user.ChatID
	fetchedUser.UserName = user.UserName
	fetchedUser.FirstName = user.FirstName
	fetchedUser.LastName = user.LastName

	user = fetchedUser
}

// Save a user to the database
func saveUser(user *firestore.User) {
	client, ctx := createClient()
	defer client.Close()

	err := firestore.SaveUser(ctx, client, *user)
	if err != nil {
		log.Printf("[User] Error saving user to firestore, %s", err.Error())
	}
}
