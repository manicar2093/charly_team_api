package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/go-rel/reltest"
	"github.com/manicar2093/health_records/internal/db/entities"
	"github.com/manicar2093/health_records/internal/user"
	"github.com/manicar2093/health_records/mocks"
	"github.com/manicar2093/health_records/pkg/apperrors"
	"github.com/manicar2093/health_records/pkg/validators"
	"github.com/stretchr/testify/suite"
)

func TestUpdater(t *testing.T) {
	suite.Run(t, new(MainTests))
}

type MainTests struct {
	suite.Suite
	repo          *reltest.Repository
	validator     *mocks.ValidatorService
	ctx           context.Context
	userUpdater   user.UserUpdater
	ordinaryError error
}

func (c *MainTests) SetupTest() {
	c.repo = reltest.New()
	c.validator = &mocks.ValidatorService{}
	c.ctx = context.Background()
	c.userUpdater = user.NewUpdateUser(c.repo, c.validator)
	c.ordinaryError = errors.New("An ordinary error :O")

}

func (c *MainTests) TearDownTest() {
	c.validator.AssertExpectations(c.T())
	c.repo.AssertExpectations(c.T())
}

func (c *MainTests) TestUpdateUser() {
	userRequest := entities.User{
		ID: 1,
	}
	c.validator.On("Validate", &userRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.repo.ExpectUpdate().ForType("entities.User").Success()

	res, err := c.userUpdater.Run(c.ctx, &userRequest)

	c.Nil(err, "should return an error")
	c.NotNil(res.UserUpdated, "should not return user updated")

}

func (c *MainTests) TestUpdateUser_UpdateError() {
	userRequest := entities.User{
		ID: 1,
	}
	c.validator.On("Validate", &userRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.repo.ExpectUpdate().ForType("entities.User").Error(c.ordinaryError)

	res, err := c.userUpdater.Run(c.ctx, &userRequest)

	c.NotNil(err, "should return an error")
	c.Nil(res, "should not return user updated")

}
func (c *MainTests) TestUpdateUser_NoUserID() {
	userRequest := entities.User{}

	res, err := c.userUpdater.Run(c.ctx, &userRequest)

	c.Nil(res, "should not return data")
	bodyError := err.(apperrors.ValidationErrors)
	c.Equal("identifier", bodyError[0].Field, "validation error is not correct")
	c.Equal("required", bodyError[0].Validation, "validation error is not correct")

}

func (c *MainTests) TestUpdateUser_NoValidRequest() {
	userRequest := entities.User{
		ID: 1,
	}
	validationErrors := apperrors.ValidationErrors{
		{Field: "name", Validation: "required"},
		{Field: "last_name", Validation: "required"},
	}

	c.validator.On("Validate", &userRequest).Return(validators.ValidateOutput{IsValid: false, Err: validationErrors})

	res, err := c.userUpdater.Run(c.ctx, &userRequest)

	c.Nil(res, "should not return data")
	errorGot, ok := err.(apperrors.ValidationErrors)
	c.True(ok, "error parsing error message")
	c.Equal(len(errorGot), 2, "error message should not be empty")

}
