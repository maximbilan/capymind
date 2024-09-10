package bot

import "github.com/capymind/internal/utils"

// Set the timezone
func setupTimezone(session Session) {

	// TO DO: Implement parsing and saving the timezone to the database

	secondsFromUTC, ok := utils.ParseTimezone(*session.Job.Input)

	if ok && secondsFromUTC != nil {
		setupTimezone(*userID, *secondsFromUTC)
		localizeAndSendMessage(*chatID, *userID, locale, "timezone_set")
		if !userExists(*userID) {
			sendStartMessage(*chatID, *userID, &callbackQuery.From.UserName, locale)
		}
		return
	}
}

// Handle the timezone command
func handleTimezone(session Session) {
	if session.Job.Input == nil {
		requestTimezone(session)
	} else {
		setupTimezone(session)
	}
}

// Set the timezone
func requestTimezone(session Session) {
	var buttons []JobResultTextButton
	timeZones := utils.GetTimeZones()
	for _, tz := range timeZones {
		callback := utils.GetTimezoneParameter(tz)
		button := JobResultTextButton{
			TextID:   tz.String(),
			Callback: callback,
		}
		buttons = append(buttons, button)
	}
	setOutputTextWithButtons("timezone_set", buttons, session)
}
