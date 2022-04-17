package main

import (
	"context"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/manicar2093/charly_team_api/internal/db/entities"
	"github.com/manicar2093/charly_team_api/internal/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/userfilters"
	"github.com/manicar2093/charly_team_api/pkg/models"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(GetUserByUUIDAWSLambdaTest))
}

type GetUserByUUIDAWSLambdaTest struct {
	suite.Suite
	ctx                    context.Context
	userByUUIDFinder       *userfilters.MockUserByUUIDFinder
	getUserByUUIDAWSLambda *GetUserByUUIDAWSLambda
	faker                  faker.Faker
}

func (c *GetUserByUUIDAWSLambdaTest) SetupTest() {
	c.ctx = context.Background()
	c.userByUUIDFinder = &userfilters.MockUserByUUIDFinder{}
	c.getUserByUUIDAWSLambda = NewGetUserByUUIDAWSLambda(c.userByUUIDFinder)
	c.faker = faker.New()
}

func (c *GetUserByUUIDAWSLambdaTest) TearDownTest() {
	c.userByUUIDFinder.AssertExpectations(c.T())
}

func (c *GetUserByUUIDAWSLambdaTest) TestHandler() {
	userUUID := c.faker.UUID().V4()
	request := userfilters.UserByUUIDFinderRequest{UserUUID: userUUID}
	userFound := entities.User{UserUUID: userUUID}
	userByUUIDFinderReturn := userfilters.UserByUUIDFinderResponse{UserFound: &userFound}
	c.userByUUIDFinder.On("Run", c.ctx, &request).Return(&userByUUIDFinderReturn, nil)

	got, err := c.getUserByUUIDAWSLambda.Handler(c.ctx, request)

	c.Nil(err)
	c.NotNil(got)
	c.Equal(&userFound, got.Body)
}

func (c *GetUserByUUIDAWSLambdaTest) TestHandler_UserNotFound() {
	userUUID := c.faker.UUID().V4()
	request := userfilters.UserByUUIDFinderRequest{UserUUID: userUUID}
	errorReturned := repositories.NotFoundError{Entity: "User", Identifier: userUUID}
	c.userByUUIDFinder.On("Run", c.ctx, &request).Return(nil, errorReturned)

	got, err := c.getUserByUUIDAWSLambda.Handler(c.ctx, request)

	c.Nil(err)
	c.NotNil(got)
	c.Equal(models.ErrorReponse{Error: errorReturned.Error()}, got.Body)
}
