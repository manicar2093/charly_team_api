package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/charly_team_api/internal/handlers/userfilters"
	"github.com/manicar2093/charly_team_api/pkg/models"
)

type GetUserByUUIDAWSLambda struct {
	userByUUIDFinder userfilters.UserByUUIDFinder
}

func NewGetUserByUUIDAWSLambda(
	userByUUIDFinder userfilters.UserByUUIDFinder,
) *GetUserByUUIDAWSLambda {
	return &GetUserByUUIDAWSLambda{userByUUIDFinder: userByUUIDFinder}
}

func (c *GetUserByUUIDAWSLambda) Handler(
	ctx context.Context,
	req userfilters.UserByUUIDFinderRequest,
) (*models.Response, error) {
	res, err := c.userByUUIDFinder.Run(ctx, &req)
	if err != nil {
		return models.CreateResponseFromError(err), nil
	}

	return models.CreateResponse(
		http.StatusOK,
		res.UserFound,
	), nil
}
