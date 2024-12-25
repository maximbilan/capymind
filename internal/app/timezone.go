package app

import (
	"log"
	"strconv"

	"github.com/capymind/internal/botservice"
	"github.com/capymind/internal/utils"
)

// Handle the timezone command
func handleTimezone(session *Session) {
	if len(session.Job.Parameters) == 0 {
		requestTimezone(session)
	} else {
		setupTimezone(session)
	}
}

// Set the timezone
func setupTimezone(session *Session) {
	secondsFromUTC, err := strconv.Atoi(session.Job.Parameters[0])
	if err != nil {
		log.Printf("[Bot] Error parsing timezone: %v", err)
		return
	}
	session.User.SecondsFromUTC = &secondsFromUTC

	if !session.User.IsOnboarded {
		session.User.IsOnboarded = true

		setOutputText("timezone_set", session)
		sendWelcome(session)
	} else {
		setOutputText("timezone_set", session)
	}
}

// Set the timezone
func requestTimezone(session *Session) {
	var buttons []botservice.BotResultTextButton
	timeZones := utils.GetTimeZones()
	for _, tz := range timeZones {
		callback := string(Timezone) + " " + tz.Parameter()
		button := botservice.BotResultTextButton{
			TextID:   tz.String(),
			Locale:   session.Locale(),
			Callback: callback,
		}
		buttons = append(buttons, button)
	}
	setOutputTextWithButtons("timezone_select", buttons, session)
}