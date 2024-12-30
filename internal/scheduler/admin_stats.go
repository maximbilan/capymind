package scheduler

import (
	"context"

	"github.com/capymind/internal/helpers"
	"github.com/capymind/internal/translator"
)

func prepareAdminStats(ctx *context.Context, locale translator.Locale) *string {
	stats := helpers.GetStats(ctx, locale, adminStorage, feedbackStorage)

	var finalString string
	for _, stat := range stats {
		finalString += stat + "\n"
	}
	return &finalString
}
