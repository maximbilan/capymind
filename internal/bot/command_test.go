package bot

import (
	"testing"
)

func TestCommands(t *testing.T) {
	if Start != "/start" {
		t.Fatalf("Expected /start, got %s", Start)
	}

	if Note != "/note" {
		t.Fatalf("Expected /note, got %s", Note)
	}

	if Last != "/last" {
		t.Fatalf("Expected /last, got %s", Last)
	}

	if Analysis != "/analysis" {
		t.Fatalf("Expected /analysis, got %s", Analysis)
	}

	if Language != "/language" {
		t.Fatalf("Expected /language, got %s", Language)
	}

	if Timezone != "/timezone" {
		t.Fatalf("Expected /timezone, got %s", Timezone)
	}

	if Help != "/help" {
		t.Fatalf("Expected /help, got %s", Help)
	}

	if None != "" {
		t.Fatalf("Expected '', got %s", None)
	}
}

func TestParseCommand(t *testing.T) {
	command, parameters := ParseCommand("/start")
	if command != Start {
		t.Fatalf("Expected /start, got %s", command)
	}
	if parameters != nil {
		t.Fatalf("Expected nil, got %v", parameters)
	}
}

func TestCommandWithParameters(t *testing.T) {
	command, parameters := ParseCommand("/language en")
	if command != Language {
		t.Fatalf("Expected /language, got %s", command)
	}
	if len(parameters) != 1 {
		t.Fatalf("Expected 1, got %d", len(parameters))
	}
	if parameters[0] != "en" {
		t.Fatalf("Expected en, got %s", parameters[0])
	}

	command, parameters = ParseCommand("/timezone 3600")
	if command != Timezone {
		t.Fatalf("Expected /timezone, got %s", command)
	}
	if len(parameters) != 1 {
		t.Fatalf("Expected 1, got %d", len(parameters))
	}
	if parameters[0] != "3600" {
		t.Fatalf("Expected 3600, got %s", parameters[0])
	}
}
