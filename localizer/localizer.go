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
