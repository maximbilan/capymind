package app

import (
	"context"
	"testing"

	"github.com/capymind/internal/database"
)

func TestSessionLocale(t *testing.T) {
	user := &database.User{}
	user.Locale = nil

	ctx := context.Background()

	session := createSession(&Job{}, user, nil, &ctx)
	if session.Locale() != "en" {
		t.Fatalf("Expected 'en', got %s", session.Locale())
	}

	locale := "uk"
	user.Locale = &locale

	session = createSession(&Job{}, user, nil, &ctx)
	if session.Locale() != "uk" {
		t.Fatalf("Expected 'uk', got %s", session.Locale())
	}
}

func TestIsAdmin(t *testing.T) {
	user := &database.User{}
	admin := database.Admin
	user.Role = nil

	ctx := context.Background()

	session := createSession(&Job{}, user, nil, &ctx)
	if session.isAdmin() {
		t.Fatalf("Expected false, got true")
	}

	user.Role = &admin

	session = createSession(&Job{}, user, nil, &ctx)
	if !session.isAdmin() {
		t.Fatalf("Expected true, got false")
	}
}

func TestCommandExecution(t *testing.T) {
	job := &Job{
		Command: "/help",
	}
	locale := "en"
	user := &database.User{}
	user.Locale = &locale

	ctx := context.Background()

	session := createSession(job, user, nil, &ctx)
	handleSession(session)

	if session.Job.Output[0].TextID != "commands_hint" {
		t.Fatalf("Expected 'commands_hint', got %s", session.Job.Output[0].TextID)
	}
}
