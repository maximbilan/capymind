package localizer

import (
	"encoding/json"
	"log"
)

type Locale string

const (
	EN Locale = "en"
)

var translations map[Locale]map[string]string

const translationsJSON = `{
    "en": {
        "welcome": "Welcome to CapyMind! Your personal mental health journal is just a few taps away. Start making entries to reflect on your thoughts and emotions.",
        "info": "CapyMind is here to assist you in maintaining a personal journal for your mental health. You can record your thoughts and feelings, track your emotional journey, and reflect on your progress over time. Use the commands to start making entries and take a step towards self-awareness and mental well-being.",
        "start_note" : "Please input your thoughts and feelings in the text field and send them to me. Your personal reflections will be safely stored in your journal.",
        "finish_note" : "Your thoughts have been successfully stored. Thank you for sharing with CapyMind. Remember, every note is a step towards better mental well-being.",
        "your_last_note": "Here is your last note: ",
        "no_notes": "You have not made any entries yet. Start by sharing your thoughts and feelings with CapyMind.",
        "commands_hint": "You can use the following commands to interact with CapyMind:\n/start - Start the bot\n/note - Make a journal entry\n/last - View your last entry\n/info - Learn more about CapyMind\n/help - Get help with using CapyMind\n"
    }
}`

func init() {
	if err := json.Unmarshal([]byte(translationsJSON), &translations); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}
}

func Localize(locale Locale, key string) string {
	return translations[locale][key]
}
