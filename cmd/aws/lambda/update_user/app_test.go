package main

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/handlers/userupdater"
	"github.com/manicar2093/charly_team_api/internal/apperrors"

	"github.com/manicar2093/charly_team_api/internal/models"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(UpdateUserAWSLambdaTests))
}

type UpdateUserAWSLambdaTests struct {
	suite.Suite
	user                entities.User
	ctx                 context.Context
	userUpdater         *userupdater.MockUserUpdater
	updateUserAWSLambda *UpdateUserAWSLambda
	ordinaryError       error
}

func (c *UpdateUserAWSLambdaTests) SetupTest() {
	c.user = entities.User{}
	c.ctx = context.Background()
	c.userUpdater = &userupdater.MockUserUpdater{}
	c.updateUserAWSLambda = NewUpdateUserAWSLambda(c.userUpdater)
	c.ordinaryError = errors.New("An ordinary error :O")

}

func (c *UpdateUserAWSLambdaTests) TearDownTest() {
	c.userUpdater.AssertExpectations(c.T())
}

func (c *UpdateUserAWSLambdaTests) TestHandler() {
	c.userUpdater.On("Run", c.ctx, &c.user).Return(
		&userupdater.UserUpdaterResponse{UserUpdated: &c.user},
		nil,
	)

	res, err := c.updateUserAWSLambda.Handler(c.ctx, c.user)

	c.Nil(err, "should not return an error")
	c.Equal(http.StatusOK, res.StatusCode, "status code not correct")
	c.IsType(&entities.User{}, res.Body, "body is not correct type")
}

func (c *UpdateUserAWSLambdaTests) TestHandler_ValidationError() {
	validationErrors := apperrors.ValidationErrors{
		{Field: "name", Validation: "required"},
		{Field: "last_name", Validation: "required"},
	}
	c.userUpdater.On("Run", c.ctx, &c.user).Return(
		nil,
		validationErrors,
	)

	res, err := c.updateUserAWSLambda.Handler(c.ctx, c.user)

	bodyAsErrorResponse := res.Body.(models.ErrorReponse)
	c.Nil(err, "should not return an error")
	c.Equal(http.StatusBadRequest, res.StatusCode, "status code not correct")
	c.Len(bodyAsErrorResponse.Error.(apperrors.ValidationErrors), 2, "not correct errors returned")
}

func (c *UpdateUserAWSLambdaTests) TestHandler_UnhandledError() {
	c.userUpdater.On("Run", c.ctx, &c.user).Return(
		nil,
		c.ordinaryError,
	)

	res, err := c.updateUserAWSLambda.Handler(c.ctx, c.user)

	bodyAsErrorResponse := res.Body.(models.ErrorReponse)
	c.Nil(err, "should not return an error")
	c.Equal(http.StatusInternalServerError, res.StatusCode, "status code not correct")
	c.Equal(c.ordinaryError.Error(), bodyAsErrorResponse.Error, "not correct error returned")
}
