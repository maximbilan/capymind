package bot

import (
	"testing"

	"github.com/capymind/internal/telegram"
)

func TestJobFromMessage(t *testing.T) {
	update := telegram.Update{
		ID: 789,
		Message: &telegram.Message{
			ID:   101,
			Text: "/language en",
			Chat: &telegram.Chat{
				ID: 456,
			},
			From: &telegram.User{
				ID:           456,
				UserName:     "test",
				FirstName:    "Test",
				LastName:     "User",
				LanguageCode: "en",
			},
		},
	}

	job := createJob(update)
	if job == nil {
		t.Fatalf("Job is nil")
	}
	if job.Command != Language {
		t.Fatalf("Command is not language")
	}
	if len(job.Parameters) == 0 {
		t.Fatalf("Parameters are not empty")
	}
	if job.Parameters[0] != "en" {
		t.Fatalf("Parameter is not en")
	}
	if *job.Input != "/language en" {
		t.Fatalf("Input is not /language en")
	}
}

func TestJobFromCallbackQuery(t *testing.T) {
	update := telegram.Update{
		ID: 789,
		CallbackQuery: &telegram.CallbackQuery{
			ID:   "123",
			Data: "/timezone 25200",
			From: &telegram.User{
				ID:           456,
				UserName:     "test",
				FirstName:    "Test",
				LastName:     "User",
				LanguageCode: "en",
			},
			Message: &telegram.Message{
				ID: 101,
				Chat: &telegram.Chat{
					ID: 456,
				},
			},
		},
	}

	job := createJob(update)
	if job == nil {
		t.Fatalf("Job is nil")
	}
	if job.Command != Timezone {
		t.Fatalf("Command is not timezone")
	}
	if len(job.Parameters) == 0 {
		t.Fatalf("Parameters are not empty")
	}
	if job.Parameters[0] != "25200" {
		t.Fatalf("Parameter is not 25200")
	}
	if *job.Input != "/timezone 25200" {
		t.Fatalf("Input is not /timezone 25200")
	}
}
