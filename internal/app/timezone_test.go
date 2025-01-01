package app

import (
	"testing"
	"time"

	"github.com/capymind/internal/database"
	"github.com/capymind/internal/mocks"
)

func TestTimezoneHandler(t *testing.T) {
	session := createSession(&Job{Command: "/timezone"}, &database.User{}, nil, nil)
	settingsStorage := &mocks.EmptySettingsStorageMock{}
	handleTimezone(session, settingsStorage)

	if session.Job.Output[0].TextID != "timezone_select" {
		t.Error("Expected 'timezone_select', got", session.Job.Output[0].TextID)
	}
	if len(session.Job.Output[0].Buttons) != 25 {
		t.Error("Expected '25', got", len(session.Job.Output[0].Buttons))
	}
}

func TestTimezoneHandlerWithParam(t *testing.T) {
	session := createSession(&Job{Command: "/timezone 7200", Parameters: []string{"7200"}}, &database.User{}, &database.Settings{}, nil)
	settingsStorage := &mocks.EmptySettingsStorageMock{}
	handleTimezone(session, settingsStorage)

	if session.Job.Output[0].TextID != "timezone_set" {
		t.Error("Expected 'timezone_set', got", session.Job.Output[0].TextID)
	}
	if *session.User.SecondsFromUTC != 7200 {
		t.Error("Expected '7200', got", false)
	}
	if session.Job.Output[1].TextID != "welcome" {
		t.Error("Expected 'welcome', got", session.Job.Output[0].TextID)
	}
}

func TestTimezoneHandlerWithParamOnboarded(t *testing.T) {
	session := createSession(&Job{Command: "/timezone 0", Parameters: []string{"0"}}, &database.User{IsOnboarded: true}, &database.Settings{}, nil)
	settingsStorage := &mocks.EmptySettingsStorageMock{}
	handleTimezone(session, settingsStorage)

	if session.Job.Output[0].TextID != "timezone_set" {
		t.Error("Expected 'timezone_set', got", session.Job.Output[0].TextID)
	}
	if *session.User.SecondsFromUTC != 0 {
		t.Error("Expected '0', got", false)
	}
}

func TestRequestTimezoneHandler(t *testing.T) {
	session := createSession(&Job{Command: "/ask_for_city"}, &database.User{}, nil, nil)
	requestTimezone(session)

	if session.Job.Output[0].TextID != "ask_for_city" {
		t.Error("Expected 'ask_for_city', got", session.Job.Output[0].TextID)
	}
	if session.User.IsTyping != true {
		t.Error("Expected 'true', got", session.User.IsTyping)
	}
}

func TestFinishCityInvalidRequestHandler(t *testing.T) {
	city := "new york"
	session := createSession(&Job{Command: "", Input: &city}, &database.User{}, &database.Settings{}, nil)
	mapsService := &mocks.InvalidMapsServiceMock{}
	settingsStorage := &mocks.EmptySettingsStorageMock{}
	finishCityRequest(session, mapsService, settingsStorage)

	if session.User.IsTyping != false {
		t.Error("Expected 'false', got", session.User.IsTyping)
	}
	if session.Job.Output[0].TextID != "timezone_not_found" {
		t.Error("Expected 'timezone_not_found', got", session.Job.Output[0].TextID)
	}
	if session.Job.Output[1].TextID != "timezone_select" {
		t.Error("Expected 'timezone_select', got", session.Job.Output[1].TextID)
	}
}

func TestFinishCityRequestHandler(t *testing.T) {
	city := "portland"
	session := createSession(&Job{Command: "", Input: &city}, &database.User{}, &database.Settings{}, nil)
	mapsService := &mocks.MapsServiceMock{}
	settingsStorage := &mocks.EmptySettingsStorageMock{}
	finishCityRequest(session, mapsService, settingsStorage)

	if session.User.IsTyping != false {
		t.Error("Expected 'false', got", session.User.IsTyping)
	}
	if *session.Settings.Location != city {
		t.Error("Expected 'portland', got", *session.Settings.Location)
	}
	if len(session.Job.Output[0].Buttons) != 2 {
		t.Error("Expected '2', got", len(session.Job.Output[0].Buttons))
	}
	if session.Job.Output[0].Buttons[0].TextID != "yes" {
		t.Error("Expected 'yes', got", session.Job.Output[0].Buttons[0].TextID)
	}
	if session.Job.Output[0].Buttons[0].Callback != "/timezone 7200" {
		t.Error("Expected '/timezone 7200', got", session.Job.Output[0].Buttons[0].Callback)
	}
	if session.Job.Output[0].Buttons[1].TextID != "no" {
		t.Error("Expected 'no', got", session.Job.Output[0].Buttons[1].TextID)
	}
	if session.Job.Output[0].Buttons[1].Callback != "/timezone" {
		t.Error("Expected '/timezone', got", session.Job.Output[0].Buttons[1].Callback)
	}
}

func TestCurrentTimeString(t *testing.T) {
	time := time.Date(2021, 1, 1, 9, 35, 0, 0, time.UTC)

	timeString := currentTimeString(time, 7200)
	if timeString != "11:35" {
		t.Error("Expected '11:35', got", timeString)
	}
}
