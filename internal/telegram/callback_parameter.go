package telegram

type CallbackParameter string

const (
	Locale CallbackParameter = "locale"
)

func (c CallbackParameter) String() string {
	return string(c)
}
