package mocks

import (
	"context"

	"github.com/capymind/internal/database"
)

type UserStorageMock struct{}

func (storage UserStorageMock) GetUser(ctx *context.Context, userID string) (*database.User, error) {
	firstName := "John"
	lastName := "Doe"
	user := database.User{
		ID:        "1",
		FirstName: &firstName,
		LastName:  &lastName,
	}
	return &user, nil
}

func (storage UserStorageMock) SaveUser(ctx *context.Context, user database.User) error {
	return nil
}

func (storage UserStorageMock) DeleteUser(ctx *context.Context, userID string) error {
	return nil
}

func (storage UserStorageMock) ForEachUser(ctx *context.Context, callback func([]database.User) error) error {
	firstName := "John"
	lastName := "Doe"
	user1 := database.User{
		ID:        "1",
		FirstName: &firstName,
		LastName:  &lastName,
	}

	firstName = "Jane"
	lastName = "Doe"
	user2 := database.User{
		ID:        "2",
		FirstName: &firstName,
		LastName:  &lastName,
	}

	callback([]database.User{user1, user2})

	return nil
}
