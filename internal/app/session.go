package app

import (
	"context"
	"time"

	"github.com/capymind/internal/database"
	"github.com/capymind/internal/translator"
)

type Session struct {
	Job      *Job
	User     *database.User
	Settings *database.Settings
	Context  *context.Context
}

// Return the locale of the current user
func (session *Session) Locale() translator.Locale {
	if session.User.Locale != nil {
		return translator.Locale(*session.User.Locale)
	}
	return translator.EN
}

// Save the user's data
func (session *Session) SaveUser(userStorage database.UserStorage) {
	if session.User.IsDeleted {
		// Do not save the user if it is deleted
		return
	}
	saveUser(session.User, session.Context, userStorage)
}

// Save the user's settings
func (session *Session) SaveSettings(settings database.Settings, settingsStorage database.SettingsStorage) {
	//coverage:ignore
	saveSettings(session.Context, session.User.ID, settings, settingsStorage)
}

// Create a session
func createSession(job *Job, user *database.User, settings *database.Settings, context *context.Context) *Session {
	session := Session{
		Job:      job,
		User:     user,
		Settings: settings,
		Context:  context,
	}
	return &session
}

func (session *Session) isAdmin() bool {
	return isAdmin(session.User)
}

// Handle the session
func handleSession(session *Session) {
	now := time.Now()
	command := session.Job.Command
	commandStr := string(command)
	session.User.LastCommand = &commandStr
	session.User.Timestamp = &now

	if command.IsAdmin() && !session.isAdmin() {
		handleHelp(session)
		return
	}

	switch command {
	case Start:
		handleStart(session)
	case Why:
		handleWhy(session)
	case Note:
		startNote(session)
	case MissingNote:
		handleMissingNote(session, noteStorage)
	case Last:
		handleLastNote(session, noteStorage)
	case Analysis:
		handleAnalysis(session, noteStorage, aiService)
	case Settings:
		handleSettings(session)
	case Language:
		handleLanguage(session)
	case Timezone:
		handleTimezone(session, settingsStorage)
	case AskForCity:
		handleCityRequest(session)
	case Reminders:
		handleReminders(session)
	case MorningReminder:
		handleMorningReminder(session)
	case EveningReminder:
		handleEveningReminder(session)
	case EnableAllReminders:
		enableAllReminders(session, settingsStorage)
	case DisableAllReminders:
		disableAllReminders(session, settingsStorage)
	case EnableMorningReminder:
		enableMorningReminder(session, settingsStorage)
	case DisableMorningReminder:
		disableMorningReminder(session, settingsStorage)
	case EnableEveningReminder:
		enableEveningReminder(session, settingsStorage)
	case DisableEveningReminder:
		disableEveningReminder(session, settingsStorage)
	case SetMorningReminderTime:
		setMorningReminderOffset(session, settingsStorage)
	case SetEveningReminderTime:
		setEveningReminderOffset(session, settingsStorage)
	case SkipReminders:
		skipReminders(session, settingsStorage)
	case Support:
		startFeedback(session)
	case Help:
		handleHelp(session)
	case Version:
		handleVersion(session)
	case SleepAnalysis:
		handleSleepAnalysis(session, noteStorage, aiService)
	case WeeklyAnalysis:
		handleWeeklyAnalysis(session, noteStorage, aiService)
	case NoteCount:
		handleNoteCount(session, noteStorage)
	case DownloadData:
		handleDownloadData(session, noteStorage, fileStorage)
	case DeleteAccount:
		handleDeleteAccount(session)
	case ForceDeleteAccount:
		handleForceDeleteAccount(session, noteStorage, userStorage, settingsStorage)
	case TotalUserCount:
		handleTotalUserCount(session, adminStorage)
	case TotalActiveUserCount:
		handleTotalActiveUserCount(session, adminStorage)
	case TotalNoteCount:
		handleTotalNoteCount(session, adminStorage)
	case Stats:
		handleStats(session, adminStorage, feedbackStorage)
	case FeedbackLastWeek:
		handleFeedbackLastWeek(session, feedbackStorage)
	case None:
		// Typing mode
		if session.User.IsTyping && session.Job.Input != nil {
			switch session.Job.LastCommand {
			case Note:
				finishNote(*session.Job.Input, session, noteStorage)
			case Support:
				finishFeedback(session, feedbackStorage)
			case AskForCity:
				finishCityRequest(session, mapsService, settingsStorage)
			default:
				// If the user is typing and the last command is not recognized
				handleHelp(session)
			}
		} else {
			if session.Job.Input != nil && len(*session.Job.Input) > 1 && (*session.Job.Input)[0] != '/' {
				// If the user is not typing and the input is not a command
				handleInputWithoutCommand(session)
			} else {
				// If the user is not typing and the input is not recognized
				handleHelp(session)
			}
		}
	default:
		// Unknown command
		handleHelp(session)
	}
}

// Finish the session. Send the output to the user
func finishSession(session *Session) {
	//coverage:ignore
	// Save the user's data
	session.SaveUser(userStorage)
	// Prepare the messages, localize and send it
	sendOutputMessages(session)
}
