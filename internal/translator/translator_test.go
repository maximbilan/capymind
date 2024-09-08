package translator

import "testing"

func TestTranslator(t *testing.T) {
	locale1 := Locale("en")
	locale2 := Locale("uk")

	want := "Analysis for the past week 🧑‍⚕️\n\n"

	if got := Translate(locale1, "weekly_analysis"); got != want {
		t.Errorf("Translate() = %v, want %v", got, want)
	}

	want = "Аналіз за останній тиждень 🧑‍⚕️\n\n"
	if got := Translate(locale2, "weekly_analysis"); got != want {
		t.Errorf("Translate() = %v, want %v", got, want)
	}
}

func TestTranslatorJSON(t *testing.T) {
	en_texts := translations["en"]
	uk_texts := translations["uk"]

	if len(en_texts) != len(uk_texts) {
		t.Errorf("Number of texts for en and uk locales is different")
	}
}
