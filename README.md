# capymind

CapyMind is a personal mental health journal designed to help you track your thoughts, emotions, and progress over time. By offering a simple and secure platform to make journal entries, set reminders, and receive personalized therapy insights, CapyMind empowers you to reflect on your mental well-being. With support for multiple languages and time zones, it fits seamlessly into your daily routine, providing a personalized space for self-reflection and growth.

# Screenshots

<table>
  <tr>
    <td><img src="assets/screen1.jpeg" alt="Image 1" width="300"/></td>
    <td><img src="assets/screen2.jpeg" alt="Image 2" width="300"/></td>
    <td><img src="assets/screen3.jpeg" alt="Image 3" width="300"/></td>
  </tr>
</table>

# Reference

<a href="t.me/capymind_bot">capymind</a> is available as a Telegram bot.

# Commands

## User commands
```
/start - Start the bot
/why - Tells the purpose of the bot
/note - Add a new note
/last - Get the last note
/analysis - Get the analysis of the last notes
/language - Change the language
/settings - Show the settings
/help - Show the help
/version - Show the version
```

## Hidden user commands
```
/missing_note - Ask to put a note from the previous text
/timezone - Change the timezone
/download_data - Download the user data (all notes)
/delete_account - Ask to delete the account
/force_delete_account - Force delete the account
/support - Give feedback or ask for support
/note_count - Count of the current user notes
/sleep_analysis - Sleep analysis of last note
/weekly_analysis - Weekly analysis of the user's journal entries for last week
```

## Admin commands
```
/total_user_count - Get the total number of users
/total_active_user_count - Get the number of active users
/total_note_count - Get the total number of notes
/feedback_last_week - Get the feedback from the last week
/stats - Get the stats (total number of users, active users, notes, feedback)
```

# Building

## Prerequisites

1. Install `Go`
2. Tunnel ex. `ngrok`

## How to run locally

1. `go build`
2. Run `ngrok` or other tunnel
3. Set up `CAPY_CLOUD_FUNCTION_URL` as a ENV variable with the tunnel url
4. `chmod +x ./scripts/setup_telegram.bot.sh`
5. `./scripts/setup_telegram.bot.sh` (Set up a Telegram token)
6. `go run cmd/main.go` to run the server

# Scripts

More information <a href="./docs/SCRIPTS.md">here</a>
