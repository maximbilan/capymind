package translator

type Locale string

const (
	EN Locale = "en"
	UK Locale = "uk"
)

func (l Locale) String() string {
	return string(l)
}
