package scheduler

import (
	"net/url"
	"testing"

	"github.com/capymind/internal/taskservice"
)

func TestParse(t *testing.T) {
	url1 := url.URL{
		RawQuery: "type=morning&offset=2",
	}

	typeStr, offset := parse(&url1)
	if *typeStr != "morning" {
		t.Errorf("Expected type=morning, got %s", *typeStr)
	}
	if offset != 2 {
		t.Errorf("Expected offset=2, got %d", offset)
	}

	url2 := url.URL{
		RawQuery: "type=evening&offset=0",
	}

	typeStr, offset = parse(&url2)
	if *typeStr != "evening" {
		t.Errorf("Expected type=evening, got %s", *typeStr)
	}
	if offset != 0 {
		t.Errorf("Expected offset=0, got %d", offset)
	}

	url3 := url.URL{
		RawQuery: "type=weekly_analysis&offset=-1",
	}

	typeStr, offset = parse(&url3)
	if *typeStr != "weekly_analysis" {
		t.Errorf("Expected type=weekly_analysis, got %s", *typeStr)
	}
	if offset != -1 {
		t.Errorf("Expected offset=-1, got %d", offset)
	}

	url4 := url.URL{
		RawQuery: "type=aaa&",
	}

	typeStr, offset = parse(&url4)
	if *typeStr != "aaa" {
		t.Errorf("Expected nil, got %s", *typeStr)
	}
	if offset != 0 {
		t.Errorf("Expected offset=0, got %d", offset)
	}

	url5 := url.URL{
		RawQuery: "offset=2",
	}

	typeStr, offset = parse(&url5)
	if *typeStr != "" {
		t.Errorf("Expected nil, got %s", *typeStr)
	}
	if offset != 2 {
		t.Errorf("Expected offset=2, got %d", offset)
	}
}

func TestGetTextMessage(t *testing.T) {
	message := getTextMessage(taskservice.Morning)
	if *message != "how_are_you_morning_monday" {
		t.Errorf("Expected how_are_you_morning_monday, got %s", *message)
	}

	message = getTextMessage(taskservice.Evening)
	if *message != "how_are_you_evening_monday" {
		t.Errorf("Expected how_are_you_evening_monday, got %s", *message)
	}

	message = getTextMessage(taskservice.Feedback)
	if *message != "ask_write_review_about_bot" {
		t.Errorf("Expected ask_write_review_about_bot, got %s", *message)
	}

	message = getTextMessage(taskservice.WeeklyAnalysis)
	if *message != "" {
		t.Errorf("Expected empty string, got %s", *message)
	}

	message = getTextMessage(taskservice.UserStats)
	if *message != "" {
		t.Errorf("Expected empty string, got %s", *message)
	}

	message = getTextMessage(taskservice.AdminStats)
	if *message != "" {
		t.Errorf("Expected empty string, got %s", *message)
	}

	message = getTextMessage("")
	if message != nil {
		t.Errorf("Expected nil, got %s", *message)
	}
}
