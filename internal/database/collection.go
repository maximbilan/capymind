package database

type Collection string

const (
	Users     Collection = "users"
	Notes     Collection = "notes"
	Feedbacks Collection = "feedbacks"
	Updates   Collection = "updates"
)

func (c Collection) String() string {
	return string(c)
}
