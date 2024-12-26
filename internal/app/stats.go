package app

import (
	"github.com/capymind/internal/helpers"
)

func handleStats(session *Session) {
	stats := helpers.GetStats(session.Context, session.Locale())

	var finalString string
	for _, stat := range stats {
		finalString += stat + "\n"
	}
	setOutputText(finalString, session)
}
