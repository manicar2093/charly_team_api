package user_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/manicar2093/charly_team_api/internal/db/entities"
	"github.com/manicar2093/charly_team_api/internal/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/user"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/manicar2093/charly_team_api/pkg/apperrors"
	"github.com/manicar2093/charly_team_api/pkg/validators"
	"github.com/stretchr/testify/suite"
)

func TestAvatarUpdater(t *testing.T) {
	suite.Run(t, new(AvatarUpdaterTests))
}

type AvatarUpdaterTests struct {
	suite.Suite
	ctx           context.Context
	userRepo      *mocks.UserRepository
	validator     *mocks.ValidatorService
	avatarUpdater user.AvatarUpdater
	userID        int32
	userUUID      string
	user          entities.User
}

func (c *AvatarUpdaterTests) SetupTest() {
	c.ctx = context.Background()
	c.userRepo = &mocks.UserRepository{}
	c.validator = &mocks.ValidatorService{}
	c.avatarUpdater = user.NewUserAvatarUpdater(c.userRepo, c.validator)
	c.userID = int32(1)
	c.userUUID = "a_uuid"
	c.user = entities.User{ID: c.userID, UserUUID: c.userUUID}
}

func (c *AvatarUpdaterTests) TearDownTest() {
	c.userRepo.AssertExpectations(c.T())
	c.validator.AssertExpectations(c.T())
}

func (c *AvatarUpdaterTests) TestRun() {
	expectedAvatarURL := "avatar/url.jpg"
	req := user.AvatarUpdaterRequest{UserUUID: c.userUUID, AvatarURL: expectedAvatarURL}
	c.validator.On("Validate", &req).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.userRepo.On("FindUserByUUID", c.ctx, c.userUUID).Return(&c.user, nil)
	c.user.AvatarUrl = expectedAvatarURL
	c.userRepo.On("UpdateUser", c.ctx, &c.user).Return(nil)

	got, err := c.avatarUpdater.Run(c.ctx, &req)

	c.Nil(err)
	c.NotNil(got)
	c.Equal(expectedAvatarURL, got.UserUpdated.AvatarUrl)

}

func (c *AvatarUpdaterTests) TestRun_UserNotFound() {
	expectedAvatarURL := "avatar/url.jpg"
	req := user.AvatarUpdaterRequest{UserUUID: c.userUUID, AvatarURL: expectedAvatarURL}
	c.validator.On("Validate", &req).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.userRepo.On("FindUserByUUID", c.ctx, c.userUUID).Return(nil, repositories.NotFoundError{})

	got, err := c.avatarUpdater.Run(c.ctx, &req)

	c.NotNil(err)
	c.Nil(got)
	c.IsType(repositories.NotFoundError{}, err)

}

func (c *AvatarUpdaterTests) TestRun_ValidationError() {
	expectedAvatarURL := "avatar/url.jpg"
	req := user.AvatarUpdaterRequest{UserUUID: c.userUUID, AvatarURL: expectedAvatarURL}
	validationErr := apperrors.ValidationErrors{{Field: "user_uuid", Validation: "required"}}
	c.validator.On("Validate", &req).Return(validators.ValidateOutput{IsValid: false, Err: validationErr})

	got, err := c.avatarUpdater.Run(c.ctx, &req)

	c.NotNil(err)
	c.Nil(got)
	c.Equal(validationErr, err)

}

func (c *AvatarUpdaterTests) TestRun_UpdateError() {
	expectedAvatarURL := "avatar/url.jpg"
	req := user.AvatarUpdaterRequest{UserUUID: c.userUUID, AvatarURL: expectedAvatarURL}
	c.validator.On("Validate", &req).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.userRepo.On("FindUserByUUID", c.ctx, c.userUUID).Return(&c.user, nil)
	c.user.AvatarUrl = expectedAvatarURL
	c.userRepo.On("UpdateUser", c.ctx, &c.user).Return(fmt.Errorf("an ordinary error"))

	got, err := c.avatarUpdater.Run(c.ctx, &req)

	c.NotNil(err)
	c.Nil(got)

}
