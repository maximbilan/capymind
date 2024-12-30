//coverage:ignore file

package mocks

import "context"

type ValidAIServiceMock struct{}

func (service ValidAIServiceMock) Request(name string, description string, systemPrompt string, userPrompt string, ctx *context.Context) *string {
	response := "valid response"
	return &response
}

type InvalidAIServiceMock struct{}

func (service InvalidAIServiceMock) Request(name string, description string, systemPrompt string, userPrompt string, ctx *context.Context) *string {
	return nil
}
