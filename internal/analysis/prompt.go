package analysis

import "github.com/capymind/internal/translator"

type Prompt struct {
	System string
	User   string
}

type PromptType string

const (
	WeeklyAnalysis PromptType = "weekly_analysis" // Analysis of the user's journal entries for last week
	SleepAnalysis  PromptType = "sleep_analysis"  // Analysis of the user's last note for sleep patterns
	QuickAnalysis  PromptType = "quick_analysis"  // Analysis of the user's last 3 journal entries
)

func getPrompt(promptType PromptType, locale translator.Locale) Prompt {
	var system string
	var user string

	switch promptType {
	case WeeklyAnalysis:
		system = translator.Prompt(locale, "ai_weekly_analysis_system_message")
		user = translator.Prompt(locale, "ai_weekly_analysis_user_message")
	case SleepAnalysis:
		system = translator.Prompt(locale, "ai_sleep_analysis_system_message")
		user = translator.Prompt(locale, "ai_sleep_analysis_user_message")
	case QuickAnalysis:
		system = translator.Prompt(locale, "ai_quick_analysis_system_message")
		user = translator.Prompt(locale, "ai_quick_analysis_user_message")
	}

	return Prompt{
		System: system,
		User:   user,
	}
}
