package userbyuuidfinder

import (
	"context"
	"fmt"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/validators"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(UserByUUIDFinderTests))
}

type UserByUUIDFinderTests struct {
	suite.Suite
	ctx                  context.Context
	userRepo             *repositories.MockUserRepository
	validator            *validators.MockValidatorService
	userByUUIDFinderImpl *userByUUIDFinderImpl
	faker                faker.Faker
}

func (c *UserByUUIDFinderTests) SetupTest() {
	c.ctx = context.Background()
	c.userRepo = &repositories.MockUserRepository{}
	c.validator = &validators.MockValidatorService{}
	c.userByUUIDFinderImpl = NewUserByUUIDFinderImpl(c.userRepo, c.validator)
	c.faker = faker.New()
}

func (c *UserByUUIDFinderTests) TearDownTest() {
	c.userRepo.AssertExpectations(c.T())
	c.validator.AssertExpectations(c.T())
}

func (c *UserByUUIDFinderTests) TestRun() {
	userUUID := c.faker.UUID().V4()
	request := UserByUUIDFinderRequest{UserUUID: userUUID}
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
	request := UserByUUIDFinderRequest{UserUUID: userUUID}
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
	request := UserByUUIDFinderRequest{UserUUID: userUUID}
	returnErr := fmt.Errorf("ordinary error")
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.userRepo.On("FindUserByUUID", c.ctx, userUUID).Return(nil, returnErr)

	got, err := c.userByUUIDFinderImpl.Run(c.ctx, &request)

	c.NotNil(err)
	c.Nil(got)
	c.Equal(returnErr, err)
}
