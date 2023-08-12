package validators

import (
	"errors"
	"net/http"
	"testing"

	"github.com/manicar2093/health_records/pkg/apperrors"
	"github.com/manicar2093/health_records/pkg/models"
	"github.com/stretchr/testify/assert"
)

type TestNastedData struct {
	RoleID   int `validate:"required" json:"role_id,omitempty"`
	GenderID int `json:"gender_id,omitempty"`
}

type TestUser struct {
	FirstName      string         `json:"name,omitempty" validate:"required"`
	LastName       string         `json:"last_name,omitempty" validate:"required"`
	Age            uint8          `validate:"gte=2,lte=130" json:"age,omitempty"`
	Email          string         `validate:"required,email" json:"email,omitempty"`
	FavouriteColor string         `validate:"iscolor" json:"favourite_color,omitempty"`
	NestedData     TestNastedData `validate:"-" json:"nested_data,omitempty"`
	MoreNestedData TestNastedData `json:"mode_nested_data,omitempty"`
}

func TestStructValidator(t *testing.T) {
	validator := NewStructValidator()

	t.Run("should return correct error count", func(t *testing.T) {
		userToValidate := TestUser{}

		validationOutput := validator.Validate(userToValidate)

		assert.False(t, validationOutput.IsValid, "validation should not be accepted")
		val, ok := validationOutput.Err.(apperrors.ValidationErrors)

		assert.True(t, ok, "error type is not correct")
		assert.Len(t, val, 6, "errors count is not correct")

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
			MoreNestedData: TestNastedData{RoleID: 1},
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
		assert.Equal(t, http.StatusBadRequest, responseGot.StatusCode, "response code incorrect")
		assert.Equal(t, http.StatusText(http.StatusBadRequest), responseGot.Status)

		body, _ := responseGot.Body.(map[string]interface{})

		errorsList := body["error"].(apperrors.ValidationErrors)

		assert.Len(t, errorsList, 1, "error count not correct")

	})

	t.Run("if error is not of required type return internal server error response", func(t *testing.T) {
		validationOutput := ValidateOutput{
			false,
			errors.New("an unexpected error"),
		}

		isValidGot, responseGot := CheckValidationErrors(validationOutput)

		assert.False(t, isValidGot, "data should not be valid")
		assert.Equal(t, http.StatusInternalServerError, responseGot.StatusCode, "response code incorrect")
		assert.Equal(t, http.StatusText(http.StatusInternalServerError), responseGot.Status)

		body, _ := responseGot.Body.(models.ErrorReponse)

		assert.NotEmpty(t, body, "reponse body should not be empty")
	})

	t.Run("if there is any error return true validation", func(t *testing.T) {
		validationOutput := ValidateOutput{
			true,
			nil,
		}

		isValidGot, responseGot := CheckValidationErrors(validationOutput)

		assert.True(t, isValidGot, "data should be valid")
		assert.Empty(t, responseGot, "response should be empty")
	})

}
