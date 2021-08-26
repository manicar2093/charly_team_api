package services

import (
	"errors"
	"net/http"
	"testing"

	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/stretchr/testify/assert"
)

type TestUser struct {
	FirstName      string `json:"name" validate:"required"`
	LastName       string `json:"last_name" validate:"required"`
	Age            uint8  `validate:"gte=2,lte=130"`
	Email          string `validate:"required,email"`
	FavouriteColor string `validate:"iscolor"`
}

func TestStructValidator(t *testing.T) {
	validator := NewStructValidator()

	t.Run("should return correct error count", func(t *testing.T) {
		userToValidate := TestUser{}

		validationOutput := validator.Validate(userToValidate)

		assert.False(t, validationOutput.IsValid, "validation should not be accepted")
		val, ok := validationOutput.Err.(apperrors.ValidationErrors)

		assert.True(t, ok, "error type is not correct")
		assert.Len(t, val, 5, "errors count is not correct")

	})

	t.Run("should return an unexpected error", func(t *testing.T) {
		toValidate := "bad data"

		validationOutput := validator.Validate(toValidate)

		assert.False(t, validationOutput.IsValid, "validation shouldnot be accepted")
		assert.NotNil(t, validationOutput.Err, "an error should have existed")
		assert.Equal(t, ErrorUnexpectedValidation, validationOutput.Err, "bad error received")
	})

	t.Run("should pass all validation", func(t *testing.T) {
		userToValidate := TestUser{
			FirstName:      "testing",
			LastName:       "testing",
			Age:            5,
			Email:          "testig@testing.com",
			FavouriteColor: "#000000",
		}

		validationOutput := validator.Validate(userToValidate)

		assert.True(t, validationOutput.IsValid, "validation should be accepted")
		assert.Nil(t, validationOutput.Err, "any error should have existed")
	})

}

func TestCheckValidationErrors(t *testing.T) {

	t.Run("if validation errors exists should create a response", func(t *testing.T) {
		validationOutput := ValidateOutput{
			false,
			apperrors.ValidationErrors{
				{Field: "name", Validation: "required"},
			},
		}

		isValidGot, responseGot := CheckValidationErrors(validationOutput)

		assert.False(t, isValidGot, "data should not be valid")
		assert.Equal(t, http.StatusBadRequest, responseGot.Code, "response code incorrect")
		assert.Equal(t, http.StatusText(http.StatusBadRequest), responseGot.Status)

		body, _ := responseGot.Body.(map[string]interface{})

		errorsList := body["errors"].(apperrors.ValidationErrors)

		assert.Len(t, errorsList, 1, "error count not correct")

		message := body["message"].(string)
		assert.Contains(t, message, "Please check documentation", "message has not documentation reading warn")

	})

	t.Run("if error is not of required type return need response", func(t *testing.T) {
		validationOutput := ValidateOutput{
			false,
			errors.New("an unexpected error"),
		}

		isValidGot, responseGot := CheckValidationErrors(validationOutput)

		assert.False(t, isValidGot, "data should not be valid")
		assert.Equal(t, http.StatusInternalServerError, responseGot.Code, "response code incorrect")
		assert.Equal(t, http.StatusText(http.StatusInternalServerError), responseGot.Status)

		body, _ := responseGot.Body.(string)

		assert.NotEmpty(t, body, "reponse body should not be empty")
	})

}
