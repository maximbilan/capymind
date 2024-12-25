package botservice

type BotMessage struct {
	UserID       string
	ChatID       int64
	UserName     string
	FirstName    string
	LastName     string
	LanguageCode string
	Text         string
}
