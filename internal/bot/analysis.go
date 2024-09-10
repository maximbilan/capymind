package bot

import (
	"github.com/capymind/internal/analysis"
)

func handleAnalysis(session Session) {
	// Get the user's notes
	notes := getNotes(session)
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
		analysis := analysis.Request(strings, session.Locale())
		if analysis != nil {
			// Send the analysis
			setOutputText(*analysis, session)
			return
		}
	}

	// If there are no notes, send a message to make a note
	var button JobResultTextButton = JobResultTextButton{
		TextID:   "make_record_to_journal",
		Callback: "note_make",
	}
	setOutputTextWithButtons("no_analysis", []JobResultTextButton{button}, session)
}