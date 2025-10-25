package app

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"strings"

	"github.com/capymind/internal/database"
)

func TestStartTherapySession(t *testing.T) {
	ctx := context.Background()
	user := &database.User{IsTyping: false}
	session := createSession(&Job{Command: Therapy}, user, nil, &ctx)

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
	session := createSession(&Job{Command: Therapy}, user, nil, &ctx)

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
	// Create a fake therapy session backend implementing both init and run endpoints
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPost && strings.HasPrefix(r.URL.Path, "/apps/capymind_agent/users/u1/sessions/"):
			// Session init endpoint
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"ok":true}`))
			return
		case r.Method == http.MethodPost && r.URL.Path == "/run_sse":
			// Message sending endpoint
			w.Header().Set("Content-Type", "text/event-stream")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("data: {\"content\":{\"parts\":[{\"text\":\"Hello, I'm here for you.\"}],\"role\":\"model\"}}\n\n"))
			return
		default:
			http.NotFound(w, r)
			return
		}
	}))
	defer ts.Close()
	t.Setenv("CAPY_THERAPY_SESSION_URL", ts.URL)

	// Inject simple HTTP client without Google auth for tests
	originalBuilder := newTherapyHTTPClient
	newTherapyHTTPClient = func(ctx context.Context, targetURL string) (*http.Client, error) {
		return &http.Client{Timeout: 5 * time.Second}, nil
	}
	defer func() { newTherapyHTTPClient = originalBuilder }()

	ctx := context.Background()
	locale := "en"
	user := &database.User{ID: "u1", Locale: &locale}
	session := createSession(&Job{Command: None}, user, nil, &ctx)

	relayTherapyMessage("hi", session)

	if len(session.Job.Output) < 2 {
		t.Fatalf("expected at least two outputs (immediate feedback + response)")
	}
	
	// First output should be a thinking message
	firstOutput := session.Job.Output[0]
	validThinkingKeys := []string{
		"therapy_thinking_1",
		"therapy_thinking_2", 
		"therapy_thinking_3",
		"therapy_thinking_4",
		"therapy_thinking_5",
	}
	
	found := false
	for _, validKey := range validThinkingKeys {
		if firstOutput.TextID == validKey {
			found = true
			break
		}
	}
	
	if !found {
		t.Fatalf("expected first output to be a thinking message, got %s", firstOutput.TextID)
	}
	
	// Second output should be the actual therapy response
	if session.Job.Output[1].TextID != "Hello, I'm here for you." {
		t.Fatalf("unexpected relay text: %s", session.Job.Output[1].TextID)
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
	// Fake backend implementing both init and run endpoints
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPost && strings.HasPrefix(r.URL.Path, "/apps/capymind_agent/users/u1/sessions/"):
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"ok":true}`))
			return
		case r.Method == http.MethodPost && r.URL.Path == "/run_sse":
			w.Header().Set("Content-Type", "text/event-stream")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("data: {\"content\":{\"parts\":[{\"text\":\"Therapist reply\"}],\"role\":\"model\"}}\n\n"))
			return
		default:
			http.NotFound(w, r)
			return
		}
	}))
	defer ts.Close()
	t.Setenv("CAPY_THERAPY_SESSION_URL", ts.URL)

	originalBuilder := newTherapyHTTPClient
	newTherapyHTTPClient = func(ctx context.Context, targetURL string) (*http.Client, error) {
		return &http.Client{Timeout: 5 * time.Second}, nil
	}
	defer func() { newTherapyHTTPClient = originalBuilder }()

	ctx := context.Background()
	future := time.Now().Add(5 * time.Minute)
	locale := "en"
	user := &database.User{ID: "u1", TherapySessionEndAt: &future, IsTyping: true, Locale: &locale}
	input := "some text"
	job := &Job{Command: None, LastCommand: Therapy, Input: &input}
	session := createSession(job, user, nil, &ctx)

	handleSession(session)

	if len(session.Job.Output) < 2 {
		t.Fatalf("expected at least two outputs (immediate feedback + response), got %v", session.Job.Output)
	}
	
	// First output should be a thinking message
	firstOutput := session.Job.Output[0]
	validThinkingKeys := []string{
		"therapy_thinking_1",
		"therapy_thinking_2", 
		"therapy_thinking_3",
		"therapy_thinking_4",
		"therapy_thinking_5",
	}
	
	found := false
	for _, validKey := range validThinkingKeys {
		if firstOutput.TextID == validKey {
			found = true
			break
		}
	}
	
	if !found {
		t.Fatalf("expected first output to be a thinking message, got %s", firstOutput.TextID)
	}
	
	// Second output should be the actual therapy response
	if session.Job.Output[1].TextID != "Therapist reply" {
		t.Fatalf("expected Therapist reply in second output, got %v", session.Job.Output)
	}
}

