package main

import (
	"context"
	"errors"
	"testing"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/manicar2093/charly_team_api/validators"
	"github.com/stretchr/testify/suite"
)

type MainTests struct {
	suite.Suite
	repo                         *reltest.Repository
	validator                    *mocks.ValidatorService
	paginator                    *mocks.Paginable
	userFilter                   *mocks.UserFilterService
	ctx                          context.Context
	filterFuncMock               mocks.FilterFunc
	ordinaryError, notFoundError error
}

func (c *MainTests) SetupTest() {
	c.repo = reltest.New()
	c.validator = &mocks.ValidatorService{}
	c.paginator = &mocks.Paginable{}
	c.userFilter = &mocks.UserFilterService{}
	c.filterFuncMock = mocks.FilterFunc{}
	c.ctx = context.Background()
	c.ordinaryError = errors.New("An ordinary error :O")
	c.notFoundError = rel.NotFoundError{}

}

func (c *MainTests) TearDownTest() {
	c.repo.AssertExpectations(c.T())
	c.validator.AssertExpectations(c.T())
	c.paginator.AssertExpectations(c.T())
	c.userFilter.AssertExpectations(c.T())
}

func (c *MainTests) TestGetUserByFilter() {
	filterName := "a_filter_name"
	userFilter := UserFilter{FilterName: filterName, Values: "values"}

	c.validator.On(
		"Validate",
		userFilter,
	).Return(
		validators.ValidateOutput{
			IsValid: true,
			Err:     nil,
		},
	)
	c.userFilter.On("GetUserFilter", userFilter.Values).Return(c.filterFuncMock)

	CreateLambdaHandlerWDependencies(c.repo, c.validator, c.paginator, c.userFilter)

}

func TestMain(t *testing.T) {
	suite.Run(t, new(MainTests))
}
