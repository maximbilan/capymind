package firestore

type User struct {
	ID     string `firestore:"id"`
	Name   string `firestore:"name"`
	Locale string `firestore:"locale"`
}
