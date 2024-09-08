package telegram

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

var baseURL string

func init() {
	baseURL = "https://api.telegram.org/bot" + os.Getenv("CAPY_TELEGRAM_BOT_TOKEN")
}

func Parse(r *http.Request) *Update {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("[Parse] Could not decode incoming update %s", err.Error())
		return nil
	}
	return &update
}

func SendMessage(chatID int64, text string, replyMarkup *InlineKeyboardMarkup) {
	var url string = baseURL + "/sendMessage"

	message := SendMessageRequest{
		ChatID:      chatID,
		Text:        text,
		ReplyMarkup: replyMarkup,
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Printf("[SendMessage] JSON parsing error: %s", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("[SendMessage] POST error: %s", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[SendMessage] Body reading error: %s", err)
		return
	}
	log.Printf("[SendMessage] Response body: %s", body)
}
