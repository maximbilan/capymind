package telegram

import "strings"

type CallbackParameter string

const (
	Locale   CallbackParameter = "locale"
	Timezone CallbackParameter = "timezone"
)

func (c CallbackParameter) String() string {
	return string(c)
}

func ParseCallbackParameter(input string, param CallbackParameter) (*string, bool) {
	parts := strings.Split(input, "_")
	if len(parts) == 2 && parts[0] == param.String() {
		return &parts[1], true
	}
	return nil, false
}

func GetParameterFmt(param CallbackParameter, value string) string {
	return param.String() + "_" + value
}
