package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/manicar2093/charly_team_api/internal/db/entities"
	"github.com/manicar2093/charly_team_api/internal/db/paginator"
	"github.com/manicar2093/charly_team_api/internal/userfilters"
	"github.com/manicar2093/charly_team_api/pkg/models"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(FindBiotestByUUIDAWSLambdaTests))
}

type FindBiotestByUUIDAWSLambdaTests struct {
	suite.Suite
	ctx                  context.Context
	allUsersFinder       *userfilters.MockAllUsersFinder
	getAllusersAWSLambda *GetAllUsersAWSLambda
}

func (c *FindBiotestByUUIDAWSLambdaTests) SetupTest() {
	c.ctx = context.Background()
	c.allUsersFinder = &userfilters.MockAllUsersFinder{}
	c.getAllusersAWSLambda = NewGetAllUsersAWSLambda(c.allUsersFinder)
}

func (c *FindBiotestByUUIDAWSLambdaTests) TearDownTest() {
	c.allUsersFinder.AssertExpectations(c.T())
}

func (c *FindBiotestByUUIDAWSLambdaTests) TestHandler() {
	request := userfilters.AllUsersFinderRequest{PageSort: paginator.PageSort{
		Page: 1,
	}}
	usersPaginator := paginator.Paginator{Data: &[]entities.User{{}, {}}}
	allUsersRunReturned := userfilters.AllUsersFinderResponse{UsersFound: &usersPaginator}
	c.allUsersFinder.On("Run", c.ctx, &request).Return(&allUsersRunReturned, nil)

	got, err := c.getAllusersAWSLambda.Handler(c.ctx, request)

	c.Nil(err)
	c.NotNil(got)
	c.Equal(&usersPaginator, got.Body)
}

func (c *FindBiotestByUUIDAWSLambdaTests) TestHandler_UnhandledError() {
	request := userfilters.AllUsersFinderRequest{PageSort: paginator.PageSort{
		Page: 1,
	}}
	allUsersErrReturned := fmt.Errorf("ordinary error")
	c.allUsersFinder.On("Run", c.ctx, &request).Return(nil, allUsersErrReturned)

	got, err := c.getAllusersAWSLambda.Handler(c.ctx, request)

	c.Nil(err)
	c.NotNil(got)
	c.IsType(models.ErrorReponse{}, got.Body)
}
