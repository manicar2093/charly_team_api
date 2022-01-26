package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/charly_team_api/handlers/userfilters/userbyuuidfinder"
	"github.com/manicar2093/charly_team_api/models"
)

type GetUserByUUIDAWSLambda struct {
	userByUUIDFinder userbyuuidfinder.UserByUUIDFinder
}

func NewGetUserByUUIDAWSLambda(
	userByUUIDFinder userbyuuidfinder.UserByUUIDFinder,
) *GetUserByUUIDAWSLambda {
	return &GetUserByUUIDAWSLambda{userByUUIDFinder: userByUUIDFinder}
}

func (c *GetUserByUUIDAWSLambda) Handler(
	ctx context.Context,
	req userbyuuidfinder.UserByUUIDFinderRequest,
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
