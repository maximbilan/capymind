package telegram

import "fmt"

type User struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot,omitempty"`
	UserName     string `json:"username,omitempty"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

func (user *User) StringID() string {
	return fmt.Sprintf("%d", user.ID)
}
