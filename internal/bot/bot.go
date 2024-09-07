package bot

import (
	"fmt"
	"log"
	"net/http"

	"github.com/capymind/internal/telegram"
	"github.com/capymind/internal/translator"
	"github.com/capymind/internal/utils"
)

var userIds *utils.ThreadSafeArray[int64]

func init() {
	userIds = utils.NewThreadSafeArray[int64]()
}

func Parse(w http.ResponseWriter, r *http.Request) {
	update := telegram.Parse(r)
	if update == nil {
		return
	}

	callbackQuery := update.CallbackQuery
	if callbackQuery != nil && callbackQuery.Data != "" {
		log.Printf("[Bot] Received callback data: %s", callbackQuery.Data)
		updatedLocale, ok := translator.ParseLocale(callbackQuery.Data)
		if ok && updatedLocale != nil {
			userId := fmt.Sprintf("%d", callbackQuery.From.ID)
			setupLocale(userId, *updatedLocale)
			localizeAndSendMessage(callbackQuery.Message.Chat.Id, userId, translator.Locale(*updatedLocale), "locale_set")
			return
		}

		secondsFromUTC, ok := utils.ParseTimezone(callbackQuery.Data)
		if ok && secondsFromUTC != nil {
			userId := fmt.Sprintf("%d", callbackQuery.From.ID)
			setupTimezone(userId, *secondsFromUTC)

			userLocale := getUserLocaleByUserId(userId)
			locale := translator.EN
			if userLocale != nil {
				locale = translator.Locale(*userLocale)
			}
			localizeAndSendMessage(callbackQuery.Message.Chat.Id, userId, locale, "timezone_set")
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
	case Locale:
		handleLocale(message, locale)
	case Timezone:
		handleTimezone(message, locale)
	case Info:
		handleInfo(message, locale)
	case Help:
		handleHelp(message, locale)
	default:
		handleUnknownState(message, locale)
	}
}
