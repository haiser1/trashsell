package helper

import (
	"strings"

	"github.com/go-playground/validator"
)

// Function to handle validation errors
func HandleValidationErrors(err error) []string {
	// Cast error to ValidationErrors type
	validationErrors := err.(validator.ValidationErrors)

	// Initialize slice to store error messages
	errorMessages := make([]string, 0)

	// Iterate through validation errors
	for _, e := range validationErrors {
		// Generate error message using field and tag
		errorMessage := strings.ToLower(e.Field()) + " " + e.Tag() + ": " + e.Param()

		// Append error message to slice
		errorMessages = append(errorMessages, errorMessage)
	}

	return errorMessages
}
