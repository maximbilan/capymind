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

func handleUser(message telegram.Message) {
	if message.Text == "" {
		return
	}
	createOrUpdateUser(message)
}

func handleStart(message telegram.Message, locale translator.Locale) {
	userId := fmt.Sprintf("%d", message.From.ID)
	localizeAndSendMessage(message.Chat.Id, userId, locale, "welcome")
	handleUser(message)
}

func handleNote(message telegram.Message, locale translator.Locale) {
	userId := message.From.ID
	userIdStr := fmt.Sprintf("%d", userId)
	localizeAndSendMessage(message.Chat.Id, userIdStr, locale, "start_note")
	userIds.Append(userId)
}

func handleLast(message telegram.Message, locale translator.Locale) {
	userId := fmt.Sprintf("%d", message.From.ID)
	note := getLastNote(message)
	if note != nil {
		var response string = translator.Translate(locale, "your_last_note") + note.Text
		sendMessage(message.Chat.Id, userId, response)
	} else {
		localizeAndSendMessage(message.Chat.Id, userId, locale, "no_notes")
	}
}

func handleAnalysis(message telegram.Message, locale translator.Locale) {
	userId := fmt.Sprintf("%d", message.From.ID)
	notes := getNotes(message)
	if len(notes) > 0 {
		analysis := "Pretty Good"
		sendMessage(message.Chat.Id, userId, analysis)
	} else {
		localizeAndSendMessage(message.Chat.Id, userId, locale, "no_notes")
	}
}

func handleLocale(message telegram.Message, locale translator.Locale) {
	userId := fmt.Sprintf("%d", message.From.ID)
	replyMarkup := telegram.InlineKeyboardMarkup{
		InlineKeyboard: [][]telegram.InlineKeyboardButton{
			{
				{Text: translator.English.String(), CallbackData: translator.GetLocaleParameter(translator.EN)},
				{Text: translator.Ukrainian.String(), CallbackData: translator.GetLocaleParameter(translator.UK)},
			},
		},
	}
	localizeAndSendMessageWithReply(message.Chat.Id, userId, locale, "language_select", &replyMarkup)
}

func handleTimezone(message telegram.Message, locale translator.Locale) {
	userId := fmt.Sprintf("%d", message.From.ID)
	timeZones := utils.GetTimeZones()

	var inlineKeyboard [][]telegram.InlineKeyboardButton
	for _, tz := range timeZones {
		inlineKeyboard = append(inlineKeyboard, []telegram.InlineKeyboardButton{
			{Text: tz.String(), CallbackData: utils.GetTimezoneParameter(tz)},
		})
	}

	replyMarkup := telegram.InlineKeyboardMarkup{
		InlineKeyboard: inlineKeyboard,
	}

	localizeAndSendMessageWithReply(message.Chat.Id, userId, locale, "timezone_select", &replyMarkup)
}

func handleUnknownState(message telegram.Message, locale translator.Locale) {
	userId := message.From.ID
	userIdStr := fmt.Sprintf("%d", userId)
	if userIds.Contains(userId) {
		saveNote(message)
		localizeAndSendMessage(message.Chat.Id, userIdStr, locale, "finish_note")
		userIds.Remove(userId)
	} else {
		handleHelp(message, locale)
	}
}

func handleInfo(message telegram.Message, locale translator.Locale) {
	userId := fmt.Sprintf("%d", message.From.ID)
	localizeAndSendMessage(message.Chat.Id, userId, locale, "info")
}

func handleHelp(message telegram.Message, locale translator.Locale) {
	userId := fmt.Sprintf("%d", message.From.ID)
	localizeAndSendMessage(message.Chat.Id, userId, locale, "commands_hint")
}
