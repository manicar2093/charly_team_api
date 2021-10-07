package main

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/go-rel/rel/reltest"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/manicar2093/charly_team_api/validators"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(MainTests))
}

type MainTests struct {
	suite.Suite
	repo          *reltest.Repository
	validator     mocks.ValidatorService
	ctx           context.Context
	ordinaryError error
}

func (c *MainTests) SetupTest() {
	c.repo = reltest.New()
	c.validator = mocks.ValidatorService{}
	c.ctx = context.Background()
	c.ordinaryError = errors.New("An ordinary error :O")

}

func (c *MainTests) TearDownTest() {
	c.validator.AssertExpectations(c.T())
	c.repo.AssertExpectations(c.T())
}

func (c *MainTests) TestUpdateUser_UpdateError() {

	userRequest := entities.User{
		ID: 1,
	}

	c.validator.On("Validate", userRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.repo.ExpectUpdate().ForType("entities.User").Return(c.ordinaryError)

	res, _ := CreateLambdaHandlerWDependencies(c.repo, &c.validator)(c.ctx, userRequest)

	c.Equal(res.StatusCode, http.StatusInternalServerError, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusInternalServerError), "http status is not correct")

}

func (c *MainTests) TestUpdateUser() {

	userRequest := entities.User{
		ID: 1,
	}

	c.validator.On("Validate", userRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.repo.ExpectUpdate().ForType("entities.User").Return(nil)

	res, _ := CreateLambdaHandlerWDependencies(c.repo, &c.validator)(c.ctx, userRequest)

	c.Equal(res.StatusCode, http.StatusOK, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusOK), "http status is not correct")

}

func (c *MainTests) TestUpdateUser_NoUserID() {

	userRequest := entities.User{}

	res, _ := CreateLambdaHandlerWDependencies(c.repo, &c.validator)(c.ctx, userRequest)

	c.Equal(res.StatusCode, http.StatusBadRequest, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusBadRequest), "http status is not correct")

	bodyError := res.Body.(apperrors.ValidationErrors)
	c.Equal("identifier", bodyError[0].Field, "validation error is not correct")
	c.Equal("required", bodyError[0].Validation, "validation error is not correct")

}

func (c *MainTests) TestUpdateUser_NoValidRequest() {

	userRequest := entities.User{
		ID: 1,
	}

	validationErrors := apperrors.ValidationErrors{
		{Field: "name", Validation: "required"},
		{Field: "last_name", Validation: "required"},
	}

	c.validator.On("Validate", userRequest).Return(validators.ValidateOutput{IsValid: false, Err: validationErrors})

	res, _ := CreateLambdaHandlerWDependencies(c.repo, &c.validator)(c.ctx, userRequest)

	c.Equal(res.StatusCode, http.StatusBadRequest, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusBadRequest), "http status is not correct")

	bodyAsMap := res.Body.(map[string]interface{})

	errorGot, ok := bodyAsMap["error"].(apperrors.ValidationErrors)
	c.True(ok, "error parsing error message")
	c.Equal(len(errorGot), 2, "error message should not be empty")

}
