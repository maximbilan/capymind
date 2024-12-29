package translator

import "testing"

func TestLanguage(t *testing.T) {
	if English.String() != "English 🇺🇸" {
		t.Fatalf("Expected 'English 🇺🇸', got %s", English.String())
	}
	if Ukrainian.String() != "Українська 🇺🇦" {
		t.Fatalf("Expected 'Українська 🇺🇦', got %s", Ukrainian.String())
	}
}
