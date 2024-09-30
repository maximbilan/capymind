package bot

import "strings"

type Command string

const (
	// User commands
	Start    Command = "/start"
	Note     Command = "/note"
	Last     Command = "/last"
	Analysis Command = "/analysis"
	Language Command = "/language"
	Timezone Command = "/timezone"
	Help     Command = "/help"

	// Hidden user commands
	NoteCount      Command = "/note_count"      // Count of the current user notes
	SleepAnalysis  Command = "/sleep_analysis"  // Sleep analysis of last note
	WeeklyAnalysis Command = "/weekly_analysis" // Weekly analysis of the user's journal entries for last week

	// Admin commands
	TotalUserCount Command = "/total_user_count"
	TotalNoteCount Command = "/total_note_count"

	// Empty command
	None Command = "" // No command, just plain text
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
