package ai

import (
	"context"
	"encoding/json"
	"log"
	"os"

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

func GetAnalysis(notes []string) *string {
	ctx := context.Background()
	client := createClient(ctx)

	prompt := "You're a professional therapist. A patient comes to you with the following notes. What would you say to them?"

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
	err = json.Unmarshal([]byte(chat.Choices[0].Message.Content), &analysis)
	if err == nil {
		log.Printf("[AI] Error parsing analysis: %s", err.Error())
		return nil
	}
	return &analysis.Text
}
