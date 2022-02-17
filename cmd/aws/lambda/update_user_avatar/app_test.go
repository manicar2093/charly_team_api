package main

import (
	"context"
	"testing"

	"github.com/manicar2093/charly_team_api/internal/db/entities"
	"github.com/manicar2093/charly_team_api/internal/handlers/user"
	"github.com/manicar2093/charly_team_api/pkg/apperrors"
	"github.com/manicar2093/charly_team_api/pkg/models"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(UpdateUserAWSLambdaTests))
}

type UpdateUserAWSLambdaTests struct {
	suite.Suite
	ctx                 context.Context
	userAvatarUpdater   *user.MockAvatarUpdater
	updateUserAWSLambda *UpdateUserAWSLambda
}

func (c *UpdateUserAWSLambdaTests) SetupTest() {
	c.ctx = context.Background()
	c.userAvatarUpdater = &user.MockAvatarUpdater{}
	c.updateUserAWSLambda = NewUpdateUserAWSLambda(c.userAvatarUpdater)
}

func (c *UpdateUserAWSLambdaTests) TearDownTest() {
	c.userAvatarUpdater.AssertExpectations(c.T())
}

func (c *UpdateUserAWSLambdaTests) TestHandler() {
	req := user.AvatarUpdaterRequest{UserUUID: "a_uuid", AvatarURL: "avatar/url"}
	userRes := &entities.User{ID: int32(1)}
	res := user.AvatarUpdaterResponse{UserUpdated: userRes}
	c.userAvatarUpdater.On("Run", c.ctx, &req).Return(&res, nil)

	got, err := c.updateUserAWSLambda.Handler(c.ctx, req)

	c.Nil(err)
	c.NotNil(got)
	c.IsType(&models.Response{}, got)

}

func (c *UpdateUserAWSLambdaTests) TestHandler_ValidationError() {
	req := user.AvatarUpdaterRequest{}
	validationErr := apperrors.ValidationErrors{{Field: "user_uuid", Validation: "required"}, {Field: "avatar_url", Validation: "required"}}
	c.userAvatarUpdater.On("Run", c.ctx, &req).Return(nil, validationErr)

	got, err := c.updateUserAWSLambda.Handler(c.ctx, req)

	c.Nil(err)
	c.NotNil(got)
	c.IsType(&models.Response{}, got)

}
