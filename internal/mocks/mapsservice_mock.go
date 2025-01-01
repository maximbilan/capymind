//coverage:ignore file

package mocks

type MapsServiceMock struct{}

func (service MapsServiceMock) GetTimezone(city string) *int {
	timezone := 7200
	return &timezone
}

type InvalidMapsServiceMock struct{}

func (service InvalidMapsServiceMock) GetTimezone(city string) *int {
	return nil
}
