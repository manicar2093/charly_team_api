package apperrors

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/manicar2093/charly_team_api/models"
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

func CheckValidationErrors(isValid bool, err error) (bool, models.Response) {

	if err == nil {
		return true, models.Response{}
	}

	validationErr, ok := err.(ValidationErrors)
	if !ok {
		return false, models.Response{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Body:   err.Error(),
		}
	}

	bytesData, err := json.Marshal(validationErr)
	if err != nil {
		log.Println("Unexpected error marshalling errors: ", err)
		return false, models.Response{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Body:   err.Error(),
		}
	}

	return false, models.Response{
		Code:   http.StatusBadRequest,
		Status: http.StatusText(http.StatusBadRequest),
		Body: map[string]interface{}{
			"message": "Request body does not satisfy needs. Please check documentation",
			"errors":  string(bytesData),
		},
	}
}
