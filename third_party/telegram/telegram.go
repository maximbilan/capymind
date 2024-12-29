package telegram

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/capymind/internal/botservice"
)

type Telegram struct{}

var baseURL string

//coverage:ignore
func init() {
	baseURL = "https://api.telegram.org/bot" + os.Getenv("CAPY_TELEGRAM_BOT_TOKEN")
}

// Parse an incoming update
func (t Telegram) Parse(body io.ReadCloser) *botservice.BotMessage {
	var update Update
	if err := json.NewDecoder(body).Decode(&update); err != nil {
		log.Printf("[Parse] Could not decode incoming update %s", err.Error())
		return nil
	}

	// Check if the message is valid
	if update.Message == nil && update.CallbackQuery == nil {
		log.Printf("[User] Invalid update: %d", update.ID)
		return nil
	}
	// Check if the message is valid (Callback variant)
	if update.Message == nil && update.CallbackQuery != nil && update.CallbackQuery.Message == nil {
		log.Printf("[User] Invalid update (with callback): %d", update.ID)
		return nil
	}

	var chatID int64
	var telegramUser *User

	// Check if the update is a callback query or a message
	callbackQuery := update.CallbackQuery
	if callbackQuery != nil && callbackQuery.Data != "" {
		chatID = callbackQuery.Message.Chat.ID
		telegramUser = callbackQuery.From
	} else {
		message := update.Message
		chatID = message.Chat.ID
		telegramUser = message.From
	}

	// Check if the user is valid
	if telegramUser == nil || telegramUser.ID == 0 {
		return nil
	}

	// Get the input from the update
	var input *string
	if callbackQuery != nil && callbackQuery.Data != "" {
		input = &callbackQuery.Data
	} else if update.Message != nil {
		message := update.Message
		input = &message.Text
	}

	// Check if the input is valid
	if input == nil || *input == "" {
		return nil
	}

	// Create a user from the telegram user
	message := botservice.BotMessage{
		UserID:       telegramUser.StringID(),
		ChatID:       chatID,
		UserName:     telegramUser.UserName,
		FirstName:    telegramUser.FirstName,
		LastName:     telegramUser.LastName,
		LanguageCode: telegramUser.LanguageCode,
		Text:         *input,
	}

	return &message
}

//coverage:ignore
func (t Telegram) SendMessage(chatID int64, text string) {
	sendMessage(chatID, text, nil)
}

//coverage:ignore
func (t Telegram) SendResult(chatID int64, result botservice.BotResult) {
	// Prepare the reply markup
	var replyMarkup *InlineKeyboardMarkup
	if len(result.Buttons) > 0 {
		var inlineKeyboard [][]InlineKeyboardButton
		for _, button := range result.Buttons {
			inlineKeyboard = append(inlineKeyboard, []InlineKeyboardButton{
				{Text: button.Text(), CallbackData: &button.Callback},
			})
		}

		replyMarkup = &InlineKeyboardMarkup{
			InlineKeyboard: inlineKeyboard,
		}
	}

	// Send the message
	sendMessage(chatID, result.Text(), replyMarkup)
}

//coverage:ignore
func sendMessage(chatID int64, text string, replyMarkup *InlineKeyboardMarkup) {
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
}
