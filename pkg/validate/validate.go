package validate

import (
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func Struct(s interface{}) error {
	validate := validator.New()
	return validate.Struct(s)
}

func Parse(err error) []ValidationError {
	if _, ok := err.(validator.ValidationErrors); !ok {
		return nil
	}

	validationErrors := make([]ValidationError, 0)
	for _, validationErr := range err.(validator.ValidationErrors) {
		validationErrors = append(validationErrors, ValidationError{
			FailedField: validationErr.StructNamespace(),
			Tag:         validationErr.Tag(),
			Value:       validationErr.Param(),
		})
	}
	return validationErrors
}
