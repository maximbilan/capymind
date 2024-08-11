package firestore

import (
	"time"

	"cloud.google.com/go/firestore"
)

type Note struct {
	ID        string                 `firestore:"id"`
	Text      string                 `firestore:"text"`
	Timestamp time.Time              `firestore:"timestamp"`
	User      *firestore.DocumentRef `firestore:"user"`
}

type User struct {
	ID   string `firestore:"id"`
	Name string `firestore:"name"`
}

const (
	users string = "users"
	notes string = "notes"
)
