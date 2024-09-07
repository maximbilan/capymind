package utils

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type TimeZoneInfo struct {
	Offset         int    // UTC offset in hours
	Description    string // Major cities or countries in this time zone
	SecondsFromUTC int    // Seconds from UTC
}

func GetTimeZones() []TimeZoneInfo {
	timezoneMap := map[int]string{
		-12: "Baker Island",
		-11: "American Samoa, Niue",
		-10: "Hawaii, Tahiti",
		-9:  "Alaska",
		-8:  "Los Angeles, Vancouver, Tijuana",
		-7:  "Denver, Phoenix, Calgary",
		-6:  "Mexico City, Chicago, Guatemala",
		-5:  "New York, Toronto, Lima",
		-4:  "Caracas, La Paz, Manaus",
		-3:  "Buenos Aires, São Paulo",
		-2:  "South Georgia and the South Sandwich Islands",
		-1:  "Azores, Cape Verde",
		0:   "London, Dublin, Lisbon",
		1:   "Berlin, Paris, Rome",
		2:   "Athens, Cairo, Jerusalem",
		3:   "Kyiv, Istanbul, Helsinki",
		4:   "Dubai, Baku, Samara",
		5:   "Karachi, Tashkent",
		6:   "Almaty, Dhaka",
		7:   "Bangkok, Jakarta",
		8:   "Beijing, Singapore, Perth",
		9:   "Tokyo, Seoul",
		10:  "Sydney, Guam",
		11:  "Solomon Islands, New Caledonia",
		12:  "Fiji, Marshall Islands",
	}

	var timeZones []TimeZoneInfo
	for offset, description := range timezoneMap {
		timeZones = append(timeZones, TimeZoneInfo{
			Offset:         offset,
			Description:    description,
			SecondsFromUTC: offset * 3600,
		})
	}

	sort.Slice(timeZones, func(i, j int) bool {
		return timeZones[i].Offset < timeZones[j].Offset
	})

	return timeZones
}

func (info TimeZoneInfo) String() string {
	return fmt.Sprintf("UTC %+d - %s", info.Offset, info.Description)
}

func GetTimezoneParameter(info TimeZoneInfo) string {
	return "timezone_" + fmt.Sprintf("%d", info.SecondsFromUTC)
}

func ParseTimezone(input string) (*int, bool) {
	parts := strings.Split(input, "_")
	if len(parts) == 2 && parts[0] == "timezone" {
		secondsFromUTC, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, false
		}
		return &secondsFromUTC, true
	}
	return nil, false
}
