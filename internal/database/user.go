//coverage:ignore file

package database

import (
	"context"
	"time"
)

type User struct {
	ID                  string     `json:"id"`
	ChatID              int64      `json:"chatId"`
	UserName            *string    `json:"username"`
	FirstName           *string    `json:"firstName"`
	LastName            *string    `json:"lastName"`
	Locale              *string    `json:"locale"`
	SecondsFromUTC      *int       `json:"secondsFromUTC"`
	LastCommand         *string    `json:"lastCommand"`
	IsTyping            bool       `json:"isTyping"`
	IsOnboarded         bool       `json:"isOnboarded"`
	Role                *Role      `json:"role"`
	Timestamp           *time.Time `json:"timestamp"`
	IsDeleted           bool       `json:"isDeleted"`
	TherapySessionEndAt *time.Time `json:"therapySessionEndAt"`
}

type UserStorage interface {
	GetUser(ctx *context.Context, userID string) (*User, error)
	SaveUser(ctx *context.Context, user User) error
	DeleteUser(ctx *context.Context, userID string) error
	ForEachUser(ctx *context.Context, callback func([]User) error) error
}

// returns true if the user is active
// `active` means `timestamp` > 7 days ex. if the user has been active in the last 7 days
func (u User) IsActive() bool {
	if u.Timestamp == nil {
		return false
	}
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	return u.Timestamp.After(sevenDaysAgo)
}

// returns true if the user is inactive
// `inactive` means `timestamp` < 14 days ex. if the user has not been active in the last 14 days
func (u User) IsNonActive() bool {
	if u.Timestamp == nil {
		return true
	}
	fourteenDaysAgo := time.Now().AddDate(0, 0, -14)
	return u.Timestamp.Before(fourteenDaysAgo)
}
