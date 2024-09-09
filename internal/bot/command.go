package bot

import "strings"

type Command string

const (
	Start    Command = "/start"
	Note     Command = "/note"
	Last     Command = "/last"
	Analysis Command = "/analysis"
	Language Command = "/language"
	Timezone Command = "/timezone"
	Help     Command = "/help"
	None     Command = "" // No command, just plain text
)

var commandsWithParams = []Command{Language, Timezone}

func (c Command) HasParam() bool {
	for _, command := range commandsWithParams {
		if c == command {
			return true
		}
	}
	return false
}

func ParseCommand(input string) (Command, *string) {
	if len(input) == 0 || input[0] != '/' {
		return None, nil
	}

	parts := strings.Split(input, " ")
	if len(parts) == 1 {
		return Command(parts[0]), nil
	} else {
		return Command(parts[0]), &parts[1]
	}
}
