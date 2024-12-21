package validations

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Property string `json:"property"`
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Message  string `json:"message"`
}

func GetValidationErrors(err error) *[]ValidationError {
	var validationErrors []ValidationError
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, err := range err.(validator.ValidationErrors) {
			var item ValidationError
			item.Property = err.Field()
			item.Tag = err.Tag()
			item.Value = err.Param()
			item.Message = customizeValidationErrors(err)
			validationErrors = append(validationErrors, item)
		}
		return &validationErrors
	}

	return nil
}

func customizeValidationErrors(err validator.FieldError) string {
	
	switch err.Tag() {
	case "password":
		return fmt.Sprintf("%s does not match required characters", err.Field())
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", err.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters long", err.Field(), err.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", err.Field(), err.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", err.Field(), err.Param())
	default:
		// Generic fallback for unsupported tags
		return fmt.Sprintf("%s is invalid", err.Field())
	}
}