package userlikeemailornamefinder

import (
	"context"
	"fmt"
	"testing"

	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/validators"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(UserLikeEmailOrNameFinderTests))
}

type UserLikeEmailOrNameFinderTests struct {
	suite.Suite
	ctx                       context.Context
	userRepo                  *repositories.MockUserRepository
	validator                 *validators.MockValidatorService
	userLikeEmailOrNameFinder *userLikeEmailOrNameFinderImpl
}

func (c *UserLikeEmailOrNameFinderTests) SetupTest() {
	c.ctx = context.Background()
	c.userRepo = &repositories.MockUserRepository{}
	c.validator = &validators.MockValidatorService{}
	c.userLikeEmailOrNameFinder = NewUserLikeEmailOrNameFinderImpl(c.userRepo, c.validator)
}

func (c *UserLikeEmailOrNameFinderTests) TearDownTest() {
	c.userRepo.AssertExpectations(c.T())
	c.validator.AssertExpectations(c.T())
}

func (c *UserLikeEmailOrNameFinderTests) TestHandler() {
	filterData := "name"
	request := UserLikeEmailOrNameFinderRequest{FilterData: filterData}
	usersFound := []entities.User{{}, {}}
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.userRepo.On("FindUserLikeEmailOrNameOrLastName", c.ctx, filterData).Return(&usersFound, nil)

	got, err := c.userLikeEmailOrNameFinder.Run(c.ctx, &request)

	c.Nil(err)
	c.NotNil(got)
	c.Equal(&usersFound, got.FetchedData)
}

func (c *UserLikeEmailOrNameFinderTests) TestHandler_ValidationErr() {
	filterData := "name"
	request := UserLikeEmailOrNameFinderRequest{FilterData: filterData}
	validationErrReturned := fmt.Errorf("ordinary error")
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{IsValid: false, Err: validationErrReturned})

	got, err := c.userLikeEmailOrNameFinder.Run(c.ctx, &request)

	c.NotNil(err)
	c.Nil(got)
	c.Equal(validationErrReturned, err)
}

func (c *UserLikeEmailOrNameFinderTests) TestHandler_RepoError() {
	filterData := "name"
	request := UserLikeEmailOrNameFinderRequest{FilterData: filterData}
	userRepoErrReturned := fmt.Errorf("ordinary error")
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.userRepo.On("FindUserLikeEmailOrNameOrLastName", c.ctx, filterData).Return(nil, userRepoErrReturned)

	got, err := c.userLikeEmailOrNameFinder.Run(c.ctx, &request)

	c.NotNil(err)
	c.Nil(got)
	c.Equal(userRepoErrReturned, err)
}
