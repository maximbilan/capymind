package analysis

import (
	"os"

	"github.com/invopop/jsonschema"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// Create a client for the OpenAI API
func createAI() *openai.Client {
	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("CAPY_AI_KEY")),
	)
	return client
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
