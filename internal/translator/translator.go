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
	// Checks if the locale is in the translations map
	if _, ok := translations[locale]; !ok {
		log.Printf("Locale %s not found in translations", locale)
		return key
	}

	// Checks if the key is in the translations map
	if _, ok := translations[locale][key]; !ok {
		log.Printf("Key %s not found in translations", key)
		return key
	}

	// Returns the translation
	return translations[locale][key]
}
