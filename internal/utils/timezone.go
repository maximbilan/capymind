package utils

import (
	"fmt"
	"sort"
)

type TimeZoneInfo struct {
	Offset         int    // UTC offset in hours
	Description    string // Major cities or countries in this time zone
	SecondsFromUTC int    // Seconds from UTC
}

func GetTimeZones() []TimeZoneInfo {
	timezoneMap := map[int]string{
		-12: "Baker Island, Howland Island (Uninhabited)",
		-11: "American Samoa, Niue",
		-10: "Hawaii, French Polynesia (Tahiti)",
		-9:  "Alaska",
		-8:  "Los Angeles, Vancouver, Tijuana",
		-7:  "Denver, Phoenix, Calgary",
		-6:  "Mexico City, Chicago, Guatemala",
		-5:  "New York, Toronto, Lima",
		-4:  "Caracas, La Paz, Manaus",
		-3:  "Buenos Aires, São Paulo, Montevideo",
		-2:  "South Georgia and the South Sandwich Islands",
		-1:  "Azores, Cape Verde",
		0:   "London, Dublin, Lisbon",
		1:   "Berlin, Paris, Rome",
		2:   "Athens, Cairo, Jerusalem",
		3:   "Kyiv, Istanbul, Moscow",
		4:   "Dubai, Baku, Samara",
		5:   "Karachi, Tashkent, Yekaterinburg",
		6:   "Almaty, Dhaka, Omsk",
		7:   "Bangkok, Jakarta, Krasnoyarsk",
		8:   "Beijing, Singapore, Perth",
		9:   "Tokyo, Seoul, Irkutsk",
		10:  "Sydney, Guam, Vladivostok",
		11:  "Solomon Islands, New Caledonia, Magadan",
		12:  "Fiji, Kamchatka, Marshall Islands",
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
