package bot

type Command string

const (
	Start  Command = "/start"
	Note   Command = "/note"
	Last   Command = "/last"
	Locale Command = "/locale"
	Info   Command = "/info"
	Help   Command = "/help"
)
