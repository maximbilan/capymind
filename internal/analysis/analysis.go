package analysis

import (
	"context"
	"fmt"

	"github.com/capymind/internal/translator"
	"github.com/capymind/third_party/openai"
)

var service openai.OpenAI

func AnalyzeQuickly(notes []string, locale translator.Locale, ctx *context.Context) *string {
	return analyzeJournal(getLocalizedPrompt(QuickAnalysis, locale), notes, locale, ctx, nil)
}

func AnalyzeLastWeek(notes []string, locale translator.Locale, ctx *context.Context) *string {
	header := "weekly_analysis"
	return analyzeJournal(getLocalizedPrompt(WeeklyAnalysis, locale), notes, locale, ctx, &header)
}

func analyzeJournal(prompt Prompt, notes []string, locale translator.Locale, ctx *context.Context, header *string) *string {
	systemPrompt := prompt.System
	userPrompt := prompt.User
	for index, note := range notes {
		userPrompt += fmt.Sprintf("%d. %s ", index+1, note)
	}

	response := service.Request("analysis", "Analysis of the user's journal entries", systemPrompt, userPrompt, ctx)

	var analysis string
	if response != nil {
		if header != nil {
			analysis = fmt.Sprintf("%s%s", translator.Translate(locale, *header), *response)
		} else {
			analysis = *response
		}
		return &analysis
	} else {
		return nil
	}
}

// Request an analysis of the user's sleep
func AnalyzeSleep(text string, locale translator.Locale, ctx *context.Context) *string {
	prompt := getLocalizedPrompt(SleepAnalysis, locale)

	systemPrompt := prompt.System
	userPrompt := prompt.User
	userPrompt += text

	response := service.Request("sleep_analysis", "Analysis of the user's sleep", systemPrompt, userPrompt, ctx)
	return response
}
