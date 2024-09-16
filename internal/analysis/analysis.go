package analysis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/capymind/internal/translator"
	"github.com/openai/openai-go"
)

// Reqeust an analysis of the user's journal entries
func Request(notes []string, locale translator.Locale, ctx *context.Context, header *string) *string {
	ai := createAI()

	systemPrompt := translator.Prompt(locale, "ai_analysis_system_message")
	userPrompt := translator.Prompt(locale, "ai_analysis_user_message")
	for index, note := range notes {
		userPrompt += fmt.Sprintf("%d. %s ", index+1, note)
	}

	var responseSchema = generateSchema[Response]()

	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        openai.F("therapy-analysis"),
		Description: openai.F("Analysis of the user's journal entries"),
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

	var analysis string
	if response.Text != "" {
		if header != nil {
			analysis = fmt.Sprintf("%s%s", translator.Translate(locale, *header), response.Text)
		} else {
			analysis = response.Text
		}
		return &analysis
	} else {
		return nil
	}
}
