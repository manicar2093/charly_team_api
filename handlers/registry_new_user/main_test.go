package main

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/validators"
	"github.com/stretchr/testify/suite"
)

type MainTests struct {
	suite.Suite
	userService   mocks.UserService
	validator     mocks.ValidatorService
	idUserCreated int32
}

func (c *MainTests) SetupTest() {
	c.userService = mocks.UserService{}
	c.validator = mocks.ValidatorService{}
	c.idUserCreated = int32(1)

}

func (c *MainTests) TearDownTest() {
	c.userService.AssertExpectations(c.T())
}

func (c *MainTests) TestRegistryNewUser() {

	userRequest := models.CreateUserRequest{
		Name:     "testing",
		LastName: "main",
		Email:    "testing@main-func.com",
		Birthday: time.Now(),
		RoleID:   3,
	}

	c.validator.On("Validate", userRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.userService.On("CreateUser", &userRequest).Return(c.idUserCreated, nil)

	res := CreateLambdaHandlerWDependencies(&c.userService, &c.validator)(userRequest)

	c.Equal(res.StatusCode, http.StatusCreated, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusCreated), "http status is not correct")

	createUserResponse := res.Body.(models.CreateUserResponse)

	c.Equal(createUserResponse.UserID, c.idUserCreated, "unexpected id user response")

}

func (c *MainTests) TestRegistryNewUserError() {

	userRequest := models.CreateUserRequest{
		Name:     "testing",
		LastName: "main",
		Email:    "testing@main-func.com",
		Birthday: time.Now(),
		RoleID:   3,
	}

	errorText := "an error"

	c.validator.On("Validate", userRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.userService.On("CreateUser", &userRequest).Return(int32(0), errors.New(errorText))

	res := CreateLambdaHandlerWDependencies(&c.userService, &c.validator)(userRequest)

	c.Equal(res.StatusCode, http.StatusInternalServerError, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusInternalServerError), "http status is not correct")

	bodyAsMap := res.Body.(map[string]interface{})

	errorGot, ok := bodyAsMap["error"].(string)
	c.True(ok, "error parsing error data")

	c.Equal(errorGot, errorText, "error does not correspond")

}

func (c *MainTests) TestRegistryNewUserNoValidReq() {

	userRequest := models.CreateUserRequest{
		Name:     "testing",
		LastName: "main",
		Email:    "testing@main-func.com",
		Birthday: time.Now(),
		RoleID:   3,
	}

	validationErrors := apperrors.ValidationErrors{
		{Field: "name", Validation: "required"},
	}

	c.validator.On("Validate", userRequest).Return(validators.ValidateOutput{IsValid: false, Err: validationErrors})

	res := CreateLambdaHandlerWDependencies(&c.userService, &c.validator)(userRequest)

	c.Equal(res.StatusCode, http.StatusBadRequest, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusBadRequest), "http status is not correct")

	bodyAsMap := res.Body.(map[string]interface{})

	errorGot, ok := bodyAsMap["error"].(apperrors.ValidationErrors)
	c.True(ok, "error parsing error message")
	c.Equal(len(errorGot), 1, "error message should not be empty")

}

func TestMain(t *testing.T) {
	suite.Run(t, new(MainTests))
}
