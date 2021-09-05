package apperrors

import (
	"fmt"
	"net/http"
)

// ValidationError represents a validation error formated to
// send to the front end
type ValidationError struct {
	Field      string `json:"tag"`
	Validation string `json:"validation"`
}

func (c ValidationError) Error() string {

	return fmt.Sprintf("field '%s' is %s", c.Field, c.Validation)
}

func (c ValidationError) StatusCode() int {
	return http.StatusBadRequest
}

// ValidationErrors is a slice of ValidationError.
// It implements error interface
type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	errorCount := len(v)

	return fmt.Sprintf("error count: %v", errorCount)
}

func (v ValidationErrors) StatusCode() int {
	return http.StatusBadRequest
}
