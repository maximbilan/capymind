package localizer

import (
	"encoding/json"
	"os"
)

var translations map[string]map[string]string

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

func Translate(lang, key string) string {
	return translations[lang][key]
}
