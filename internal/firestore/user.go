package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
)

type User struct {
	ID             string  `firestore:"id"`
	ChatID         int64   `firestore:"chatID"`
	Name           *string `firestore:"name"` // Deprecated
	UserName       *string `json:"username"`
	FirstName      *string `json:"firstName"`
	LastName       *string `json:"lastName"`
	Locale         *string `firestore:"locale"`
	LastChatID     *int64  `firestore:"lastChatId"` // Deprecated
	SecondsFromUTC *int    `firestore:"secondsFromUTC"`
	IsWriting      bool    `firestore:"isWriting"` // Deprecated
	LastCommand    string  `firestore:"lastCommand"`
	IsTyping       bool    `firestore:"isTyping"`
}

func NewUser(ctx context.Context, client *firestore.Client, user User) error {
	_, err := client.Collection(users.String()).Doc(user.ID).Set(ctx, map[string]interface{}{
		"id":        user.ID,
		"username":  user.UserName,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
	}, firestore.MergeAll)
	return err
}

func UserExists(ctx context.Context, client *firestore.Client, userID string) (bool, error) {
	user, _ := getUser(ctx, client, userID)
	if user != nil {
		return user.Name != nil, nil
	}
	return false, nil
}

func getUser(ctx context.Context, client *firestore.Client, userID string) (*User, error) {
	doc, err := client.Collection(users.String()).Doc(userID).Get(ctx)
	if err != nil {
		return nil, err
	}

	var user User
	doc.DataTo(&user)
	return &user, nil
}

func UserLocale(ctx context.Context, client *firestore.Client, userID string) (*string, error) {
	user, err := getUser(ctx, client, userID)
	if err != nil {
		return nil, err
	}
	return user.Locale, nil
}

func UpdateUserLocale(ctx context.Context, client *firestore.Client, userID string, locale string) error {
	_, err := client.Collection(users.String()).Doc(userID).Set(ctx, map[string]interface{}{
		"locale": locale,
	}, firestore.MergeAll)
	return err
}

func UpdateUserTimezone(ctx context.Context, client *firestore.Client, userID string, secondsFromUTC int) error {
	_, err := client.Collection(users.String()).Doc(userID).Set(ctx, map[string]interface{}{
		"secondsFromUTC": secondsFromUTC,
	}, firestore.MergeAll)
	return err
}

func SaveLastChatID(ctx context.Context, client *firestore.Client, userID string, chatID int64) error {
	_, err := client.Collection(users.String()).Doc(userID).Set(ctx, map[string]interface{}{
		"id":     userID,
		"chatId": userID,
	}, firestore.MergeAll)
	return err
}

func ForEachUser(ctx context.Context, client *firestore.Client, callback func([]User) error) error {
	var lastDoc *firestore.DocumentSnapshot
	for {
		query := client.Collection(users.String()).OrderBy(firestore.DocumentID, firestore.Asc).Limit(100)
		if lastDoc != nil {
			query = query.StartAfter(lastDoc)
		}

		docs, err := query.Documents(ctx).GetAll()
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

func UserTimezone(ctx context.Context, client *firestore.Client, userID string) (*int, error) {
	user, err := getUser(ctx, client, userID)
	if err != nil {
		return nil, err
	}
	return user.SecondsFromUTC, nil
}

func updateWritingMode(ctx context.Context, client *firestore.Client, userID string, state bool) error {
	_, err := client.Collection(users.String()).Doc(userID).Set(ctx, map[string]interface{}{
		"isWriting": state,
	}, firestore.MergeAll)
	return err
}

func StartWriting(ctx context.Context, client *firestore.Client, userID string) error {
	return updateWritingMode(ctx, client, userID, true)
}

func StopWriting(ctx context.Context, client *firestore.Client, userID string) error {
	return updateWritingMode(ctx, client, userID, false)
}

func UserWritingStatus(ctx context.Context, client *firestore.Client, userID string) (bool, error) {
	user, err := getUser(ctx, client, userID)
	if err != nil {
		return false, err
	}
	return user.IsWriting, nil
}
