package main

import (
	"context"
	"testing"

	"github.com/manicar2093/health_records/internal/db/entities"
	"github.com/manicar2093/health_records/internal/user"
	"github.com/manicar2093/health_records/mocks"
	"github.com/manicar2093/health_records/pkg/apperrors"
	"github.com/manicar2093/health_records/pkg/models"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(UpdateUserAWSLambdaTests))
}

type UpdateUserAWSLambdaTests struct {
	suite.Suite
	ctx                 context.Context
	userAvatarUpdater   *mocks.AvatarUpdater
	updateUserAWSLambda *UpdateUserAWSLambda
}

func (c *UpdateUserAWSLambdaTests) SetupTest() {
	c.ctx = context.Background()
	c.userAvatarUpdater = &mocks.AvatarUpdater{}
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
