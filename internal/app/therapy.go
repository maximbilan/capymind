package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
    "strings"
	"time"

	"github.com/google/uuid"
)

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
	// Resolve base URL and token
	baseURL := os.Getenv("CAPY_THERAPY_SESSION_URL")
	if baseURL == "" {
		log.Printf("[TherapySession] missing CAPY_THERAPY_SESSION_URL")
		return nil
	}
	token := os.Getenv("CAPY_AGENT_TOKEN")
	if token == "" {
		log.Printf("[TherapySession] missing CAPY_AGENT_TOKEN")
		return nil
	}

	// Ensure therapy session identifier on the user entity
	if session.User.TherapySessionId == nil || *session.User.TherapySessionId == "" {
		newID := uuid.NewString()
		session.User.TherapySessionId = &newID
	}

	userID := session.User.ID
	therapySessionID := *session.User.TherapySessionId
	locale := session.Locale().String()

	client := &http.Client{Timeout: 15 * time.Second}

	// 1) Create/init the therapy session
	initURL := fmt.Sprintf("%s/apps/capymind_agent/users/%s/sessions/%s", baseURL, userID, therapySessionID)
	initBody := map[string]any{
		"state": map[string]any{
			"preferred_language": locale,
		},
	}
	initBodyBytes, _ := json.Marshal(initBody)
	initReq, err := http.NewRequest("POST", initURL, bytes.NewBuffer(initBodyBytes))
	if err != nil {
		log.Printf("[TherapySession] init request build error: %v", err)
		return nil
	}
	initReq.Header.Set("Authorization", "Bearer "+token)
	initReq.Header.Set("Content-Type", "application/json")

    initResp, err := client.Do(initReq)
    if err != nil {
        log.Printf("[TherapySession] init request error: %v", err)
        return nil
    }
    defer initResp.Body.Close()
    proceed := false
    if initResp.StatusCode >= 200 && initResp.StatusCode < 300 {
        proceed = true
    } else {
        body, _ := io.ReadAll(initResp.Body)
        // Allow existing session scenario to proceed
        if initResp.StatusCode == 400 && strings.Contains(string(body), "Session already exists") {
            log.Printf("[TherapySession] init session exists, proceeding: %s", therapySessionID)
            proceed = true
        } else {
            log.Printf("[TherapySession] init non-2xx: %d body=%s", initResp.StatusCode, string(body))
        }
    }
    if !proceed {
        return nil
    }

	// 2) Send user message via run_sse
	runURL := fmt.Sprintf("%s/run_sse", baseURL)
	runBody := map[string]any{
		"app_name":   "capymind_agent",
		"user_id":    userID,
		"session_id": therapySessionID,
		"new_message": map[string]any{
			"role": "user",
			"parts": []map[string]string{
				{"text": text},
			},
		},
		"streaming": false,
	}
	runBodyBytes, _ := json.Marshal(runBody)
	runReq, err := http.NewRequest("POST", runURL, bytes.NewBuffer(runBodyBytes))
	if err != nil {
		log.Printf("[TherapySession] run request build error: %v", err)
		return nil
	}
	runReq.Header.Set("Authorization", "Bearer "+token)
	runReq.Header.Set("Content-Type", "application/json")

	runResp, err := client.Do(runReq)
	if err != nil {
		log.Printf("[TherapySession] run request error: %v", err)
		return nil
	}
	defer runResp.Body.Close()
	runRespBody, err := io.ReadAll(runResp.Body)
	if err != nil {
		log.Printf("[TherapySession] run read error: %v", err)
		return nil
	}
	if runResp.StatusCode < 200 || runResp.StatusCode >= 300 {
		log.Printf("[TherapySession] run non-2xx: %d body=%s", runResp.StatusCode, string(runRespBody))
		return nil
	}
    respStr := string(runRespBody)
    if respStr == "" {
        return nil
    }

    // Try to extract plain text from JSON response
    // Support responses that are either raw JSON or lines prefixed with "data: "
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

// Relay a user message to the therapy session backend and append the reply
func relayTherapyMessage(text string, session *Session) {
	//coverage:ignore
	// Send immediate typing acknowledgement is already enabled via IsTyping
	reply := callTherapySessionEndpoint(text, session)
	if reply != nil && *reply != "" {
		setOutputRawText(*reply, session)
	}
}

// End the therapy session and notify the user
func endTherapySession(session *Session) {
	session.User.IsTyping = false
	session.User.TherapySessionEndAt = nil
	setOutputText("therapy_session_ended", session)
}
