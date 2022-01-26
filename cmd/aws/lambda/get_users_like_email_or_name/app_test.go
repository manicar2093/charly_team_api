package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/handlers/userfilters/userlikeemailornamefinder"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(GetUsersLikeEmailOrNameAWSLambdaTests))
}

type GetUsersLikeEmailOrNameAWSLambdaTests struct {
	suite.Suite
	ctx                              context.Context
	userLikeEmailOrNameFinder        *userlikeemailornamefinder.MockUserLikeEmailOrNameFinder
	getUsersLikeEmailOrNameAWSLambda *GetUsersLikeEmailOrNameAWSLambda
}

func (c *GetUsersLikeEmailOrNameAWSLambdaTests) SetupTest() {
	c.ctx = context.Background()
	c.userLikeEmailOrNameFinder = &userlikeemailornamefinder.MockUserLikeEmailOrNameFinder{}
	c.getUsersLikeEmailOrNameAWSLambda = NewGetUsersLikeEmailOrNameAWSLambda(c.userLikeEmailOrNameFinder)
}

func (c *GetUsersLikeEmailOrNameAWSLambdaTests) TearDownTest() {
	c.userLikeEmailOrNameFinder.AssertExpectations(c.T())
}

func (c *GetUsersLikeEmailOrNameAWSLambdaTests) TestHandler() {
	filterData := "name"
	request := userlikeemailornamefinder.UserLikeEmailOrNameFinderRequest{FilterData: filterData}
	usersFound := []entities.User{{}, {}}
	userLikeEmailReturned := userlikeemailornamefinder.UserLikeEmailOrNameFinderResponse{FetchedData: &usersFound}
	c.userLikeEmailOrNameFinder.On("Run", c.ctx, &request).Return(&userLikeEmailReturned, nil)

	got, err := c.getUsersLikeEmailOrNameAWSLambda.Handler(c.ctx, request)

	c.Nil(err)
	c.NotNil(got)
	c.IsType(&models.Response{}, got)

}

func (c *GetUsersLikeEmailOrNameAWSLambdaTests) TestHandler_UnhandledError() {
	filterData := "name"
	request := userlikeemailornamefinder.UserLikeEmailOrNameFinderRequest{FilterData: filterData}
	errorReturned := fmt.Errorf("ordinary error")
	c.userLikeEmailOrNameFinder.On("Run", c.ctx, &request).Return(nil, errorReturned)

	got, err := c.getUsersLikeEmailOrNameAWSLambda.Handler(c.ctx, request)

	c.Nil(err)
	c.NotNil(got)
	c.IsType(&models.Response{}, got)

}
