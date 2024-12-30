package scheduler

import (
	"testing"

	"github.com/capymind/internal/taskservice"
)

func TestPrepareMorningMessage(t *testing.T) {
	scheduledMessage := taskservice.ScheduledTask{
		ChatID: 1,
		Text:   "Bot result",
		Type:   taskservice.Morning,
		Locale: "en",
	}

	result := prepareBotResult(scheduledMessage)
	if result.TextID != "Bot result" {
		t.Error("Expected Bot result, got nil")
	}
	if result.Locale != "en" {
		t.Error("Expected en, got nil")
	}
	if result.Buttons[0].TextID != "make_record_to_journal" {
		t.Error("Expected make_record_to_journal, got nil")
	}
	if result.Buttons[0].Locale != "en" {
		t.Error("Expected en, got nil")
	}
	if result.Buttons[0].Callback != "/note" {
		t.Error("Expected /note, got nil")
	}
}

func TestPrepareFeedbackMessage(t *testing.T) {
	scheduledMessage := taskservice.ScheduledTask{
		ChatID: 1,
		Text:   "Bot result",
		Type:   taskservice.Feedback,
		Locale: "en",
	}

	result := prepareBotResult(scheduledMessage)
	if result.TextID != "Bot result" {
		t.Error("Expected Bot result, got nil")
	}
	if result.Locale != "en" {
		t.Error("Expected en, got nil")
	}
	if result.Buttons[0].TextID != "feedback_button" {
		t.Error("Expected feedback_button, got nil")
	}
	if result.Buttons[0].Locale != "en" {
		t.Error("Expected en, got nil")
	}
	if result.Buttons[0].Callback != "/support" {
		t.Error("Expected /support, got nil")
	}
}

func TestPrepareRegularMessage(t *testing.T) {
	scheduledMessage := taskservice.ScheduledTask{
		ChatID: 1,
		Text:   "Bot result",
		Type:   taskservice.Regular,
		Locale: "en",
	}

	result := prepareBotResult(scheduledMessage)
	if result.TextID != "Bot result" {
		t.Error("Expected Bot result, got nil")
	}
	if result.Locale != "en" {
		t.Error("Expected en, got nil")
	}
	if len(result.Buttons) != 0 {
		t.Error("Expected no buttons, got some")
	}
}
