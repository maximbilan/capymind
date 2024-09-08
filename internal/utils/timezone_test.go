package utils

import "testing"

func TestTimezoneParameter(t *testing.T) {
	want := "timezone_3600"

	info := TimeZoneInfo{
		Offset:         1,
		Description:    "Berlin, Paris, Rome",
		SecondsFromUTC: 3600,
	}
	param := GetTimezoneParameter(info)

	if want != param {
		t.Fatalf("Wrong parameter. %s doesn't equal %s", want, param)
	}
}
