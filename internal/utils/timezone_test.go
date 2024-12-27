package utils

import (
	"testing"
)

func TestTimeZones(t *testing.T) {
	list := GetTimeZones()
	if len(list) != 25 {
		t.Fatalf("Expected 25, got %d", len(list))
	}

	if list[0].Offset != -11 {
		t.Fatalf("Expected -12, got %d", list[0].Offset)
	}

	if list[15].Description != "GMT +4" {
		t.Fatalf("Expected GMT +4, got %s", list[15].Description)
	}

	if list[24].SecondsFromUTC != 46800 {
		t.Fatalf("Expected 43200, got %d", list[24].SecondsFromUTC)
	}

	if list[5].Parameter() != "-21600" {
		t.Fatalf("Expected -7, got %s", list[5].Parameter())
	}
}
