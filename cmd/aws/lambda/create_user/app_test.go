package main

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/manicar2093/charly_team_api/internal/db/entities"
	"github.com/manicar2093/charly_team_api/internal/user"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/manicar2093/charly_team_api/pkg/apperrors"
	"github.com/manicar2093/charly_team_api/pkg/models"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(CreateUserAWSLambdaTests))
}

type CreateUserAWSLambdaTests struct {
	suite.Suite
	user                entities.User
	userCreateReq       user.UserCreatorRequest
	ctx                 context.Context
	userCreator         *mocks.UserCreator
	createUserAWSLambda *CreateUserAWSLambda
	ordinaryError       error
}

func (c *CreateUserAWSLambdaTests) SetupTest() {
	c.user = entities.User{}
	c.userCreateReq = user.UserCreatorRequest{}
	c.ctx = context.Background()
	c.userCreator = &mocks.UserCreator{}
	c.createUserAWSLambda = NewUserCreatorAWSLambda(c.userCreator)
	c.ordinaryError = errors.New("An ordinary error :O")

}

func (c *CreateUserAWSLambdaTests) TearDownTest() {
	c.userCreator.AssertExpectations(c.T())
}

func (c *CreateUserAWSLambdaTests) TestHandler() {
	c.userCreator.On("Run", c.ctx, &c.userCreateReq).Return(
		&user.UserCreatorResponse{UserCreated: &c.user},
		nil,
	)

	res, err := c.createUserAWSLambda.Handler(c.ctx, c.userCreateReq)

	c.Nil(err, "should not return an error")
	c.Equal(http.StatusCreated, res.StatusCode, "status code not correct")
	c.IsType(&entities.User{}, res.Body, "body is not correct type")
}

func (c *CreateUserAWSLambdaTests) TestHandler_ValidationError() {
	validationErrors := apperrors.ValidationErrors{
		{Field: "name", Validation: "required"},
		{Field: "last_name", Validation: "required"},
	}
	c.userCreator.On("Run", c.ctx, &c.userCreateReq).Return(
		nil,
		validationErrors,
	)

	res, err := c.createUserAWSLambda.Handler(c.ctx, c.userCreateReq)

	bodyAsErrorResponse := res.Body.(models.ErrorReponse)
	c.Nil(err, "should not return an error")
	c.Equal(http.StatusBadRequest, res.StatusCode, "status code not correct")
	c.Len(bodyAsErrorResponse.Error.(apperrors.ValidationErrors), 2, "not correct errors returned")
}

func (c *CreateUserAWSLambdaTests) TestHandler_UnhandledError() {
	c.userCreator.On("Run", c.ctx, &c.userCreateReq).Return(
		nil,
		c.ordinaryError,
	)

	res, err := c.createUserAWSLambda.Handler(c.ctx, c.userCreateReq)

	bodyAsErrorResponse := res.Body.(models.ErrorReponse)
	c.Nil(err, "should not return an error")
	c.Equal(http.StatusInternalServerError, res.StatusCode, "status code not correct")
	c.Equal(c.ordinaryError.Error(), bodyAsErrorResponse.Error, "not correct error returned")
}
