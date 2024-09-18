package bot

import "strings"

type Command string

const (
	Start         Command = "/start"
	Note          Command = "/note"
	Last          Command = "/last"
	Analysis      Command = "/analysis"
	Language      Command = "/language"
	Timezone      Command = "/timezone"
	Help          Command = "/help"
	SleepAnalysis Command = "/sleep_analysis" // Sleep analysis of last note
	NotesCount    Command = "/notes_count"    // Count of notes
	None          Command = ""                // No command, just plain text
)

// Parse the command from the input
func ParseCommand(input string) (Command, []string) {
	if len(input) == 0 || input[0] != '/' {
		return None, nil
	}

	parts := strings.Split(input, " ")
	if len(parts) == 1 {
		return Command(parts[0]), nil
	} else {
		// Return the command and the rest of the input as parameters
		return Command(parts[0]), parts[1:]
	}
}
