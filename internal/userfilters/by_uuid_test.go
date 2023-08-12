package userfilters_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/manicar2093/health_records/internal/db/entities"
	"github.com/manicar2093/health_records/internal/userfilters"
	"github.com/manicar2093/health_records/mocks"
	"github.com/manicar2093/health_records/pkg/apperrors"
	"github.com/manicar2093/health_records/pkg/validators"
	"github.com/stretchr/testify/suite"
)

func TestByUUID(t *testing.T) {
	suite.Run(t, new(UserByUUIDFinderTests))
}

type UserByUUIDFinderTests struct {
	suite.Suite
	ctx                  context.Context
	userRepo             *mocks.UserRepository
	validator            *mocks.ValidatorService
	userByUUIDFinderImpl *userfilters.UserByUUIDFinderImpl
	faker                faker.Faker
}

func (c *UserByUUIDFinderTests) SetupTest() {
	c.ctx = context.Background()
	c.userRepo = &mocks.UserRepository{}
	c.validator = &mocks.ValidatorService{}
	c.userByUUIDFinderImpl = userfilters.NewUserByUUIDFinderImpl(c.userRepo, c.validator)
	c.faker = faker.New()
}

func (c *UserByUUIDFinderTests) TearDownTest() {
	c.userRepo.AssertExpectations(c.T())
	c.validator.AssertExpectations(c.T())
}

func (c *UserByUUIDFinderTests) TestRun() {
	userUUID := c.faker.UUID().V4()
	request := userfilters.UserByUUIDFinderRequest{UserUUID: userUUID}
	userRepoReturn := entities.User{UserUUID: userUUID}
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.userRepo.On("FindUserByUUID", c.ctx, userUUID).Return(&userRepoReturn, nil)

	got, err := c.userByUUIDFinderImpl.Run(c.ctx, &request)

	c.Nil(err)
	c.NotNil(got)
	c.Equal(got.UserFound, &userRepoReturn)
}

func (c *UserByUUIDFinderTests) TestRun_ValidationError() {
	userUUID := c.faker.UUID().V4()
	request := userfilters.UserByUUIDFinderRequest{UserUUID: userUUID}
	validationErr := apperrors.ValidationErrors{
		{Field: "user_uuid", Validation: "required"},
	}
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{IsValid: false, Err: validationErr})

	got, err := c.userByUUIDFinderImpl.Run(c.ctx, &request)

	c.NotNil(err)
	c.Nil(got)
	c.Equal(validationErr, err)
}

func (c *UserByUUIDFinderTests) TestRun_UserRepoErr() {
	userUUID := c.faker.UUID().V4()
	request := userfilters.UserByUUIDFinderRequest{UserUUID: userUUID}
	returnErr := fmt.Errorf("ordinary error")
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.userRepo.On("FindUserByUUID", c.ctx, userUUID).Return(nil, returnErr)

	got, err := c.userByUUIDFinderImpl.Run(c.ctx, &request)

	c.NotNil(err)
	c.Nil(got)
	c.Equal(returnErr, err)
}
