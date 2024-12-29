package app

import (
	"context"
	"testing"

	"github.com/capymind/internal/database"
)

func TestHelpCommand(t *testing.T) {
	job := &Job{
		Command: "/help",
	}
	locale := "en"
	user := &database.User{}
	user.Locale = &locale

	ctx := context.Background()

	session := createSession(job, user, &ctx)
	handleSession(session)

	if session.Job.Output[0].TextID != "commands_hint" {
		t.Fatalf("Expected 'commands_hint', got %s", session.Job.Output[0].TextID)
	}
}

func TestWhyCommand(t *testing.T) {
	job := &Job{
		Command: "/why",
	}
	locale := "en"
	user := &database.User{}
	user.Locale = &locale

	ctx := context.Background()

	session := createSession(job, user, &ctx)
	handleSession(session)

	if session.Job.Output[0].TextID != "why_descr" {
		t.Fatalf("Expected 'why_descr', got %s", session.Job.Output[0].TextID)
	}
}
