package bot

import (
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
	Command    Command
	Parameters []string
	Input      *string
	Output     *JobResult
}

// Creates a job from an update
func createJob(update telegram.Update) *Job {
	var input *string

	// Check if the update is a callback query or a message
	callbackQuery := update.CallbackQuery
	if callbackQuery != nil && callbackQuery.Data != "" {
		input = &callbackQuery.Data
	} else {
		message := update.Message
		input = &message.Text
	}

	// Check if the input is valid
	if input == nil {
		return nil
	}

	// Create a command and parameters from the input
	command, parameters := ParseCommand(*input)

	// Create a job with the command and parameters
	job := Job{
		Command:    command,
		Parameters: parameters,
		Input:      input,
	}

	return &job
}
