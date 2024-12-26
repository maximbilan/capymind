package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/capymind/internal/database"
)

type UserStorage struct{}

// Get a user from the database
func (storage UserStorage) GetUser(ctx *context.Context, userID string) (*database.User, error) {
	doc, err := client.Collection(database.Users.String()).Doc(userID).Get(*ctx)
	if err != nil {
		return nil, err
	}

	var user database.User
	doc.DataTo(&user)
	return &user, nil
}

// Save a user to the database
func (storage UserStorage) SaveUser(ctx *context.Context, user database.User) error {
	_, err := client.Collection(database.Users.String()).Doc(user.ID).Set(*ctx, user)
	return err
}

// Delete a user from the database
func (storage UserStorage) DeleteUser(ctx *context.Context, userID string) error {
	_, err := client.Collection(database.Users.String()).Doc(userID).Delete(*ctx)
	return err
}

// Iterate over all users
func (storage UserStorage) ForEachUser(ctx *context.Context, callback func([]database.User) error) error {
	var lastDoc *firestore.DocumentSnapshot
	for {
		query := client.Collection(database.Users.String()).OrderBy(firestore.DocumentID, firestore.Asc).Limit(100)
		if lastDoc != nil {
			query = query.StartAfter(lastDoc)
		}

		docs, err := query.Documents(*ctx).GetAll()
		if err != nil {
			return err
		}

		var users []database.User
		for _, doc := range docs {
			var user database.User
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
