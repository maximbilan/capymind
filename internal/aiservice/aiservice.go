//coverage:ignore file

package aiservice

import "context"

type AIService interface {
	Request(name string, description string, systemPrompt string, userPrompt string, ctx *context.Context) *string
}
