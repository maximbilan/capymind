package analysis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/capymind/internal/translator"
	"github.com/openai/openai-go"
)

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

	response := request("analysis", "Analysis of the user's journal entries", systemPrompt, userPrompt, ctx)

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

	response := request("sleep_analysis", "Analysis of the user's sleep", systemPrompt, userPrompt, ctx)
	return response
}

// Request an AI analysis
func request(name string, description string, systemPrompt string, userPrompt string, ctx *context.Context) *string {
	ai := createAI()

	var responseSchema = generateSchema[Response]()

	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        openai.F(name),
		Description: openai.F(description),
		Schema:      openai.F(responseSchema),
		Strict:      openai.Bool(true),
	}

	chat, err := ai.Chat.Completions.New(*ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemPrompt),
			openai.UserMessage(userPrompt),
		}),
		ResponseFormat: openai.F[openai.ChatCompletionNewParamsResponseFormatUnion](
			openai.ResponseFormatJSONSchemaParam{
				Type:       openai.F(openai.ResponseFormatJSONSchemaTypeJSONSchema),
				JSONSchema: openai.F(schemaParam),
			},
		),
		Model: openai.F(openai.ChatModelGPT4oMini),
	})

	response := Response{}
	if err != nil {
		log.Printf("[AI] Error parsing analysis: %s", err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(chat.Choices[0].Message.Content), &response)
	if err != nil {
		log.Printf("[AI] Error parsing analysis: %s", err.Error())
		return nil
	}

	return &response.Text
}
