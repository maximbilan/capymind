package utils

import (
	"testing"
)

func TestTimeZones(t *testing.T) {
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

	if list[3].String() != "UTC -9 - Alaska" {
		t.Fatalf("Expected UTC -9 - Alaska, got %s", list[3].String())
	}

	if list[5].Parameter() != "-25200" {
		t.Fatalf("Expected -7, got %s", list[5].Parameter())
	}
}
