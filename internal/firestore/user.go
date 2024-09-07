package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
)

type User struct {
	ID             string  `firestore:"id"`
	Name           string  `firestore:"name"`
	Locale         *string `firestore:"locale"`
	LastChatId     *int    `firestore:"lastChatId"`
	SecondsFromUTC *int64  `firestore:"secondsFromUTC"`
}

func NewUser(ctx context.Context, client *firestore.Client, user User) error {
	_, err := client.Collection(users.String()).Doc(user.ID).Set(ctx, map[string]interface{}{
		"name": user.Name,
	}, firestore.MergeAll)
	return err
}

func getUser(ctx context.Context, client *firestore.Client, userId string) (*User, error) {
	doc, err := client.Collection(users.String()).Doc(userId).Get(ctx)
	if err != nil {
		return nil, err
	}

	var user User
	doc.DataTo(&user)
	return &user, nil
}

func UserLocale(ctx context.Context, client *firestore.Client, userId string) (*string, error) {
	user, err := getUser(ctx, client, userId)
	if err != nil {
		return nil, err
	}
	return user.Locale, nil
}

func UpdateUserLocale(ctx context.Context, client *firestore.Client, userId string, locale string) error {
	_, err := client.Collection(users.String()).Doc(userId).Set(ctx, map[string]interface{}{
		"locale": locale,
	}, firestore.MergeAll)
	return err
}

func UpdateUserTimezone(ctx context.Context, client *firestore.Client, userId string, secondsFromUTC int) error {
	_, err := client.Collection(users.String()).Doc(userId).Set(ctx, map[string]interface{}{
		"secondsFromUTC": secondsFromUTC,
	}, firestore.MergeAll)
	return err
}

func SaveLastChatId(ctx context.Context, client *firestore.Client, userId string, chatId int) error {
	_, err := client.Collection(users.String()).Doc(userId).Set(ctx, map[string]interface{}{
		"lastChatId": chatId,
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
