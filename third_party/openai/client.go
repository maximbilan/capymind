//coverage:ignore file

package openai

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/invopop/jsonschema"
	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

type Response struct {
	Text string `json:"text"`
}

type OpenAI struct{}

// Create a client for the OpenAI API
func createAI() *openai.Client {
	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("CAPY_AI_KEY")),
	)
	return &client
}

// Generate a JSON schema for a given type
func generateSchema[T any]() interface{} {
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v T
	schema := reflector.Reflect(v)
	return schema
}

// Request an AI analysis
func (service OpenAI) Request(name string, description string, systemPrompt string, userPrompt string, ctx *context.Context) *string {
	ai := createAI()

	var responseSchema = generateSchema[Response]()

	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        name,
		Description: openai.String(description),
		Schema:      responseSchema,
		Strict:      openai.Bool(true),
	}

	chat, err := ai.Chat.Completions.New(*ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemPrompt),
			openai.UserMessage(userPrompt),
		},
		ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
			OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{
				JSONSchema: schemaParam,
			},
		},
		Model: openai.ChatModelO3Mini,
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
