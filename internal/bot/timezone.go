package bot

import (
	"log"
	"strconv"

	"github.com/capymind/internal/utils"
)

// Handle the timezone command
func handleTimezone(session Session) {
	if session.Job.Input == nil {
		requestTimezone(session)
	} else {
		setupTimezone(session)
	}
}

// Set the timezone
func setupTimezone(session Session) {
	secondsFromUTC, err := strconv.Atoi(*session.Job.Input)
	if err != nil {
		log.Printf("[Bot] Error parsing timezone: %v", err)
		return
	}

	session.User.SecondsFromUTC = &secondsFromUTC
	setOutputText("timezone_set", session)
}

// Set the timezone
func requestTimezone(session Session) {
	var buttons []JobResultTextButton
	timeZones := utils.GetTimeZones()
	for _, tz := range timeZones {
		callback := string(Timezone) + " " + tz.String()
		button := JobResultTextButton{
			TextID:   tz.String(),
			Callback: callback,
		}
		buttons = append(buttons, button)
	}
	setOutputTextWithButtons("timezone_set", buttons, session)
}
