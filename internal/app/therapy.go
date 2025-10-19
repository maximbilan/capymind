package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"google.golang.org/api/idtoken"
)

// Allow tests to inject a custom HTTP client builder.
var newTherapyHTTPClient = buildTherapyHTTPClient

// Internal configuration for therapy session calls
type therapyConfig struct {
	baseURL   string
	userID    string
	sessionID string
	locale    string
	ctx       context.Context
	token     string
}

// Start therapy session
func startTherapySession(session *Session) {
	setOutputText("start_therapy_session", session)
	session.User.IsTyping = true
	endAt := time.Now().Add(30 * time.Minute)
	session.User.TherapySessionEndAt = &endAt
}

// HTTP client to therapy session endpoint
func callTherapySessionEndpoint(text string, session *Session) *string {
	//coverage:ignore
	cfg := buildTherapyConfig(session)
	if cfg == nil {
		return nil
	}

	client, err := newTherapyHTTPClient(cfg.ctx, cfg.baseURL)
	if err != nil {
		log.Printf("[TherapySession] failed to create authenticated client: %v", err)
		return nil
	}

	if !initTherapySession(client, cfg) {
		return nil
	}

	return sendTherapyMessage(client, cfg, text)
}

// Relay a user message to the therapy session backend and append the reply
func relayTherapyMessage(text string, session *Session) {
	//coverage:ignore
	// Send immediate typing acknowledgement is already enabled via IsTyping
	reply := callTherapySessionEndpoint(text, session)
	if reply != nil && *reply != "" {
		setOutputRawText(*reply, session)
	}
}

// Build an HTTP client configured with ID token authentication for therapy session calls
func buildTherapyHTTPClient(ctx context.Context, targetURL string) (*http.Client, error) {
	// Local development path: plain HTTP client; per-request header is set using cfg.token
	if os.Getenv("CLOUD") == "false" {
		return &http.Client{Timeout: 120 * time.Second}, nil
	}

	// Cloud path (default): use Google ID token auth
	client, err := idtoken.NewClient(ctx, targetURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create ID token client: %w", err)
	}
	// Set timeout on the client
	client.Timeout = 120 * time.Second
	return client, nil
}

// Build configuration from environment and session; ensures a session ID exists
func buildTherapyConfig(session *Session) *therapyConfig {
	baseURL := os.Getenv("CAPY_THERAPY_SESSION_URL")
	if baseURL == "" {
		log.Printf("[TherapySession] missing CAPY_THERAPY_SESSION_URL")
		return nil
	}

	if session.User.TherapySessionId == nil || *session.User.TherapySessionId == "" {
		newID := uuid.NewString()
		session.User.TherapySessionId = &newID
	}

	userID := session.User.ID
	sessionID := *session.User.TherapySessionId
	locale := session.Locale().String()
	var token string
	if os.Getenv("CLOUD") == "false" {
		token = strings.TrimSpace(os.Getenv("CAPY_AGENT_TOKEN"))
	}

	return &therapyConfig{
		baseURL:   baseURL,
		userID:    userID,
		sessionID: sessionID,
		locale:    locale,
		ctx:       context.Background(),
		token:     token,
	}
}

// Initialize or validate the therapy session on the backend
func initTherapySession(client *http.Client, cfg *therapyConfig) bool {
	initURL := fmt.Sprintf("%s/apps/capymind_agent/users/%s/sessions/%s", cfg.baseURL, cfg.userID, cfg.sessionID)
	initBody := map[string]any{
		"state": map[string]any{
			"preferred_language": cfg.locale,
		},
	}
	initBodyBytes, _ := json.Marshal(initBody)

	req, err := http.NewRequest("POST", initURL, bytes.NewBuffer(initBodyBytes))
	if err != nil {
		log.Printf("[TherapySession] init request build error: %v", err)
		return false
	}
	// ID token client automatically adds Authorization header
	req.Header.Set("Content-Type", "application/json")
	if cfg.token != "" {
		req.Header.Set("Authorization", "Bearer "+cfg.token)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[TherapySession] init request error: %v", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return true
	}

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == 400 && strings.Contains(string(body), "Session already exists") {
		log.Printf("[TherapySession] init session exists, proceeding: %s", cfg.sessionID)
		return true
	}

	log.Printf("[TherapySession] init non-2xx: %d body=%s", resp.StatusCode, string(body))
	return false
}

// Send the user's message to the run_sse endpoint and parse the reply
func sendTherapyMessage(client *http.Client, cfg *therapyConfig, text string) *string {
	runURL := fmt.Sprintf("%s/run_sse", cfg.baseURL)
	runBody := map[string]any{
		"app_name":   "capymind_agent",
		"user_id":    cfg.userID,
		"session_id": cfg.sessionID,
		"new_message": map[string]any{
			"role": "user",
			"parts": []map[string]string{
				{"text": text},
			},
		},
		"streaming": false,
	}
	runBodyBytes, _ := json.Marshal(runBody)

	req, err := http.NewRequest("POST", runURL, bytes.NewBuffer(runBodyBytes))
	if err != nil {
		log.Printf("[TherapySession] run request build error: %v", err)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")
	if cfg.token != "" {
		req.Header.Set("Authorization", "Bearer "+cfg.token)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[TherapySession] run request error: %v", err)
		return nil
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[TherapySession] run read error: %v", err)
		return nil
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Printf("[TherapySession] run non-2xx: %d body=%s", resp.StatusCode, string(respBody))
		return nil
	}

	respStr := string(respBody)
	if respStr == "" {
		return nil
	}
	return parseRunResponse(respStr)
}

// Parse a run_sse HTTP response body, extracting plain text if present
func parseRunResponse(respStr string) *string {
	extractJSON := func(s string) string {
		s = strings.TrimSpace(s)
		if strings.HasPrefix(s, "data:") {
			// If multiple lines, pick the last data line
			lines := strings.Split(s, "\n")
			for i := len(lines) - 1; i >= 0; i-- {
				line := strings.TrimSpace(lines[i])
				if strings.HasPrefix(line, "data:") {
					return strings.TrimSpace(strings.TrimPrefix(line, "data:"))
				}
			}
			return strings.TrimSpace(strings.TrimPrefix(lines[len(lines)-1], "data:"))
		}
		return s
	}

	type runSseContentPart struct {
		Text string `json:"text"`
	}
	type runSseContent struct {
		Parts []runSseContentPart `json:"parts"`
	}
	type runSseResponse struct {
		Content runSseContent `json:"content"`
	}

	jsonCandidate := extractJSON(respStr)
	var parsed runSseResponse
	if err := json.Unmarshal([]byte(jsonCandidate), &parsed); err == nil {
		if len(parsed.Content.Parts) > 0 && parsed.Content.Parts[0].Text != "" {
			onlyText := parsed.Content.Parts[0].Text
			return &onlyText
		}
	}

	// Fallback: return body as-is
	return &respStr
}

// End the therapy session and notify the user
func endTherapySession(session *Session) {
	session.User.IsTyping = false
	session.User.TherapySessionEndAt = nil
	setOutputText("therapy_session_ended", session)
}
