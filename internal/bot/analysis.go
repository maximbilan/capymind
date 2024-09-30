package bot

import (
	"github.com/capymind/internal/analysis"
)

// Handle the analysis command
func handleAnalysis(session *Session) {
	// Get the user's notes
	notes := getNotes(session, 5)
	if len(notes) > 0 {
		// Prepare the strings for analysis
		var strings []string
		for _, note := range notes {
			if note.Text != "" {
				strings = append(strings, note.Text)
			}
		}

		// Send a message to wait for the analysis
		sendMessage("analysis_waiting", session)

		// Request the analysis
		analysis := analysis.AnalyzeQuickly(strings, session.Locale(), session.Context)
		if analysis != nil {
			// Send the analysis
			setOutputText(*analysis, session)
			return
		}
	}

	// If there are no notes, send a message to make a note
	var button JobResultTextButton = JobResultTextButton{
		TextID:   "make_record_to_journal",
		Callback: string(Note),
	}
	setOutputTextWithButtons("no_analysis", []JobResultTextButton{button}, session)
}
