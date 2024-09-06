package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/capymind/internal/translator"
	"github.com/invopop/jsonschema"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type Analysis struct {
	Text string `json:"text"`
}

func createClient(ctx context.Context) *openai.Client {
	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("CAPY_AI_KEY")),
	)
	return client
}

func generateSchema[T any]() interface{} {
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v T
	schema := reflector.Reflect(v)
	return schema
}

func GetAnalysis(notes []string, locale translator.Locale) *string {
	ctx := context.Background()
	client := createClient(ctx)

	prompt := translator.Translate(locale, "ai_analysis_prompt")
	for index, note := range notes {
		prompt += fmt.Sprintf("%d. %s ", index+1, note)
	}

	var responseSchema = generateSchema[Analysis]()

	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Schema: openai.F(responseSchema),
		Strict: openai.Bool(true),
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

	analysis := Analysis{}
	if err == nil {
		log.Printf("[AI] Error parsing analysis: %s", err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(chat.Choices[0].Message.Content), &analysis)
	if err == nil {
		log.Printf("[AI] Error parsing analysis: %s", err.Error())
		return nil
	}
	return &analysis.Text
}
