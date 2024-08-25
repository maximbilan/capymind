package translator

type Language string

const (
	English   Language = "English"
	Ukrainian Language = "Ukrainian"
)

func (l Language) String() string {
	return string(l)
}
