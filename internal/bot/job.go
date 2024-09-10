package bot

import (
	"github.com/capymind/internal/telegram"
)

type Job struct {
	Command    Command
	Parameters []string
	Input      *string
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

func (job Job) execute() {
	command := job.Command
	parameters := job.Parameters

	switch command {
	case Start:
		// handleStart(message, locale)
	case Note:
		// handleNote(message, locale)
	case Last:
		// handleLast(message, locale)
	case Analysis:
		// handleAnalysis(message, locale)
	case Language:
		// handleLanguage(message, locale)
	case Timezone:
		// handleTimezone(message, locale)
	case Help:
		// handleHelp(message, locale)
	case None:
		// Typing mode
	default:
		// Unknown command
		// handleUnknownState(message, locale)
	}
}
