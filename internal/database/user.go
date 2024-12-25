package database

import (
	"context"
	"time"
)

type User struct {
	ID             string     `json:"id"`
	ChatID         int64      `json:"chatId"`
	UserName       *string    `json:"username"`
	FirstName      *string    `json:"firstName"`
	LastName       *string    `json:"lastName"`
	Locale         *string    `json:"locale"`
	SecondsFromUTC *int       `json:"secondsFromUTC"`
	LastCommand    *string    `json:"lastCommand"`
	IsTyping       bool       `json:"isTyping"`
	IsOnboarded    bool       `json:"isOnboarded"`
	Role           *Role      `json:"role"`
	Timestamp      *time.Time `json:"timestamp"`
}

type UserStorage interface {
	GetUser(ctx *context.Context, userID string) (*User, error)
	SaveUser(ctx *context.Context, user User) error
	ForEachUser(ctx *context.Context, callback func([]User) error) error
}
