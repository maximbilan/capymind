package analysis

import (
	"testing"

	"github.com/capymind/internal/translator"
)

func TestPrompts(t *testing.T) {
	system1, user1 := getPrompt(WeeklyAnalysis)
	if system1 != "ai_weekly_analysis_system_message" {
		t.Fatalf("Expected ai_weekly_analysis_system_message, got %s", system1)
	}
	if user1 != "ai_weekly_analysis_user_message" {
		t.Fatalf("Expected ai_weekly_analysis_user_message, got %s", user1)
	}

	system2, user2 := getPrompt(SleepAnalysis)
	if system2 != "ai_sleep_analysis_system_message" {
		t.Fatalf("Expected ai_sleep_analysis_system_message, got %s", system2)
	}
	if user2 != "ai_sleep_analysis_user_message" {
		t.Fatalf("Expected ai_sleep_analysis_user_message, got %s", user2)
	}

	system3, user3 := getPrompt(QuickAnalysis)
	if system3 != "ai_quick_analysis_system_message" {
		t.Fatalf("Expected ai_quick_analysis_system_message, got %s", system3)
	}
	if user3 != "ai_quick_analysis_user_message" {
		t.Fatalf("Expected ai_quick_analysis_user_message, got %s", user3)
	}
}

func TestLocalizedPrompts(t *testing.T) {
	locale1 := translator.Locale("en")
	prompt1 := getLocalizedPrompt(QuickAnalysis, locale1)
	if prompt1.System != "You are a therapist at CapyMind, specializing in analyzing user journal entries and providing insights to support their mental well-being." {
		t.Fatalf("Expected 'You are a therapist at CapyMind, specializing in analyzing user journal entries and providing insights to support their mental well-being.', got %s", prompt1.System)
	}

	locale2 := translator.Locale("uk")
	prompt2 := getLocalizedPrompt(SleepAnalysis, locale2)
	if prompt2.User != "Нижче наведено записи з мого останнього сну. Будь ласка, надайте відгук. Записи: " {
		t.Fatalf("Expected 'Нижче наведено записи з мого останнього сну. Будь ласка, надайте відгук. Записи: ', got %s", prompt2.User)
	}
}
