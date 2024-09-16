package translator

import "testing"

func TestTranslator(t *testing.T) {
	locale1 := Locale("en")
	locale2 := Locale("uk")

	want := "Latest analysis 🧑‍⚕️\n\n"

	if got := Translate(locale1, "weekly_analysis"); got != want {
		t.Errorf("Translate() = %v, want %v", got, want)
	}

	want = "Аналіз за останній час 🧑‍⚕️\n\n"
	if got := Translate(locale2, "weekly_analysis"); got != want {
		t.Errorf("Translate() = %v, want %v", got, want)
	}

	// Test for missing translation
	want = "no_existing_id"
	if got := Translate(locale1, "no_existing_id"); got != want {
		t.Errorf("Translate() = %v, want %v", got, want)
	}

	// Test prompt
	want = "Below is a list of my recent journal entries. Please provide feedback: "
	if got := Prompt(locale1, "ai_analysis_user_message"); got != want {
		t.Errorf("Prompt() = %v, want %v", got, want)
	}
}

func TestTranslatorJSON(t *testing.T) {
	en_texts := translations["en"]
	uk_texts := translations["uk"]

	if len(en_texts) != len(uk_texts) {
		t.Errorf("Number of texts for en and uk locales is different")
	}
}

func TestPromptJSON(t *testing.T) {
	en_prompts := prompts["en"]
	uk_prompts := prompts["uk"]

	if len(en_prompts) != len(uk_prompts) {
		t.Errorf("Number of prompts for en and uk locales is different")
	}
}
