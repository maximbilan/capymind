package firestore

type Endpoint string

const (
	users     Endpoint = "users"
	notes     Endpoint = "notes"
	feedbacks Endpoint = "feedbacks"
)

func (e Endpoint) String() string {
	return string(e)
}
