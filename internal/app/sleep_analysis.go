package app

import (
	"log"
	"strings"

	"github.com/capymind/internal/aiservice"
	"github.com/capymind/internal/analysis"
	"github.com/capymind/internal/botservice"
	"github.com/capymind/internal/database"
	"github.com/capymind/internal/translator"
)

// Handle the sleep analysis command
func handleSleepAnalysis(session *Session, noteStorage database.NoteStorage, aiService aiservice.AIService) {
	sendMessage("analysis_waiting", session)

	userID := session.User.ID
	note, err := noteStorage.LastNote(session.Context, userID)
	if err != nil {
		log.Printf("[Bot] Error getting last note from firestore, %s", err.Error())
	}

	if note != nil {
		sleepAnalysis := analysis.AnalyzeSleep(aiService, note.Text, session.Locale(), session.Context)
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
	var button botservice.BotResultTextButton = botservice.BotResultTextButton{
		TextID:   "sleep_analysis",
		Locale:   session.Locale(),
		Callback: string(SleepAnalysis),
	}
	setOutputTextWithButtons("do_you_want_sleep_analysis", []botservice.BotResultTextButton{button}, session)
}

func checkIfNoteADream(text string, locale translator.Locale) bool {
	keywords := translator.SearchKeywords(locale, "dreams")
	for _, keyword := range keywords {
		if strings.Contains(strings.ToLower(text), keyword) {
			return true
		}
	}
	return false
}

func handleWeeklyAnalysis(session *Session, noteStorage database.NoteStorage, aiService aiservice.AIService) {
	sendMessage("analysis_waiting", session)

	userID := session.User.ID
	notes, err := noteStorage.GetNotesForLastWeek(session.Context, userID)
	if err != nil {
		log.Printf("[Bot] Error getting notes from firestore, %s", err.Error())
	}

	if len(notes) > 0 {
		var strings []string
		for _, note := range notes {
			if note.Text != "" {
				strings = append(strings, note.Text)
			}
		}

		analysis := analysis.AnalyzeLastWeek(aiService, strings, session.Locale(), session.Context)
		if analysis != nil {
			setOutputText(*analysis, session)
		} else {
			sendNoNotes(session)
		}
	} else {
		sendNoNotes(session)
	}
}
