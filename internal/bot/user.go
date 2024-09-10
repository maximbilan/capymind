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

	// Check if the user exists
	fetchedUser := fetchUserFromDB(user.ID)
	if fetchedUser == nil {
		return
	}

	// Update the user's data if necessary
	var hasChanges bool = false
	if fetchedUser.ChatID != user.ChatID {
		fetchedUser.ChatID = user.ChatID
		hasChanges = true
	}
	if fetchedUser.UserName != user.UserName {
		fetchedUser.UserName = user.UserName
		hasChanges = true
	}
	if fetchedUser.FirstName != user.FirstName {
		fetchedUser.FirstName = user.FirstName
		hasChanges = true
	}
	if fetchedUser.LastName != user.LastName {
		fetchedUser.LastName = user.LastName
		hasChanges = true
	}

	// Save the user to the database if there are changes
	if hasChanges {
		saveUserToDB(fetchedUser)
	}

	user = fetchedUser
}

// Fetch a user from the database
func fetchUserFromDB(userID string) *firestore.User {
	client, ctx := createClient()
	defer client.Close()

	user, err := firestore.GetUser(ctx, client, userID)
	if err != nil {
		log.Printf("[User] Error fetching user from firestore, %s", err.Error())
	}
	return user
}

// Save a user to the database
func saveUserToDB(user *firestore.User) {
	client, ctx := createClient()
	defer client.Close()

	err := firestore.SaveUser(ctx, client, *user)
	if err != nil {
		log.Printf("[User] Error saving user to firestore, %s", err.Error())
	}
}
