package app

import (
	"github.com/capymind/internal/database"
	"github.com/capymind/internal/telegram"
)

type JobResultTextButton struct {
	TextID   string
	Callback string
}

type JobResult struct {
	TextID  string
	Buttons []JobResultTextButton
}

type Job struct {
	Command     Command
	LastCommand Command
	Parameters  []string
	Input       *string // Raw input
	Output      []JobResult
}

// Creates a job from an update
func createJob(update telegram.Update, user *database.User) *Job {
	var input *string

	// Check if the update is a callback query or a message
	callbackQuery := update.CallbackQuery
	if callbackQuery != nil && callbackQuery.Data != "" {
		input = &callbackQuery.Data
	} else if update.Message != nil {
		message := update.Message
		input = &message.Text
	}

	// Check if the input is valid
	if input == nil || *input == "" {
		return nil
	}

	// Create a command and parameters from the input
	command, parameters := ParseCommand(*input)

	// Gets the previous command
	var lastCommand Command
	if user != nil && user.LastCommand != nil {
		lastCommand, _ = ParseCommand(*user.LastCommand)
	} else {
		lastCommand = None
	}

	// Create a job with the command and parameters
	job := Job{
		Command:     command,
		LastCommand: lastCommand,
		Parameters:  parameters,
		Input:       input,
	}

	return &job
}
