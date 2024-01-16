package validator

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Property string `json:"property"`
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Message  string `json:"message"`
}

func ParseValidationErrors(error error) *[]ValidationError {
	var convertedValidationErrors []ValidationError
	var validationErrors validator.ValidationErrors
	if errors.As(error, &validationErrors) {
		for _, fieldError := range error.(validator.ValidationErrors) {
			var validationError ValidationError
			validationError.Property = fieldError.Field()
			validationError.Tag = fieldError.Tag()
			validationError.Value = fieldError.Param()
			convertedValidationErrors = append(convertedValidationErrors, validationError)
		}
		return &convertedValidationErrors
	}
	return nil
}
