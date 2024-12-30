package scheduler

import (
	"context"
	"log"

	"github.com/capymind/internal/aiservice"
	"github.com/capymind/internal/analysis"
	"github.com/capymind/internal/database"
	"github.com/capymind/internal/translator"
)

func prepareWeeklyAnalysis(user *database.User, ctx *context.Context, locale translator.Locale, aiService aiservice.AIService) *string {
	notes, err := noteStorage.GetNotesForLastWeek(ctx, user.ID)
	if err != nil {
		log.Printf("[Scheduler] Error getting notes from firestore, %s", err.Error())
		return nil
	}

	if len(notes) > 0 {
		var strings []string
		for _, note := range notes {
			if note.Text != "" {
				strings = append(strings, note.Text)
			}
		}
		return analysis.AnalyzeLastWeek(aiService, strings, locale, ctx)
	} else {
		return nil
	}
}
