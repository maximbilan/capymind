package app

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/capymind/internal/database"
)

func TestStartTherapySession(t *testing.T) {
	ctx := context.Background()
	user := &database.User{IsTyping: false}
	session := createSession(&Job{Command: TherapySession}, user, nil, &ctx)

	startTherapySession(session)

	if session.Job.Output[0].TextID != "start_therapy_session" {
		t.Fatalf("expected start_therapy_session, got %s", session.Job.Output[0].TextID)
	}
	if !session.User.IsTyping {
		t.Fatalf("expected IsTyping true")
	}
	if session.User.TherapySessionEndAt == nil {
		t.Fatalf("expected TherapySessionEndAt to be set")
	}
	if time.Until(*session.User.TherapySessionEndAt) < 29*time.Minute || time.Until(*session.User.TherapySessionEndAt) > 31*time.Minute {
		t.Fatalf("expected end time around 30m, got %s", session.User.TherapySessionEndAt.String())
	}
}

func TestEndTherapySession(t *testing.T) {
	ctx := context.Background()
	endAt := time.Now().Add(10 * time.Minute)
	user := &database.User{IsTyping: true, TherapySessionEndAt: &endAt}
	session := createSession(&Job{Command: TherapySession}, user, nil, &ctx)

	endTherapySession(session)

	if session.Job.Output[0].TextID != "therapy_session_ended" {
		t.Fatalf("expected therapy_session_ended, got %s", session.Job.Output[0].TextID)
	}
	if session.User.IsTyping {
		t.Fatalf("expected IsTyping false")
	}
	if session.User.TherapySessionEndAt != nil {
		t.Fatalf("expected TherapySessionEndAt to be nil")
	}
}

func TestRelayTherapyMessage(t *testing.T) {
	// Create a fake therapysession endpoint
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, I'm here for you."))
	}))
	defer ts.Close()
	os.Setenv("CAPY_THERAPYSESSION_URL", ts.URL)
	defer os.Unsetenv("CAPY_THERAPYSESSION_URL")

	ctx := context.Background()
	locale := "en"
	user := &database.User{Locale: &locale}
	session := createSession(&Job{Command: None}, user, nil, &ctx)

	relayTherapyMessage("hi", session)

	if len(session.Job.Output) == 0 {
		t.Fatalf("expected at least one output")
	}
	if session.Job.Output[0].TextID != "Hello, I'm here for you." {
		t.Fatalf("unexpected relay text: %s", session.Job.Output[0].TextID)
	}
}

func TestHandleSession_AutoEndWhenExpired(t *testing.T) {
	ctx := context.Background()
	past := time.Now().Add(-1 * time.Minute)
	user := &database.User{TherapySessionEndAt: &past, IsTyping: true}
	job := &Job{Command: None}
	session := createSession(job, user, nil, &ctx)

	handleSession(session)

	if len(session.Job.Output) == 0 || session.Job.Output[0].TextID != "therapy_session_ended" {
		t.Fatalf("expected therapy_session_ended first, got %v", session.Job.Output)
	}
}

func TestHandleSession_ForwardDuringActive(t *testing.T) {
	// Fake endpoint
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Therapist reply"))
	}))
	defer ts.Close()
	os.Setenv("CAPY_THERAPYSESSION_URL", ts.URL)
	defer os.Unsetenv("CAPY_THERAPYSESSION_URL")

	ctx := context.Background()
	future := time.Now().Add(5 * time.Minute)
	locale := "en"
	user := &database.User{TherapySessionEndAt: &future, IsTyping: true, Locale: &locale}
	input := "some text"
	job := &Job{Command: None, LastCommand: TherapySession, Input: &input}
	session := createSession(job, user, nil, &ctx)

	handleSession(session)

	if len(session.Job.Output) == 0 || session.Job.Output[0].TextID != "Therapist reply" {
		t.Fatalf("expected Therapist reply, got %v", session.Job.Output)
	}
}

func TestHandleSession_EndOnOtherCommand(t *testing.T) {
	ctx := context.Background()
	future := time.Now().Add(5 * time.Minute)
	locale := "en"
	user := &database.User{TherapySessionEndAt: &future, IsTyping: true, Locale: &locale}
	job := &Job{Command: Help}
	session := createSession(job, user, nil, &ctx)

	handleSession(session)

	if len(session.Job.Output) < 2 {
		t.Fatalf("expected at least two outputs, got %d", len(session.Job.Output))
	}
	if session.Job.Output[0].TextID != "therapy_session_ended" {
		t.Fatalf("expected therapy_session_ended first, got %s", session.Job.Output[0].TextID)
	}
	if session.Job.Output[1].TextID != "commands_hint" {
		t.Fatalf("expected commands_hint second, got %s", session.Job.Output[1].TextID)
	}
}
