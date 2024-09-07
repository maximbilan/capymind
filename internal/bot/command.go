package bot

type Command string

const (
	Start    Command = "/start"
	Note     Command = "/note"
	Last     Command = "/last"
	Analysis Command = "/analysis"
	Locale   Command = "/locale"
	Timezone Command = "/timezone"
	Help     Command = "/help"
)
