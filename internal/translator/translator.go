package translator

import (
	"encoding/json"
	"log"
)

var translations map[Locale]map[string]string

func init() {
	if err := json.Unmarshal([]byte(translationsJSON), &translations); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}
}

func Translate(locale Locale, key string) string {
	return translations[locale][key]
}
