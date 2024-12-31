package app

import (
	"testing"

	"github.com/capymind/internal/database"
)

func TestLanguageHandler(t *testing.T) {
	session := createSession(&Job{Command: "/language"}, &database.User{}, nil, nil)
	handleLanguage(session)

	if session.Job.Output[0].TextID != "language_select" {
		t.Errorf("Expected language_select, got %s", session.Job.Output[0].TextID)
	}
	if session.Job.Output[0].Buttons[0].TextID != "English 🇺🇸" {
		t.Errorf("Expected French, got %s", session.Job.Output[0].Buttons[0].TextID)
	}
	if session.Job.Output[0].Buttons[1].TextID != "Українська 🇺🇦" {
		t.Errorf("Expected Українська 🇺🇦, got %s", session.Job.Output[0].Buttons[1].TextID)
	}
}

func TestLanguageHandlerWithParameters(t *testing.T) {
	session := createSession(&Job{Command: "/language", Parameters: []string{"uk"}}, &database.User{}, nil, nil)
	handleLanguage(session)

	if session.Job.Parameters[0] != "uk" {
		t.Errorf("Expected uk, got %s", session.Job.Parameters[0])
	}
	if *session.User.Locale != "uk" {
		t.Errorf("Expected uk, got %s", *session.User.Locale)
	}

	if session.Job.Output[0].TextID != "locale_set" {
		t.Errorf("Expected locale_set, got %s", session.Job.Output[0].TextID)
	}
	if session.Job.Output[1].TextID != "timezone_select" {
		t.Errorf("Expected timezone_select, got %d", len(session.Job.Output))
	}
}

func TestLanguageHandlerWithParametersAndTimezone(t *testing.T) {
	time := 123456789
	session := createSession(&Job{Command: "/language", Parameters: []string{"en"}}, &database.User{SecondsFromUTC: &time}, nil, nil)
	handleLanguage(session)

	if session.Job.Parameters[0] != "en" {
		t.Errorf("Expected en, got %s", session.Job.Parameters[0])
	}
	if *session.User.Locale != "en" {
		t.Errorf("Expected en, got %s", *session.User.Locale)
	}
	if session.Job.Output[0].TextID != "locale_set" {
		t.Errorf("Expected locale_set, got %s", session.Job.Output[0].TextID)
	}
}
