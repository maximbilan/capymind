//coverage:ignore file

package mapsservice

type MapsService interface {
	GetTimezone(city string) *int
}
