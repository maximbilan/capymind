package botservice

import "github.com/capymind/internal/translator"

type BotResultTextButton struct {
	TextID   string
	Locale   translator.Locale
	Callback string
}

type BotResult struct {
	TextID  string
	Locale  translator.Locale
	Buttons []BotResultTextButton
}

func (result *BotResult) Text() string {
	return translator.Translate(result.Locale, result.TextID)
}

func (button *BotResultTextButton) Text() string {
	return translator.Translate(button.Locale, button.TextID)
}
