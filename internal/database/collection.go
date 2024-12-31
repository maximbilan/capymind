package database

type Collection string

const (
	Users              Collection = "users"
	Notes              Collection = "notes"
	Feedbacks          Collection = "feedbacks"
	SettingsCollection Collection = "settings"
)

func (c Collection) String() string {
	return string(c)
}
