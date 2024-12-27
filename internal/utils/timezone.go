package utils

import (
	"fmt"
	"sort"
)

type TimeZoneInfo struct {
	Offset         int    // UTC offset in hours
	Description    string // ex. "GMT +1"
	SecondsFromUTC int    // Seconds from UTC
}

// Return a list of time zones
func GetTimeZones() []TimeZoneInfo {
	gmtValues := []int{-11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	var timeZones []TimeZoneInfo
	for _, value := range gmtValues {
		timeZones = append(timeZones, TimeZoneInfo{
			Offset:         value,
			Description:    fmt.Sprintf("GMT %+d", value),
			SecondsFromUTC: value * 3600,
		})
	}

	sort.Slice(timeZones, func(i, j int) bool {
		return timeZones[i].Offset < timeZones[j].Offset
	})

	return timeZones
}

// Return the time zone info as a parameter. Example: "-25200"
func (info TimeZoneInfo) Parameter() string {
	return fmt.Sprintf("%d", info.SecondsFromUTC)
}
