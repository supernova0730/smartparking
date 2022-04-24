package apiError

import "net/http"

const (
	ClientNotFound        = "CLIENT_NOT_FOUND"
	ClientWithEmailExists = "CLIENT_WITH_EMAIL_EXISTS"
	ClientWithPhoneExists = "CLIENT_WITH_PHONE_EXISTS"
)

var clientErrors = []*apiError{
	{
		ID:      ClientNotFound,
		Message: "client not found",
		Status:  http.StatusNotFound,
	},
	{
		ID:      ClientWithEmailExists,
		Message: "client with email already exists",
		Status:  http.StatusBadRequest,
	},
	{
		ID:      ClientWithPhoneExists,
		Message: "client with phone already exists",
		Status:  http.StatusBadRequest,
	},
}
