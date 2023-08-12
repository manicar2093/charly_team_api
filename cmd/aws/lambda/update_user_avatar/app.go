package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/health_records/internal/user"
	"github.com/manicar2093/health_records/pkg/models"
)

type UpdateUserAWSLambda struct {
	userAvatarUpdater user.AvatarUpdater
}

func NewUpdateUserAWSLambda(userAvatarUpdater user.AvatarUpdater) *UpdateUserAWSLambda {
	return &UpdateUserAWSLambda{userAvatarUpdater: userAvatarUpdater}
}

func (c *UpdateUserAWSLambda) Handler(ctx context.Context, req user.AvatarUpdaterRequest) (*models.Response, error) {
	res, err := c.userAvatarUpdater.Run(ctx, &req)
	if err != nil {
		return models.CreateResponseFromError(err), nil
	}

	return models.CreateResponse(
		http.StatusOK,
		res.UserUpdated,
	), nil
}
