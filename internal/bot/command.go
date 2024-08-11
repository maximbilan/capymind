package bot

type Command string

const (
	Start Command = "/start"
	Note  Command = "/note"
	Last  Command = "/last"
	Info  Command = "/info"
	Help  Command = "/help"
)
