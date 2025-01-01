package googlemaps

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"googlemaps.github.io/maps"
)

type GoogleMapsService struct{}

var client *maps.Client

func createClient() {
	if client != nil {
		return
	}

	var err error
	client, err = maps.NewClient(maps.WithAPIKey(os.Getenv("CAPY_GOOGLE_MAPS_API_KEY")))
	if err != nil {
		fmt.Println("Error creating client: ", err)
	}
}

func (service GoogleMapsService) GetTimezone(city string) *int {
	createClient()
	context := context.Background()

	r := &maps.GeocodingRequest{
		Address: city,
	}
	results, err := client.Geocode(context, r)
	if err != nil {
		log.Printf("Error getting geocode for city %s: %v\n", city, err)
		return nil
	}

	if len(results) == 0 {
		log.Printf("No results for city %s\n", city)
		return nil
	}

	location := results[0].Geometry.Location

	timezoneRequest := &maps.TimezoneRequest{
		Location:  &maps.LatLng{Lat: location.Lat, Lng: location.Lng},
		Timestamp: time.Now(),
	}

	timezoneResult, err := client.Timezone(context, timezoneRequest)
	if err != nil {
		log.Printf("Error getting timezone for city %s: %v\n", city, err)
		return nil
	}

	if timezoneResult == nil {
		log.Printf("No timezone result for city %s\n", city)
		return nil
	}

	return &timezoneResult.RawOffset
}
