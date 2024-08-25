package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var baseURL string

func init() {
	baseURL = "https://api.telegram.org/bot" + os.Getenv("CAPY_TELEGRAM_BOT_TOKEN")
}

func Parse(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming update %s", err.Error())
		return nil, err
	}
	return &update, nil
}

func SendMessage(chatId int, text string) (string, error) {
	log.Printf("Sending %s to chat_id: %d", text, chatId)
	var telegramApi string = baseURL + "/sendMessage"
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s", bodyString)

	return bodyString, nil
}

func SendMessageWithButtons(chatID int, text string, replyMarkup InlineKeyboardMarkup) error {
	var url string = baseURL + "/sendMessage"

	message := SendMessageRequest{
		ChatID:      chatID,
		Text:        text,
		ReplyMarkup: replyMarkup,
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Printf("Telegram response: %s", body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message: %s", resp.Status)
	}

	return nil
}
