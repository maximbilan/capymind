package app

import "strings"

type Command string

const (
	// User commands
	Start    Command = "/start"    // Start the bot
	Why      Command = "/why"      // Tells the purpose of the bot
	Note     Command = "/note"     // Add a new note
	Last     Command = "/last"     // Get the last note
	Analysis Command = "/analysis" // Get the analysis of the last notes
	Language Command = "/language" // Change the language
	Settings Command = "/settings" // Show the settings
	Help     Command = "/help"     // Show the help
	Version  Command = "/version"  // Show the version

	// Hidden user commands
	MissingNote        Command = "/missing_note"         // Ask to put a note from the previous text
	DownloadData       Command = "/download_data"        // Download the user data (all notes)
	DeleteAccount      Command = "/delete_account"       // Ask to delete the account
	ForceDeleteAccount Command = "/force_delete_account" // Force delete the account
	Support            Command = "/support"              // Give feedback or ask for support
	NoteCount          Command = "/note_count"           // Count of the current user notes
	SleepAnalysis      Command = "/sleep_analysis"       // Sleep analysis of last note
	WeeklyAnalysis     Command = "/weekly_analysis"      // Weekly analysis of the user's journal entries for last week

	// Reminder commands
	Reminders              Command = "/reminders"                 // Set reminders
	EnableAllReminders     Command = "/enable_all_reminders"      // Enable all reminders
	DisableAllReminders    Command = "/disable_all_reminders"     // Disable all reminders
	MorningReminder        Command = "/morning_reminder"          // Set morning reminder
	EnableMorningReminder  Command = "/enable_morning_reminder"   // Enable morning reminder
	DisableMorningReminder Command = "/disable_morning_reminder"  // Disable morning
	SetMorningReminderTime Command = "/set_morning_reminder_time" // Set morning reminder time
	EveningReminder        Command = "/evening_reminder"          // Set evening reminder
	EnableEveningReminder  Command = "/enable_evening_reminder"   // Enable evening reminder
	DisableEveningReminder Command = "/disable_evening_reminder"  // Disable evening reminder
	SetEveningReminderTime Command = "/set_evening_reminder_time" // Set evening reminder time
	SkipReminders          Command = "/skip_reminders"            // Skip the reminders (during onboarding)

	// Timezone commands
	Timezone   Command = "/timezone"     // Change the timezone (manually)
	AskForCity Command = "/ask_for_city" // Ask for the city to set the timezone

	// Admin commands
	TotalUserCount       Command = "/total_user_count"        // Get the total number of users
	TotalActiveUserCount Command = "/total_active_user_count" // Get the number of active users
	TotalNoteCount       Command = "/total_note_count"        // Get the total number of notes
	FeedbackLastWeek     Command = "/feedback_last_week"      // Get the feedback from the last week
	Stats                Command = "/stats"                   // Get the stats (total number of users, active users, notes, feedback)

	// Empty command
	None Command = "" // No command, just plain text
)

var adminCommands = []Command{
	TotalUserCount,
	TotalActiveUserCount,
	TotalNoteCount,
	FeedbackLastWeek,
	Stats,
}

// Check if the command is an admin command
func (c Command) IsAdmin() bool {
	for _, cmd := range adminCommands {
		if c == cmd {
			return true
		}
	}
	return false
}

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
