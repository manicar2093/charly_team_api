package apperrors

import (
	"fmt"
)

// ValidationError represents a validation error formated to
// send to the front end
type ValidationError struct {
	Field      string `json:"tag"`
	Validation string `json:"validation"`
}

// ValidationErrors is a slice of ValidationError.
// It implements error interface
type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	errorCount := len(v)

	return fmt.Sprintf("error count: %v", errorCount)
}
