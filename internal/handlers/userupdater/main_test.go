package userupdater

import (
	"context"
	"errors"
	"testing"

	"github.com/go-rel/rel/reltest"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/pkg/apperrors"
	"github.com/manicar2093/charly_team_api/pkg/validators"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(MainTests))
}

type MainTests struct {
	suite.Suite
	repo          *reltest.Repository
	validator     *validators.MockValidatorService
	ctx           context.Context
	userUpdater   UserUpdater
	ordinaryError error
}

func (c *MainTests) SetupTest() {
	c.repo = reltest.New()
	c.validator = &validators.MockValidatorService{}
	c.ctx = context.Background()
	c.userUpdater = NewUpdateUser(c.repo, c.validator)
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
	c.repo.ExpectUpdate().ForType("entities.User").Return(nil)

	res, err := c.userUpdater.Run(c.ctx, &userRequest)

	c.Nil(err, "should return an error")
	c.NotNil(res.UserUpdated, "should not return user updated")

}

func (c *MainTests) TestUpdateUser_UpdateError() {
	userRequest := entities.User{
		ID: 1,
	}
	c.validator.On("Validate", &userRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.repo.ExpectUpdate().ForType("entities.User").Return(c.ordinaryError)

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
