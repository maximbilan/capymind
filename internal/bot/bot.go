package bot

import (
	"fmt"
	"log"
	"net/http"

	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
	"github.com/capymind/internal/utils"
)

func Parse(w http.ResponseWriter, r *http.Request) {
	update := telegram.Parse(r)
	if update == nil {
		return
	}

	callbackQuery := update.CallbackQuery
	if callbackQuery != nil && callbackQuery.Data != "" {
		userId := fmt.Sprintf("%d", callbackQuery.From.ID)
		userLocale := getUserLocaleByUserId(userId)
		locale := translator.EN
		if userLocale != nil {
			locale = translator.Locale(*userLocale)
		}

		if callbackQuery.Data == "note_make" {
			localizeAndSendMessage(callbackQuery.Message.Chat.Id, userId, locale, "start_note")
			startWritingMode(userId)
			return
		} else if callbackQuery.Data == "help" {
			sendHelpMessage(callbackQuery.Message.Chat.Id, userId, locale)
			return
		} else if callbackQuery.Data == "locale_setup" {
			sendLanguageSetMessage(callbackQuery.Message.Chat.Id, userId, locale)
		} else if callbackQuery.Data == "timezone_setup" {
			sendTimezoneSetMessage(callbackQuery.Message.Chat.Id, userId, locale)
		}

		log.Printf("[Bot] Received callback data: %s", callbackQuery.Data)
		updatedLocale, ok := translator.ParseLocale(callbackQuery.Data)
		if ok && updatedLocale != nil {
			userId := fmt.Sprintf("%d", callbackQuery.From.ID)
			setupLocale(userId, *updatedLocale)
			localizeAndSendMessage(callbackQuery.Message.Chat.Id, userId, translator.Locale(*updatedLocale), "locale_set")

			// If the user is setting the locale for the first time, also set the timezone
			if getTimeZone(userId) == nil {
				sendTimezoneSetMessage(callbackQuery.Message.Chat.Id, userId, locale)
			}
			return
		}

		secondsFromUTC, ok := utils.ParseTimezone(callbackQuery.Data)
		if ok && secondsFromUTC != nil {
			setupTimezone(userId, *secondsFromUTC)
			localizeAndSendMessage(callbackQuery.Message.Chat.Id, userId, locale, "timezone_set")
			if !userExists(userId) {
				sendStartMessage(callbackQuery.Message.Chat.Id, userId, callbackQuery.From.Username, locale)
			}
			return
		}

		return
	}

	message := update.Message

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
