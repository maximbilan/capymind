package firestore

type Endpoint string

const (
	users     Endpoint = "users"
	notes     Endpoint = "notes"
	feedbacks Endpoint = "feedbacks"
	updates   Endpoint = "updates"
)

func (e Endpoint) String() string {
	return string(e)
}
