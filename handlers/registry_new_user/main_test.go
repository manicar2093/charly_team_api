package main

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"

	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/validators"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(MainTests))
}

type MainTests struct {
	suite.Suite
	userService MockUserService
	validator   validators.MockValidatorService
	ctx         context.Context
	userCreated *entities.User
}

func (c *MainTests) SetupTest() {
	c.userService = MockUserService{}
	c.validator = validators.MockValidatorService{}
	c.ctx = context.Background()
	c.userCreated = &entities.User{ID: int32(1)}

}

func (c *MainTests) TearDownTest() {
	c.userService.AssertExpectations(c.T())
}

func (c *MainTests) TestRegistryNewUser_Admin() {

	userRequest := models.CreateUserRequest{
		Name:     "testing",
		LastName: "main",
		Email:    "testing@main-func.com",
		Birthday: time.Now(),
		RoleID:   1,
	}

	c.validator.On("Validate", &userRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.userService.On("CreateUser", c.ctx, &userRequest).Return(c.userCreated, nil)

	res, _ := CreateLambdaHandlerWDependencies(&c.userService, &c.validator)(c.ctx, userRequest)

	c.Equal(res.StatusCode, http.StatusCreated, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusCreated), "http status is not correct")

	createUserResponse := res.Body.(*entities.User)

	c.Equal(createUserResponse.ID, c.userCreated.ID, "unexpected id user response")

}

func (c *MainTests) TestRegistryNewUser_Error() {

	userRequest := models.CreateUserRequest{
		Name:     "testing",
		LastName: "main",
		Email:    "testing@main-func.com",
		Birthday: time.Now(),
		RoleID:   1,
	}

	errorText := "an error"

	c.validator.On("Validate", &userRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.userService.On("CreateUser", c.ctx, &userRequest).Return(c.userCreated, errors.New(errorText))

	res, _ := CreateLambdaHandlerWDependencies(&c.userService, &c.validator)(c.ctx, userRequest)

	c.Equal(res.StatusCode, http.StatusInternalServerError, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusInternalServerError), "http status is not correct")

	bodyError := res.Body.(models.ErrorReponse)

	c.Equal(bodyError.Error, errorText, "error does not correspond")

}

func (c *MainTests) TestRegistryNewUser_AdminNoValidReq() {

	userRequest := models.CreateUserRequest{
		Name:     "testing",
		LastName: "main",
		Email:    "testing@main-func.com",
		Birthday: time.Now(),
		RoleID:   1,
	}

	validationErrors := apperrors.ValidationErrors{
		{Field: "name", Validation: "required"},
	}

	c.validator.On("Validate", &userRequest).Return(validators.ValidateOutput{IsValid: false, Err: validationErrors})

	res, _ := CreateLambdaHandlerWDependencies(&c.userService, &c.validator)(c.ctx, userRequest)

	c.Equal(res.StatusCode, http.StatusBadRequest, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusBadRequest), "http status is not correct")

	bodyAsMap := res.Body.(map[string]interface{})

	errorGot, ok := bodyAsMap["error"].(apperrors.ValidationErrors)
	c.True(ok, "error parsing error message")
	c.Equal(len(errorGot), 1, "error message should not be empty")

}

func (c *MainTests) TestRegistryNewUser_Customer() {

	userRequest := models.CreateUserRequest{
		Name:          "testing",
		LastName:      "main",
		Email:         "testing@main-func.com",
		Birthday:      time.Now(),
		RoleID:        3,
		GenderID:      1,
		BoneDensityID: 1,
		BiotypeID:     1,
	}

	c.validator.On("Validate", &userRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.validator.On("Validate", userRequest.GetCustomerValidations()).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.userService.On("CreateUser", c.ctx, &userRequest).Return(c.userCreated, nil)

	res, _ := CreateLambdaHandlerWDependencies(&c.userService, &c.validator)(c.ctx, userRequest)

	c.Equal(res.StatusCode, http.StatusCreated, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusCreated), "http status is not correct")

	createUserResponse := res.Body.(*entities.User)

	c.Equal(createUserResponse.ID, c.userCreated.ID, "unexpected id user response")

}

func (c *MainTests) TestRegistryNewUser_CustomerNoValidReq() {

	userRequest := models.CreateUserRequest{
		Name:     "testing",
		LastName: "main",
		Email:    "testing@main-func.com",
		Birthday: time.Now(),
		RoleID:   3,
	}

	validationErrors := apperrors.ValidationErrors{
		{Field: "bone_density_id", Validation: "required"},
		{Field: "biotype_id", Validation: "required"},
		{Field: "gender_id", Validation: "required"},
	}

	c.validator.On("Validate", &userRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.validator.On("Validate", userRequest.GetCustomerValidations()).Return(validators.ValidateOutput{IsValid: false, Err: validationErrors})

	res, _ := CreateLambdaHandlerWDependencies(&c.userService, &c.validator)(c.ctx, userRequest)

	c.Equal(res.StatusCode, http.StatusBadRequest, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusBadRequest), "http status is not correct")

	bodyAsMap := res.Body.(map[string]interface{})

	errorGot, ok := bodyAsMap["error"].(apperrors.ValidationErrors)
	c.True(ok, "error parsing error message")
	c.Equal(len(errorGot), 3, "error message should not be empty")

}
