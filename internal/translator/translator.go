package translator

import (
	"encoding/json"
	"log"
)

var translations map[Locale]map[string]string
var prompts map[Locale]map[string]string
var searchKeywords map[Locale]map[string][]string

func init() {
	if err := json.Unmarshal([]byte(translationsJSON), &translations); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}
	if err := json.Unmarshal([]byte(promptsJSON), &prompts); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}
	if err := json.Unmarshal([]byte(searchKeywordsJSON), &searchKeywords); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}
}

func localize(data *map[Locale]map[string]string, locale Locale, key string) string {
	// Checks if the locale is in the map
	if _, ok := (*data)[locale]; !ok {
		return key
	}

	// Checks if the key is in the map
	if _, ok := (*data)[locale][key]; !ok {
		return key
	}

	// Returns the localized string
	return (*data)[locale][key]
}

func Translate(locale Locale, key string) string {
	return localize(&translations, locale, key)
}

func Prompt(locale Locale, key string) string {
	return localize(&prompts, locale, key)
}

func SearchKeywords(locale Locale, key string) []string {
	// Checks if the locale is in the map
	if _, ok := searchKeywords[locale]; !ok {
		return nil
	}

	// Checks if the key is in the map
	if _, ok := searchKeywords[locale][key]; !ok {
		return nil
	}

	// Returns the search keywords
	return searchKeywords[locale][key]
}
