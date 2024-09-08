package utils

import (
	"testing"
)

func TestTimezoneParameter(t *testing.T) {
	want := "timezone_3600"

	info := TimeZoneInfo{
		Offset:         1,
		Description:    "Berlin, Paris, Rome",
		SecondsFromUTC: 3600,
	}
	param := GetTimezoneParameter(info)

	if want != param {
		t.Fatalf("Incorrect parameter. %s doesn't equal %s", want, param)
	}
}

func TestParseTimezone(t *testing.T) {
	want, ok := ParseTimezone("timezone_3600")
	if !ok || *want != 3600 {
		t.Fatalf("Incorrect timezone. %v doesn't equal %v", want, 3600)
	}

	want, ok = ParseTimezone("timezone_abc")
	if ok {
		t.Fatalf("Incorrect timezone. %v is not nil", want)
	}

	want, ok = ParseTimezone("timezone")
	if ok {
		t.Fatalf("Incorrect timezone. %v is not nil", want)
	}

	want, ok = ParseTimezone("timezone_")
	if ok {
		t.Fatalf("Incorrect timezone. %v is not nil", want)
	}

	want, ok = ParseTimezone("timezone_0")
	if !ok || *want != 0 {
		t.Fatalf("Incorrect timezone. %v doesn't equal %v", want, 0)
	}
}

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
