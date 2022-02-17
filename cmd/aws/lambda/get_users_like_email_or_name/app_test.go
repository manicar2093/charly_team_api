package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/manicar2093/charly_team_api/internal/db/entities"
	"github.com/manicar2093/charly_team_api/internal/handlers/userfilters"
	"github.com/manicar2093/charly_team_api/pkg/models"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(GetUsersLikeEmailOrNameAWSLambdaTests))
}

type GetUsersLikeEmailOrNameAWSLambdaTests struct {
	suite.Suite
	ctx                              context.Context
	userLikeEmailOrNameFinder        *userfilters.MockUserLikeEmailOrNameFinder
	getUsersLikeEmailOrNameAWSLambda *GetUsersLikeEmailOrNameAWSLambda
}

func (c *GetUsersLikeEmailOrNameAWSLambdaTests) SetupTest() {
	c.ctx = context.Background()
	c.userLikeEmailOrNameFinder = &userfilters.MockUserLikeEmailOrNameFinder{}
	c.getUsersLikeEmailOrNameAWSLambda = NewGetUsersLikeEmailOrNameAWSLambda(c.userLikeEmailOrNameFinder)
}

func (c *GetUsersLikeEmailOrNameAWSLambdaTests) TearDownTest() {
	c.userLikeEmailOrNameFinder.AssertExpectations(c.T())
}

func (c *GetUsersLikeEmailOrNameAWSLambdaTests) TestHandler() {
	filterData := "name"
	request := userfilters.UserLikeEmailOrNameFinderRequest{FilterData: filterData}
	usersFound := []entities.User{{}, {}}
	userLikeEmailReturned := userfilters.UserLikeEmailOrNameFinderResponse{FetchedData: &usersFound}
	c.userLikeEmailOrNameFinder.On("Run", c.ctx, &request).Return(&userLikeEmailReturned, nil)

	got, err := c.getUsersLikeEmailOrNameAWSLambda.Handler(c.ctx, request)

	c.Nil(err)
	c.NotNil(got)
	c.IsType(&models.Response{}, got)

}

func (c *GetUsersLikeEmailOrNameAWSLambdaTests) TestHandler_UnhandledError() {
	filterData := "name"
	request := userfilters.UserLikeEmailOrNameFinderRequest{FilterData: filterData}
	errorReturned := fmt.Errorf("ordinary error")
	c.userLikeEmailOrNameFinder.On("Run", c.ctx, &request).Return(nil, errorReturned)

	got, err := c.getUsersLikeEmailOrNameAWSLambda.Handler(c.ctx, request)

	c.Nil(err)
	c.NotNil(got)
	c.IsType(&models.Response{}, got)

}
