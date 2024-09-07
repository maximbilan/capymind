package analysis

import (
	"context"
	"os"

	"github.com/invopop/jsonschema"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

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
