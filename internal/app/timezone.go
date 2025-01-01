package app

import (
	"log"
	"strconv"

	"github.com/capymind/internal/botservice"
	"github.com/capymind/internal/database"
	"github.com/capymind/internal/mapsservice"
	"github.com/capymind/internal/utils"
)

// Handle the timezone command
func handleTimezone(session *Session, settingsStorage database.SettingsStorage) {
	if len(session.Job.Parameters) == 0 {
		requestTimezone(session)
	} else {
		setupTimezone(session, settingsStorage)
	}
}

// Set the timezone
func setupTimezone(session *Session, settingsStorage database.SettingsStorage) {
	secondsFromUTC, err := strconv.Atoi(session.Job.Parameters[0])
	if err != nil {
		log.Printf("[Bot] Error parsing timezone: %v", err)
		return
	}
	session.User.SecondsFromUTC = &secondsFromUTC
	session.Settings.SecondsFromUTC = &secondsFromUTC

	if !session.User.IsOnboarded {
		session.User.IsOnboarded = true

		setOutputText("timezone_set", session)
		sendWelcome(session)
	} else {
		setOutputText("timezone_set", session)
	}

	session.SaveSettings(*session.Settings, settingsStorage)
}

// Set the timezone
func requestTimezone(session *Session) {
	var buttons []botservice.BotResultTextButton
	timeZones := utils.GetTimeZones()
	for _, tz := range timeZones {
		callback := string(Timezone) + " " + tz.Parameter()
		button := botservice.BotResultTextButton{
			TextID:   tz.Description,
			Locale:   session.Locale(),
			Callback: callback,
		}
		buttons = append(buttons, button)
	}
	setOutputTextWithButtons("timezone_select", buttons, session)
}

func handleTimezoneByLocation(session *Session, mapsService mapsservice.MapsService) {
	result := mapsService.GetTimezone("Kyiv")
	if result == nil {
		log.Printf("[Bot] Error getting timezone")
		return
	}

	log.Print("[Bot] Timezone: ", *result)
}
