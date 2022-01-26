package main

import (
	"context"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/handlers/userfilters/userbyuuidfinder"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(GetUserByUUIDAWSLambdaTest))
}

type GetUserByUUIDAWSLambdaTest struct {
	suite.Suite
	ctx                    context.Context
	userByUUIDFinder       *userbyuuidfinder.MockUserByUUIDFinder
	getUserByUUIDAWSLambda *GetUserByUUIDAWSLambda
	faker                  faker.Faker
}

func (c *GetUserByUUIDAWSLambdaTest) SetupTest() {
	c.ctx = context.Background()
	c.userByUUIDFinder = &userbyuuidfinder.MockUserByUUIDFinder{}
	c.getUserByUUIDAWSLambda = NewGetUserByUUIDAWSLambda(c.userByUUIDFinder)
	c.faker = faker.New()
}

func (c *GetUserByUUIDAWSLambdaTest) TearDownTest() {
	c.userByUUIDFinder.AssertExpectations(c.T())
}

func (c *GetUserByUUIDAWSLambdaTest) TestHandler() {
	userUUID := c.faker.UUID().V4()
	request := userbyuuidfinder.UserByUUIDFinderRequest{UserUUID: userUUID}
	userFound := entities.User{UserUUID: userUUID}
	userByUUIDFinderReturn := userbyuuidfinder.UserByUUIDFinderResponse{UserFound: &userFound}
	c.userByUUIDFinder.On("Run", c.ctx, &request).Return(&userByUUIDFinderReturn, nil)

	got, err := c.getUserByUUIDAWSLambda.Handler(c.ctx, request)

	c.Nil(err)
	c.NotNil(got)
	c.Equal(&userFound, got.Body)
}

func (c *GetUserByUUIDAWSLambdaTest) TestHandler_UserNotFound() {
	userUUID := c.faker.UUID().V4()
	request := userbyuuidfinder.UserByUUIDFinderRequest{UserUUID: userUUID}
	errorReturned := repositories.NotFoundError{Entity: "User", Identifier: userUUID}
	c.userByUUIDFinder.On("Run", c.ctx, &request).Return(nil, errorReturned)

	got, err := c.getUserByUUIDAWSLambda.Handler(c.ctx, request)

	c.Nil(err)
	c.NotNil(got)
	c.Equal(models.ErrorReponse{Error: errorReturned.Error()}, got.Body)
}
