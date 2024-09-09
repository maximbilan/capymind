package bot

import (
	"github.com/capymind/internal/firestore"
	"github.com/capymind/internal/telegram"
)

type Job struct {
	Command    Command
	Parameters []string
}

// creates a job from an update for the user
func createJob(update telegram.Update) (*Job, *firestore.User) {
	var chatID int64
	var telegramUser *telegram.User
	var input *string

	// Check if the update is a callback query or a message
	callbackQuery := update.CallbackQuery
	if callbackQuery != nil && callbackQuery.Data != "" {
		chatID = callbackQuery.Message.Chat.ID
		telegramUser = callbackQuery.From
		input = &callbackQuery.Data
	} else {
		message := update.Message
		chatID = message.Chat.ID
		telegramUser = message.From
		input = &message.Text
	}

	// Check if the user is valid and if the input is valid
	if telegramUser == nil || telegramUser.ID == 0 || input == nil {
		return nil, nil
	}

	// Create a user from the telegram user
	user := createUser(chatID, *telegramUser)

	// Finish this (multiple return values)
	command, parameters := ParseCommand(*input)

	// Create a job with the command and parameters
	job := Job{
		Command:    command,
		Parameters: parameters,
	}

	return &job, &user
}

func createUser(chatID int64, tUser telegram.User) *firestore.User {
	user := firestore.User{
		ID:        tUser.StringID(),
		ChatID:    chatID,
		UserName:  &tUser.UserName,
		FirstName: &tUser.FirstName,
		LastName:  &tUser.LastName,
	}
	return &user
}
