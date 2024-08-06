package localizer

import (
	"encoding/json"
	"os"
)

type Locale string

const (
	EN Locale = "en"
)

var translations map[Locale]map[string]string

func Path() string {
	var path = "localizer/translations.json"
	if os.Getenv("DEBUG_MODE") == "true" {
		path = "./../" + path
	} else {
		path = "./" + path
	}
	return path
}

func Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&translations)
	if err != nil {
		return err
	}
	return nil
}

func Localize(locale Locale, key string) string {
	return translations[locale][key]
}
