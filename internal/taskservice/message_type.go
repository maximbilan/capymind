package taskservice

import "time"

type MessageType string

const (
	Regular        MessageType = ""
	Morning        MessageType = "morning"
	Evening        MessageType = "evening"
	WeeklyAnalysis MessageType = "weekly_analysis"
	UserStats      MessageType = "user_stats"
)

func GetMessage(messageType MessageType, weekday time.Weekday) string {
	if weekday < 0 || weekday > 6 {
		return ""
	}

	var dayStr string
	switch weekday {
	case time.Monday:
		dayStr = "monday"
	case time.Tuesday:
		dayStr = "tuesday"
	case time.Wednesday:
		dayStr = "wednesday"
	case time.Thursday:
		dayStr = "thursday"
	case time.Friday:
		dayStr = "friday"
	case time.Saturday:
		dayStr = "saturday"
	case time.Sunday:
		dayStr = "sunday"
	}

	switch messageType {
	case Morning:
		return "how_are_you_morning_" + dayStr
	case Evening:
		return "how_are_you_evening_" + dayStr
	}
	return ""
}
