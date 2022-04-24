package apiError

import "net/http"

const (
	InvalidPassword = "INVALID_PASSWORD"
	TokenInvalid    = "TOKEN_INVALID"
	TokenExpired    = "TOKEN_EXPIRED"
)

var authErrors = []*apiError{
	{
		ID:      InvalidPassword,
		Message: "Invalid Password",
		Status:  http.StatusBadRequest,
	},
	{
		ID:      TokenInvalid,
		Message: "Invalid Token",
		Status:  http.StatusUnauthorized,
	},
	{
		ID:      TokenExpired,
		Message: "token was expired",
		Status:  http.StatusUnauthorized,
	},
}
