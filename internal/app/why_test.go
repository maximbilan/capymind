package app

import (
	"testing"

	"github.com/capymind/internal/database"
)

func TestWhyHandler(t *testing.T) {
	session := createSession(&Job{Command: "/why"}, &database.User{}, nil, nil)
	handleWhy(session)

	if session.Job.Output[0].TextID != "why_descr" {
		t.Errorf("Expected output text to be 'why_descr', got %s", session.Job.Output[0].TextID)
	}
}
