package apperrors

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckValidationErrors(t *testing.T) {

	t.Run("if validation errors exists should create a response", func(t *testing.T) {
		isValidData, errData := false, ValidationErrors{
			{Field: "name", Validation: "required"},
		}

		isValidGot, responseGot := CheckValidationErrors(isValidData, errData)

		assert.False(t, isValidGot, "data should not be valid")
		assert.Equal(t, http.StatusBadRequest, responseGot.Code, "response code incorrect")
		assert.Equal(t, http.StatusText(http.StatusBadRequest), responseGot.Status)

		body, _ := responseGot.Body.(map[string]interface{})

		var errorsList []map[string]interface{}

		err := json.Unmarshal([]byte(body["errors"].(string)), &errorsList)
		assert.Nil(t, err, "error on unmarshal errors from reponse")
		assert.Len(t, errorsList, 1, "error count not correct")

		message := body["message"].(string)
		assert.Contains(t, message, "Please check documentation", "message has not documentation reading warn")
	})

	t.Run("if error is not of required type return need response", func(t *testing.T) {
		isValidData, errData := false, errors.New("an unexpected error")

		isValidGot, responseGot := CheckValidationErrors(isValidData, errData)

		assert.False(t, isValidGot, "data should not be valid")
		assert.Equal(t, http.StatusInternalServerError, responseGot.Code, "response code incorrect")
		assert.Equal(t, http.StatusText(http.StatusInternalServerError), responseGot.Status)

		body, _ := responseGot.Body.(string)

		assert.NotEmpty(t, body, "reponse body should not be empty")
	})

}
