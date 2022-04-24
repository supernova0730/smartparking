package apiError

import "net/http"

const (
	TaxNotFound = "TAX_NOT_FOUND"
)

var taxErrors = []*apiError{
	{
		ID:      TaxNotFound,
		Message: "tax not found",
		Status:  http.StatusNotFound,
	},
}
