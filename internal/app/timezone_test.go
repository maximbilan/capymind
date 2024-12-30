package app

import (
	"testing"

	"github.com/capymind/internal/database"
)

func TestTimezoneHandler(t *testing.T) {
	session := createSession(&Job{Command: "/timezone"}, &database.User{}, nil)
	handleTimezone(session)

	if session.Job.Output[0].TextID != "timezone_select" {
		t.Error("Expected 'timezone_select', got", session.Job.Output[0].TextID)
	}
	if len(session.Job.Output[0].Buttons) != 25 {
		t.Error("Expected '25', got", len(session.Job.Output[0].Buttons))
	}
}

func TestTimezoneHandlerWithParam(t *testing.T) {
	session := createSession(&Job{Command: "/timezone 7200", Parameters: []string{"7200"}}, &database.User{}, nil)
	handleTimezone(session)

	if session.Job.Output[0].TextID != "timezone_set" {
		t.Error("Expected 'timezone_set', got", session.Job.Output[0].TextID)
	}
	if *session.User.SecondsFromUTC != 7200 {
		t.Error("Expected '7200', got", false)
	}
	if session.Job.Output[1].TextID != "welcome" {
		t.Error("Expected 'welcome', got", session.Job.Output[0].TextID)
	}
}

func TestTimezoneHandlerWithParamOnboarded(t *testing.T) {
	session := createSession(&Job{Command: "/timezone 0", Parameters: []string{"0"}}, &database.User{IsOnboarded: true}, nil)
	handleTimezone(session)

	if session.Job.Output[0].TextID != "timezone_set" {
		t.Error("Expected 'timezone_set', got", session.Job.Output[0].TextID)
	}
	if *session.User.SecondsFromUTC != 0 {
		t.Error("Expected '0', got", false)
	}
}
