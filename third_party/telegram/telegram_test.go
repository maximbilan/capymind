package telegram

import (
	"bytes"
	"io"
	"testing"
)

func TestValidUser(t *testing.T) {
	const json = `{
		"update_id": 272414882,
		"message": {
			"message_id": 1238,
			"from": {
				"id": 630456345,
				"is_bot":false,
				"first_name": "Max",
				"username": "m4x",
				"language_code": "en"
			},
			"chat": {
				"id": 630456345,
				"first_name": "Max",
				"username": "m4x",
				"type": "private"
			},
			"date": 1735397527,
			"text": "/total_active_user_count"
		}
	}`

	body := io.NopCloser(bytes.NewReader([]byte(json)))

	client := Telegram{}
	message := client.Parse(body)
	if message.UserID != "630456345" {
		t.Errorf("Expected UserID to be 630456345, got %s", message.UserID)
	}
	if message.ChatID != 630456345 {
		t.Errorf("Expected ChatID to be 630456345, got %d", message.ChatID)
	}
	if message.UserName != "m4x" {
		t.Errorf("Expected UserName to be m4x, got %s", message.UserName)
	}
	if message.FirstName != "Max" {
		t.Errorf("Expected FirstName to be Max, got %s", message.FirstName)
	}
	if message.LastName != "" {
		t.Errorf("Expected LastName to be empty, got %s", message.LastName)
	}
	if message.LanguageCode != "en" {
		t.Errorf("Expected LanguageCode to be en, got %s", message.LanguageCode)
	}
	if message.Text != "/total_active_user_count" {
		t.Errorf("Expected Text to be /total_active_user_count, got %s", message.Text)
	}
}

func TestCallback(t *testing.T) {
	const json = `{
		"update_id": 272414882,
		"callback_query": {
			"id": "1238",
			"from": {
				"id": 630456345,
				"is_bot":false,
				"first_name": "Max",
				"username": "m4x",
				"language_code": "en"
			},
			"message": {
				"message_id": 1238,
				"from": {
					"id": 630456345,
					"is_bot":false,
					"first_name": "Max",
					"username": "m4x",
					"language_code": "en"
				},
				"chat": {
					"id": 630456345,
					"first_name": "Max",
					"username": "m4x",
					"type": "private"
				},
				"date": 1735397527,
				"text": "Language"
			},
			"date": 1735397527,
			"data": "/language"
		}
	}`

	body := io.NopCloser(bytes.NewReader([]byte(json)))

	client := Telegram{}
	message := client.Parse(body)
	if message.UserID != "630456345" {
		t.Errorf("Expected UserID to be 630456345, got %s", message.UserID)
	}
	if message.ChatID != 630456345 {
		t.Errorf("Expected ChatID to be 630456345, got %d", message.ChatID)
	}
	if message.UserName != "m4x" {
		t.Errorf("Expected UserName to be m4x, got %s", message.UserName)
	}
	if message.FirstName != "Max" {
		t.Errorf("Expected FirstName to be Max, got %s", message.FirstName)
	}
	if message.LastName != "" {
		t.Errorf("Expected LastName to be empty, got %s", message.LastName)
	}
	if message.LanguageCode != "en" {
		t.Errorf("Expected LanguageCode to be en, got %s", message.LanguageCode)
	}
	if message.Text != "/language" {
		t.Errorf("Expected Text to be /total_active_user_count, got %s", message.Text)
	}
}
