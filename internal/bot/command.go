package bot

type Command string

const (
	Start    Command = "/start"
	Note     Command = "/note"
	Last     Command = "/last"
	Analysis Command = "/analysis"
	Language Command = "/language"
	Timezone Command = "/timezone"
	Help     Command = "/help"
)
