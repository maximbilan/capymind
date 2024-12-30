package scheduler

import (
	"context"
	"fmt"
	"log"

	"github.com/capymind/internal/database"
	"github.com/capymind/internal/translator"
)

func prepareUserStats(user *database.User, ctx *context.Context, locale translator.Locale) *string {
	count, err := noteStorage.NotesCount(ctx, user.ID)
	if err != nil {
		log.Printf("[Scheduler] Error getting notes count from firestore, %s", err.Error())
		return nil
	}
	// Send only if the user has more than one note in the journal
	if count > 1 {
		localizedMessage := fmt.Sprintf(translator.Translate(locale, "user_progress_message"), count)
		return &localizedMessage
	} else {
		return nil
	}
}
