package userfilters_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/manicar2093/health_records/internal/db/entities"
	"github.com/manicar2093/health_records/internal/db/paginator"
	"github.com/manicar2093/health_records/internal/userfilters"
	"github.com/manicar2093/health_records/mocks"
	"github.com/manicar2093/health_records/pkg/apperrors"
	"github.com/manicar2093/health_records/pkg/validators"
	"github.com/stretchr/testify/suite"
)

func TestFindAll(t *testing.T) {
	suite.Run(t, new(AllUsersFinderTests))
}

type AllUsersFinderTests struct {
	suite.Suite
	ctx           context.Context
	userRepo      *mocks.UserRepository
	validator     *mocks.ValidatorService
	allUserFinder *userfilters.AllUsersFinderImpl
}

func (c *AllUsersFinderTests) SetupTest() {
	c.ctx = context.Background()
	c.userRepo = &mocks.UserRepository{}
	c.validator = &mocks.ValidatorService{}
	c.allUserFinder = userfilters.NewAllUsersFinderImpl(c.userRepo, c.validator)
}

func (c *AllUsersFinderTests) TearDownTest() {
	c.userRepo.AssertExpectations(c.T())
	c.validator.AssertExpectations(c.T())
}

func (c *AllUsersFinderTests) TestRun() {
	request := userfilters.AllUsersFinderRequest{PageSort: paginator.PageSort{
		Page: 1,
	}}
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	usersFound := []entities.User{{}, {}}
	findAllUsersReturn := paginator.Paginator{Data: usersFound}
	c.userRepo.On("FindAllUsers", c.ctx, &request.PageSort).Return(&findAllUsersReturn, nil)

	got, err := c.allUserFinder.Run(c.ctx, &request)

	c.Nil(err)
	c.NotNil(got)
	c.Equal(usersFound, got.UsersFound.Data)
}

func (c *AllUsersFinderTests) TestRun_ValidationError() {
	request := userfilters.AllUsersFinderRequest{PageSort: paginator.PageSort{
		Page: 1,
	}}
	validationErrReturn := apperrors.ValidationErrors{{Field: "page", Validation: "required"}}
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{IsValid: false, Err: validationErrReturn})

	got, err := c.allUserFinder.Run(c.ctx, &request)

	c.NotNil(err)
	c.Nil(got)
	c.Equal(validationErrReturn, err)
}

func (c *AllUsersFinderTests) TestRun_UserRepoErr() {
	request := userfilters.AllUsersFinderRequest{PageSort: paginator.PageSort{
		Page: 1,
	}}
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	userRepoErr := fmt.Errorf("ordinary error")
	c.userRepo.On("FindAllUsers", c.ctx, &request.PageSort).Return(nil, userRepoErr)

	got, err := c.allUserFinder.Run(c.ctx, &request)

	c.NotNil(err)
	c.Nil(got)
	c.Equal(userRepoErr, err)
}
