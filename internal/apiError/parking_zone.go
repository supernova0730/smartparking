package apiError

import "net/http"

const (
	ParkingZoneNotFound = "PARKING_ZONE_NOT_FOUND"
)

var parkingZoneErrors = []*apiError{
	{
		ID:      ParkingZoneNotFound,
		Message: "parking zone not found",
		Status:  http.StatusNotFound,
	},
}
