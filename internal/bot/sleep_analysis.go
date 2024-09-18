package bot

import (
	"log"
	"strings"

	"github.com/capymind/internal/analysis"
	"github.com/capymind/internal/firestore"
	"github.com/capymind/internal/translator"
)

// Handle the sleep analysis command
func handleSleepAnalysis(session *Session) {
	setOutputText("analysis_waiting", session)

	userID := session.User.ID
	note, err := firestore.LastNote(session.Context, userID)
	if err != nil {
		log.Printf("[Bot] Error getting last note from firestore, %s", err.Error())
	}

	if note != nil {
		sleepAnalysis := analysis.AnalyzeSleep(note.Text, session.Locale(), session.Context)
		if sleepAnalysis != nil {
			setOutputText(*sleepAnalysis, session)
		} else {
			sendNoNotes(session)
		}
	} else {
		sendNoNotes(session)
	}
}

// Ask the user if they want a sleep analysis
func askForSleepAnalysis(session *Session) {
	var button JobResultTextButton = JobResultTextButton{
		TextID:   "sleep_analysis",
		Callback: string(SleepAnalysis),
	}
	setOutputTextWithButtons("do_you_want_sleep_analysis", []JobResultTextButton{button}, session)
}

func checkIfNoteADream(text string, locale translator.Locale) bool {
	keywords := translator.SearchKeywords(locale, "dreams")
	for _, keyword := range keywords {
		if strings.Contains(text, keyword) {
			return true
		}
	}
	return false
}
