package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/charly_team_api/internal/db/entities"
	"github.com/manicar2093/charly_team_api/internal/user"
	"github.com/manicar2093/charly_team_api/pkg/models"
)

type UpdateUserAWSLambda struct {
	userUpdater user.UserUpdater
}

func NewUpdateUserAWSLambda(userUpdater user.UserUpdater) *UpdateUserAWSLambda {
	return &UpdateUserAWSLambda{userUpdater}
}

func (c *UpdateUserAWSLambda) Handler(ctx context.Context, req entities.User) (*models.Response, error) {
	res, err := c.userUpdater.Run(ctx, &req)

	if err != nil {
		return models.CreateResponseFromError(err), nil
	}

	return models.CreateResponse(
			http.StatusOK,
			res.UserUpdated),
		nil
}
