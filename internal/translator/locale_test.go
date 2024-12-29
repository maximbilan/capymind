package translator

import "testing"

func TestLocale(t *testing.T) {
	if EN.String() != "en" {
		t.Fatalf("Expected 'en', got %s", EN.String())
	}
	if UK.String() != "uk" {
		t.Fatalf("Expected 'uk', got %s", UK.String())
	}
}
