package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ValidateStruct(data interface{}) []string {
	validate = validator.New(validator.WithRequiredStructEnabled())
	validationErrors := []string{}

	if ValidationErr := validate.Struct(data); ValidationErr != nil {
		for _, err := range ValidationErr.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%v is %v", err.Field(), err.Tag()))

		}
	}

	return validationErrors
}
