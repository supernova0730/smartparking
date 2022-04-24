package apiError

import "net/http"

const (
	ParkingPlaceNotFound = "PARKING_PLACE_NOT_FOUND"
)

var parkingPlaceErrors = []*apiError{
	{
		ID:      ParkingPlaceNotFound,
		Message: "parking place not found",
		Status:  http.StatusNotFound,
	},
}
