package main

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/handlers/usercreator"
	"github.com/manicar2093/charly_team_api/internal/apperrors"
	"github.com/manicar2093/charly_team_api/internal/models"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(CreateUserAWSLambdaTests))
}

type CreateUserAWSLambdaTests struct {
	suite.Suite
	user                entities.User
	userCreateReq       usercreator.UserCreatorRequest
	ctx                 context.Context
	userCreator         *usercreator.MockUserCreator
	createUserAWSLambda *CreateUserAWSLambda
	ordinaryError       error
}

func (c *CreateUserAWSLambdaTests) SetupTest() {
	c.user = entities.User{}
	c.userCreateReq = usercreator.UserCreatorRequest{}
	c.ctx = context.Background()
	c.userCreator = &usercreator.MockUserCreator{}
	c.createUserAWSLambda = NewUserCreatorAWSLambda(c.userCreator)
	c.ordinaryError = errors.New("An ordinary error :O")

}

func (c *CreateUserAWSLambdaTests) TearDownTest() {
	c.userCreator.AssertExpectations(c.T())
}

func (c *CreateUserAWSLambdaTests) TestHandler() {
	c.userCreator.On("Run", c.ctx, &c.userCreateReq).Return(
		&usercreator.UserCreatorResponse{UserCreated: &c.user},
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
