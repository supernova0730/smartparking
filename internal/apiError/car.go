package apiError

import "net/http"

const (
	CarIDInvalid = "CAR_ID_INVALID"
	CarNotFound  = "CAR_NOT_FOUND"
	NotYourCar   = "NOT_YOUR_CAR"
)

var carErrors = []*apiError{
	{
		ID:      CarIDInvalid,
		Message: "car id is invalid",
		Status:  http.StatusBadRequest,
	},
	{
		ID:      CarNotFound,
		Message: "car not found",
		Status:  http.StatusNotFound,
	},
	{
		ID:      NotYourCar,
		Message: "not your car",
		Status:  http.StatusForbidden,
	},
}
