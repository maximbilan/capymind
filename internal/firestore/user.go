package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
)

type User struct {
	ID     string `firestore:"id"`
	Name   string `firestore:"name"`
	Locale string `firestore:"locale"`
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
	return &user.Locale, nil
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
