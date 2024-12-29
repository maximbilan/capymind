package app

import (
	"testing"

	"github.com/capymind/internal/database"
)

func TestHelpHandler(t *testing.T) {
	session := createSession(&Job{Command: "/help"}, &database.User{}, nil)
	handleHelp(session)
	if session.Job.Output[0].TextID != "commands_hint" {
		t.Errorf("Expected output text to be 'commands_hint', got %s", session.Job.Output[0].TextID)
	}
}
