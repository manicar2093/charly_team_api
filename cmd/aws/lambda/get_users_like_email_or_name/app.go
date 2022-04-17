package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/charly_team_api/internal/userfilters"
	"github.com/manicar2093/charly_team_api/pkg/models"
)

type GetUsersLikeEmailOrNameAWSLambda struct {
	userLikeEmailOrNameFinder userfilters.UserLikeEmailOrNameFinder
}

func NewGetUsersLikeEmailOrNameAWSLambda(userLikeEmailOrNameFinder userfilters.UserLikeEmailOrNameFinder) *GetUsersLikeEmailOrNameAWSLambda {
	return &GetUsersLikeEmailOrNameAWSLambda{userLikeEmailOrNameFinder: userLikeEmailOrNameFinder}
}

func (c *GetUsersLikeEmailOrNameAWSLambda) Handler(ctx context.Context, req userfilters.UserLikeEmailOrNameFinderRequest) (*models.Response, error) {
	res, err := c.userLikeEmailOrNameFinder.Run(ctx, &req)

	if err != nil {
		return models.CreateResponseFromError(err), nil
	}

	return models.CreateResponse(
		http.StatusOK,
		res.FetchedData,
	), nil
}
