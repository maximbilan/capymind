package app

import (
	"github.com/capymind/internal/botservice"
	"github.com/capymind/internal/database"
)

type Job struct {
	Command     Command
	LastCommand Command
	Parameters  []string
	Input       *string // Raw input
	Output      []botservice.BotResult
}

// Creates a job from an update
func createJob(input string, user *database.User) *Job {
	// Create a command and parameters from the input
	command, parameters := ParseCommand(input)

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
		Input:       &input,
	}

	return &job
}
