package app

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/capymind/internal/botservice"
	"github.com/capymind/internal/database"
	"github.com/capymind/internal/mapsservice"
	"github.com/capymind/internal/translator"
	"github.com/capymind/internal/utils"
)

// Handle the timezone command
func handleTimezone(session *Session, settingsStorage database.SettingsStorage) {
	if len(session.Job.Parameters) == 0 {
		requestTimezoneManually(session)
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

// Set the timezone manually
func requestTimezoneManually(session *Session) {
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

func requestTimezone(session *Session) {
	session.User.IsTyping = true
	setOutputText("ask_for_city", session)
}

func finishCityRequest(session *Session, mapsService mapsservice.MapsService, settingsStorage database.SettingsStorage) {
	session.User.IsTyping = false

	city := *session.Job.Input
	secondsFromUTC := mapsService.GetTimezone(city)
	if secondsFromUTC == nil {
		setOutputText("timezone_not_found", session)
		requestTimezoneManually(session)
		return
	}

	session.Settings.Location = &city
	session.SaveSettings(*session.Settings, settingsStorage)

	text := translator.Translate(session.Locale(), "is_this_your_time")
	text = text + currentTimeString(time.Now(), *secondsFromUTC)

	var yesButton botservice.BotResultTextButton = botservice.BotResultTextButton{
		TextID:   "yes",
		Locale:   session.Locale(),
		Callback: string(Timezone) + fmt.Sprintf(" %d", *secondsFromUTC),
	}
	var noButton botservice.BotResultTextButton = botservice.BotResultTextButton{
		TextID:   "no",
		Locale:   session.Locale(),
		Callback: string(Timezone),
	}

	setOutputTextWithButtons(text, []botservice.BotResultTextButton{yesButton, noButton}, session)
}

func currentTimeString(currentTime time.Time, offset int) string {
	utcTime := currentTime.UTC().Add(time.Duration(offset) * time.Second)
	return utcTime.Format("15:04")
}
