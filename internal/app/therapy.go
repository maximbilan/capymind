package app

import (
    "bytes"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/capymind/internal/translator"
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
    url := translator.Prompt(session.Locale(), "therapysession_url")
    if url == "therapysession_url" {
        // Fallback to env var if not set in prompts
        u := os.Getenv("CAPY_THERAPYSESSION_URL")
        if u == "" {
            return nil
        }
        url = u
    }

    payload := fmt.Sprintf(`{"user_id":"%s","message":%q}`, session.User.ID, text)
    resp, err := httpPostJSON(url, payload)
    if err != nil {
        log.Printf("[TherapySession] request error: %v", err)
        return nil
    }
    if resp == "" {
        return nil
    }
    return &resp
}

func httpPostJSON(url string, payload string) (string, error) {
    //coverage:ignore
    req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
    if err != nil {
        return "", err
    }
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{Timeout: 15 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    return string(body), nil
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
