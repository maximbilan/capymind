package app

import (
	"testing"

	"github.com/capymind/internal/database"
)

func TestFirstStartHandler(t *testing.T) {
	session := createSession(&Job{Command: "/start"}, &database.User{}, nil)
	handleStart(session)

	if session.Job.Output[0].TextID != "welcome_onboarding" {
		t.Error("Expected 'welcome_onboarding', got", session.Job.Output[0].TextID)
	}
	if session.Job.Output[1].TextID != "language_select" {
		t.Error("Expected 'language_select', got", session.Job.Output[1].TextID)
	}
}

func TestSecondStartHandler(t *testing.T) {
	session := createSession(&Job{Command: "/start"}, &database.User{
		IsOnboarded: true,
	}, nil)
	handleStart(session)

	if session.Job.Output[0].TextID != "timezone_select" {
		t.Error("Expected 'timezone_select', got", session.Job.Output[0].TextID)
	}
}

func TestThirdStartHandler(t *testing.T) {
	secondsFromUTC := 7200
	session := createSession(&Job{Command: "/start"}, &database.User{
		IsOnboarded:    true,
		SecondsFromUTC: &secondsFromUTC,
	}, nil)
	handleStart(session)

	if session.Job.Output[0].TextID != "welcome" {
		t.Error("Expected 'welcome', got", session.Job.Output[0].TextID)
	}
	if session.Job.Output[0].Buttons[0].TextID != "make_record_to_journal" {
		t.Error("Expected 'make_record_to_journal', got", session.Job.Output[0].Buttons[0].TextID)
	}
	if session.Job.Output[0].Buttons[1].TextID != "how_to_use" {
		t.Error("Expected 'how_to_use', got", session.Job.Output[0].Buttons[1].TextID)
	}
}
