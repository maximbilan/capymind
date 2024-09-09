package bot

import (
	"log"
	"net/http"

	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
	"github.com/capymind/internal/utils"
)

func Parse(w http.ResponseWriter, r *http.Request) {
	update := telegram.Parse(r)
	if update == nil {
		// No update to process
		return
	}

	// Create a job for the user

	// Fetch the user's data from the database

	// Update the user's data in the database if necessary

	// Create a session

	// Process the job

	callbackQuery := update.CallbackQuery
	if callbackQuery != nil && callbackQuery.Data != "" {
		userID := callbackQuery.UserID()
		chatID := callbackQuery.ChatID()
		if userID == nil || chatID == nil {
			return
		}

		userLocale := getUserLocaleByUserID(*userID)
		locale := translator.EN
		if userLocale != nil {
			locale = translator.Locale(*userLocale)
		}

		if callbackQuery.Data == "note_make" {
			localizeAndSendMessage(*chatID, *userID, locale, "start_note")
			startWritingMode(*userID)
			return
		} else if callbackQuery.Data == "help" {
			sendHelpMessage(*chatID, *userID, locale)
			return
		} else if callbackQuery.Data == "locale_setup" {
			sendLanguageSetMessage(*chatID, *userID, locale)
		} else if callbackQuery.Data == "timezone_setup" {
			sendTimezoneSetMessage(*chatID, *userID, locale)
		}

		log.Printf("[Bot] Received callback data: %s", callbackQuery.Data)
		updatedLocale, ok := translator.ParseLocale(callbackQuery.Data)
		if ok && updatedLocale != nil {
			setupLocale(*userID, *updatedLocale)
			newLocale := translator.Locale(*updatedLocale)
			localizeAndSendMessage(*chatID, *userID, newLocale, "locale_set")

			// If the user is setting the locale for the first time, also set the timezone
			if getTimeZone(*userID) == nil {
				sendTimezoneSetMessage(*chatID, *userID, newLocale)
			}
			return
		}

		secondsFromUTC, ok := utils.ParseTimezone(callbackQuery.Data)
		if ok && secondsFromUTC != nil {
			setupTimezone(*userID, *secondsFromUTC)
			localizeAndSendMessage(*chatID, *userID, locale, "timezone_set")
			if !userExists(*userID) {
				sendStartMessage(*chatID, *userID, &callbackQuery.From.UserName, locale)
			}
			return
		}

		return
	}

	if update.Message == nil {
		return
	}
	message := *update.Message

	var locale translator.Locale
	userLocale := getUserLocale(message)
	if userLocale != nil {
		locale = *userLocale
	} else {
		locale = translator.EN
	}

	text := message.Text
	command := Command(text)

	log.Printf("[Bot] Received message text: %s", text)

	switch command {
	case Start:
		handleStart(message, locale)
	case Note:
		handleNote(message, locale)
	case Last:
		handleLast(message, locale)
	case Analysis:
		handleAnalysis(message, locale)
	case Language:
		handleLanguage(message, locale)
	case Timezone:
		handleTimezone(message, locale)
	case Help:
		handleHelp(message, locale)
	default:
		handleUnknownState(message, locale)
	}
}