func TestRelayTherapyMessage_ExistingSessionContinues(t *testing.T) {
	// Fake backend: init returns 400 Session already exists; run_sse returns a reply
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPost && strings.HasPrefix(r.URL.Path, "/apps/capymind_agent/users/u1/sessions/"):
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"detail":"Session already exists: abc-123"}`))
			return
		case r.Method == http.MethodPost && r.URL.Path == "/run_sse":
			w.Header().Set("Content-Type", "text/event-stream")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("data: {\"content\":{\"parts\":[{\"text\":\"Hello again\"}],\"role\":\"model\"}}\n\n"))
			return
		default:
			http.NotFound(w, r)
			return
		}
	}))
	defer ts.Close()
	t.Setenv("CAPY_THERAPY_SESSION_URL", ts.URL)

	originalBuilder := newTherapyHTTPClient
	newTherapyHTTPClient = func(ctx context.Context, targetURL string) (*http.Client, error) {
		return &http.Client{Timeout: 5 * time.Second}, nil
	}
	defer func() { newTherapyHTTPClient = originalBuilder }()

	ctx := context.Background()
	locale := "en"
	user := &database.User{ID: "u1", Locale: &locale}
	session := createSession(&Job{Command: None}, user, nil, &ctx)

	relayTherapyMessage("hi", session)

	if len(session.Job.Output) < 2 {
		t.Fatalf("expected at least two outputs (immediate feedback + response)")
	}
	
	// First output should be a thinking message
	firstOutput := session.Job.Output[0]
	validThinkingKeys := []string{
		"therapy_thinking_1",
		"therapy_thinking_2", 
		"therapy_thinking_3",
		"therapy_thinking_4",
		"therapy_thinking_5",
	}
	
	found := false
	for _, validKey := range validThinkingKeys {
		if firstOutput.TextID == validKey {
			found = true
			break
		}
	}
	
	if !found {
		t.Fatalf("expected first output to be a thinking message, got %s", firstOutput.TextID)
	}
	
	// Second output should be the actual therapy response
	if session.Job.Output[1].TextID != "Hello again" {
		t.Fatalf("unexpected relay text: %s", session.Job.Output[1].TextID)
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

func TestGetRandomThinkingMessage(t *testing.T) {
	ctx := context.Background()
	user := &database.User{ID: "test_user_123"}
	session := createSession(&Job{Command: Therapy}, user, nil, &ctx)

	// Test that we get a valid thinking message key
	messageKey := getRandomThinkingMessage(session)
	
	validKeys := []string{
		"therapy_thinking_1",
		"therapy_thinking_2", 
		"therapy_thinking_3",
		"therapy_thinking_4",
		"therapy_thinking_5",
	}
	
	found := false
	for _, validKey := range validKeys {
		if messageKey == validKey {
			found = true
			break
		}
	}
	
	if !found {
		t.Fatalf("expected one of %v, got %s", validKeys, messageKey)
	}
}

func TestRelayTherapyMessageWithImmediateFeedback(t *testing.T) {
	// Create a mock server that simulates the therapy endpoint
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a successful therapy response
		response := `{"content": {"parts": [{"text": "I understand your concerns. Let's work through this together."}]}}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Set environment variables for testing
	originalURL := os.Getenv("CAPY_THERAPY_SESSION_URL")
	originalCloud := os.Getenv("CLOUD")
	originalToken := os.Getenv("CAPY_AGENT_TOKEN")
	
	os.Setenv("CAPY_THERAPY_SESSION_URL", server.URL)
	os.Setenv("CLOUD", "false")
	os.Setenv("CAPY_AGENT_TOKEN", "test_token")
	
	defer func() {
		os.Setenv("CAPY_THERAPY_SESSION_URL", originalURL)
		os.Setenv("CLOUD", originalCloud)
		os.Setenv("CAPY_AGENT_TOKEN", originalToken)
	}()

	ctx := context.Background()
	user := &database.User{
		ID: "test_user",
		TherapySessionId: stringPtr("test_session_id"),
		TherapySessionEndAt: timePtr(time.Now().Add(30 * time.Minute)),
		IsTyping: true,
	}
	session := createSession(&Job{Command: Therapy}, user, nil, &ctx)

	// Call relayTherapyMessage
	relayTherapyMessage("I'm feeling anxious about work", session)

	// Check that we got at least 2 outputs: immediate feedback + actual response
	if len(session.Job.Output) < 2 {
		t.Fatalf("expected at least 2 outputs, got %d", len(session.Job.Output))
	}

	// Check that the first output is a thinking message
	firstOutput := session.Job.Output[0]
	validThinkingKeys := []string{
		"therapy_thinking_1",
		"therapy_thinking_2", 
		"therapy_thinking_3",
		"therapy_thinking_4",
		"therapy_thinking_5",
	}
	
	found := false
	for _, validKey := range validThinkingKeys {
		if firstOutput.TextID == validKey {
			found = true
			break
		}
	}
	
	if !found {
		t.Fatalf("expected first output to be a thinking message, got %s", firstOutput.TextID)
	}

	// Check that the second output is the actual therapy response
	secondOutput := session.Job.Output[1]
	if !strings.Contains(secondOutput.TextID, "I understand your concerns") {
		t.Fatalf("expected therapy response in second output TextID, got: %s", secondOutput.TextID)
	}
}

// Helper functions for creating test data
func stringPtr(s string) *string {
	return &s
}

func timePtr(t time.Time) *time.Time {
	return &t
}
