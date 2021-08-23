package services

import (
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

		isCorrect, err := validator.Validate(userToValidate)

		assert.False(t, isCorrect, "validation should not be accepted")
		val, ok := err.(apperrors.ValidationErrors)

		assert.True(t, ok, "error type is not correct")
		assert.Len(t, val, 5, "errors count is not correct")

	})

	t.Run("should return an unexpected error", func(t *testing.T) {
		toValidate := "bad data"

		isCorrect, err := validator.Validate(toValidate)

		assert.False(t, isCorrect, "validation shouldnot be accepted")
		assert.NotNil(t, err, "any error should have existed")
		assert.Equal(t, ErrorUnexpectedValidation, err, "bad error received")
	})

	t.Run("should pass all validation", func(t *testing.T) {
		userToValidate := TestUser{
			FirstName:      "testing",
			LastName:       "testing",
			Age:            5,
			Email:          "testig@testing.com",
			FavouriteColor: "#000000",
		}

		isCorrect, err := validator.Validate(userToValidate)

		assert.True(t, isCorrect, "validation should be accepted")
		assert.Nil(t, err, "any error should have existed")
	})

}
