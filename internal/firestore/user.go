package firestore

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

type User struct {
	ID             string     `firestore:"id"`
	ChatID         int64      `firestore:"chatId"`
	UserName       *string    `json:"username"`
	FirstName      *string    `json:"firstName"`
	LastName       *string    `json:"lastName"`
	Locale         *string    `firestore:"locale"`
	SecondsFromUTC *int       `firestore:"secondsFromUTC"`
	LastCommand    *string    `firestore:"lastCommand"`
	IsTyping       bool       `firestore:"isTyping"`
	IsOnboarded    bool       `firestore:"isOnboarded"`
	Role           *Role      `firestore:"role"`
	Timestamp      *time.Time `firestore:"timestamp"`
}

// Get a user from the database
func GetUser(ctx *context.Context, userID string) (*User, error) {
	doc, err := client.Collection(users.String()).Doc(userID).Get(*ctx)
	if err != nil {
		return nil, err
	}

	var user User
	doc.DataTo(&user)
	return &user, nil
}

// Save a user to the database
func SaveUser(ctx *context.Context, user User) error {
	_, err := client.Collection(users.String()).Doc(user.ID).Set(*ctx, user)
	return err
}

// Iterate over all users
func ForEachUser(ctx *context.Context, callback func([]User) error) error {
	var lastDoc *firestore.DocumentSnapshot
	for {
		query := client.Collection(users.String()).OrderBy(firestore.DocumentID, firestore.Asc).Limit(100)
		if lastDoc != nil {
			query = query.StartAfter(lastDoc)
		}

		docs, err := query.Documents(*ctx).GetAll()
		if err != nil {
			return err
		}

		var users []User
		for _, doc := range docs {
			var user User
			doc.DataTo(&user)
			users = append(users, user)
		}

		err = callback(users)
		if err != nil {
			return err
		}

		if len(docs) < 100 {
			break
		}
		lastDoc = docs[len(docs)-1]
	}
	return nil
}
