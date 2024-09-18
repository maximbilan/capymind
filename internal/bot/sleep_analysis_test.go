package bot

import (
	"testing"

	"github.com/capymind/internal/translator"
)

func TestCheckIfNoteADream(t *testing.T) {
	text1 := "I had a dream last night"
	locale1 := translator.Locale("en")
	if !checkIfNoteADream(text1, locale1) {
		t.Fatalf("Expected true, got false")
	}

	text2 := "Снилось щось цікаве"
	locale2 := translator.Locale("uk")
	if !checkIfNoteADream(text2, locale2) {
		t.Fatalf("Expected true, got false")
	}
}
