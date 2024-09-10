package utils

import (
	"testing"
)

func TestGetTimeZones(t *testing.T) {
	list := GetTimeZones()
	if len(list) != 25 {
		t.Fatalf("Expected 25, got %d", len(list))
	}

	if list[0].Offset != -12 {
		t.Fatalf("Expected -12, got %d", list[0].Offset)
	}

	if list[15].Description != "Kyiv, Istanbul, Helsinki" {
		t.Fatalf("Expected Kyiv, Istanbul, Helsinki, got %s", list[15].Description)
	}

	if list[24].SecondsFromUTC != 43200 {
		t.Fatalf("Expected 43200, got %d", list[24].SecondsFromUTC)
	}
}
