package database

type Collection string

const (
	users     Collection = "users"
	notes     Collection = "notes"
	feedbacks Collection = "feedbacks"
	updates   Collection = "updates"
)

func (c Collection) String() string {
	return string(c)
}
