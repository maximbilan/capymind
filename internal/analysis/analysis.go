package analysis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/capymind/internal/translator"
	"github.com/openai/openai-go"
)

func Request(notes []string, locale translator.Locale) *string {
	ctx := context.Background()
	client := createClient(ctx)

	prompt := translator.Translate(locale, "ai_analysis_prompt")
	for index, note := range notes {
		prompt += fmt.Sprintf("%d. %s ", index+1, note)
	}

	var responseSchema = generateSchema[Response]()

	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        openai.F("therapy-analysis"),
		Description: openai.F("Analysis of the user's journal entries"),
		Schema:      openai.F(responseSchema),
		Strict:      openai.Bool(true),
	}

	chat, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		}),
		ResponseFormat: openai.F[openai.ChatCompletionNewParamsResponseFormatUnion](
			openai.ResponseFormatJSONSchemaParam{
				Type:       openai.F(openai.ResponseFormatJSONSchemaTypeJSONSchema),
				JSONSchema: openai.F(schemaParam),
			},
		),
		// only certain models can perform structured outputs
		Model: openai.F(openai.ChatModelGPT4o2024_08_06),
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
