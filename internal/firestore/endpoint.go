package firestore

type Endpoint string

const (
	users Endpoint = "users"
	notes Endpoint = "notes"
)

func (e Endpoint) String() string {
	return string(e)
}
