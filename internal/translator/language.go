package translator

type Language string

const (
	English   Language = "English 🇺🇸"
	Ukrainian Language = "Українська 🇺🇦"
)

func (l Language) String() string {
	return string(l)
}
