package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/handlers/userupdater"
	"github.com/manicar2093/charly_team_api/internal/models"
)

type UpdateUserAWSLambda struct {
	userUpdater userupdater.UserUpdater
}

func NewUpdateUserAWSLambda(userUpdater userupdater.UserUpdater) *UpdateUserAWSLambda {
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
