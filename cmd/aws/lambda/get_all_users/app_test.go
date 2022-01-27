package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/paginator"
	"github.com/manicar2093/charly_team_api/handlers/userfilters/allusersfinder"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(FindBiotestByUUIDAWSLambdaTests))
}

type FindBiotestByUUIDAWSLambdaTests struct {
	suite.Suite
	ctx                  context.Context
	allUsersFinder       *allusersfinder.MockAllUsersFinder
	getAllusersAWSLambda *GetAllUsersAWSLambda
}

func (c *FindBiotestByUUIDAWSLambdaTests) SetupTest() {
	c.ctx = context.Background()
	c.allUsersFinder = &allusersfinder.MockAllUsersFinder{}
	c.getAllusersAWSLambda = NewGetAllUsersAWSLambda(c.allUsersFinder)
}

func (c *FindBiotestByUUIDAWSLambdaTests) TearDownTest() {
	c.allUsersFinder.AssertExpectations(c.T())
}

func (c *FindBiotestByUUIDAWSLambdaTests) TestHandler() {
	request := allusersfinder.AllUsersFinderRequest{PageSort: paginator.PageSort{
		Page: 1,
	}}
	usersPaginator := paginator.Paginator{Data: &[]entities.User{{}, {}}}
	allUsersRunReturned := allusersfinder.AllUsersFinderResponse{UsersFound: &usersPaginator}
	c.allUsersFinder.On("Run", c.ctx, &request).Return(&allUsersRunReturned, nil)

	got, err := c.getAllusersAWSLambda.Handler(c.ctx, request)

	c.Nil(err)
	c.NotNil(got)
	c.Equal(&usersPaginator, got.Body)
}

func (c *FindBiotestByUUIDAWSLambdaTests) TestHandler_UnhandledError() {
	request := allusersfinder.AllUsersFinderRequest{PageSort: paginator.PageSort{
		Page: 1,
	}}
	allUsersErrReturned := fmt.Errorf("ordinary error")
	c.allUsersFinder.On("Run", c.ctx, &request).Return(nil, allUsersErrReturned)

	got, err := c.getAllusersAWSLambda.Handler(c.ctx, request)

	c.Nil(err)
	c.NotNil(got)
	c.IsType(models.ErrorReponse{}, got.Body)
}
