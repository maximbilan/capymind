package translator

import "strings"

func ParseLocale(input string) (*string, bool) {
	parts := strings.Split(input, "_")
	if len(parts) == 2 && parts[0] == "locale" {
		return &parts[1], true
	}
	return nil, false
}

func GetLocaleParameter(locale Locale) string {
	return "locale_" + locale.String()
}
