package botservice

type BotResultTextButton struct {
	TextID   string
	Callback string
}

type BotResult struct {
	TextID  string
	Buttons []BotResultTextButton
}
