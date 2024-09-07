package scheduler

type MessageType string

const (
	Regular        MessageType = ""
	Morning        MessageType = "morning"
	Evening        MessageType = "evening"
	WeeklyAnalysis MessageType = "weekly_analysis"
)
