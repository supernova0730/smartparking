package apiError

import "net/http"

const (
	EmailVerificationCodeExpired = "EMAIL_VERIFICATION_CODE_EXPIRED"
	EmailVerificationCodeInvalid = "EMAIL_VERIFICATION_CODE_INVALID"
)

var emailVerificationErrors = []*apiError{
	{
		ID:      EmailVerificationCodeExpired,
		Message: "email verification code is expired",
		Status:  http.StatusBadRequest,
	},
	{
		ID:      EmailVerificationCodeInvalid,
		Message: "email verification code is invalid",
		Status:  http.StatusBadRequest,
	},
}
