package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/charly_team_api/internal/handlers/userfilters/userlikeemailornamefinder"
	"github.com/manicar2093/charly_team_api/internal/models"
)

type GetUsersLikeEmailOrNameAWSLambda struct {
	userLikeEmailOrNameFinder userlikeemailornamefinder.UserLikeEmailOrNameFinder
}

func NewGetUsersLikeEmailOrNameAWSLambda(userLikeEmailOrNameFinder userlikeemailornamefinder.UserLikeEmailOrNameFinder) *GetUsersLikeEmailOrNameAWSLambda {
	return &GetUsersLikeEmailOrNameAWSLambda{userLikeEmailOrNameFinder: userLikeEmailOrNameFinder}
}

func (c *GetUsersLikeEmailOrNameAWSLambda) Handler(ctx context.Context, req userlikeemailornamefinder.UserLikeEmailOrNameFinderRequest) (*models.Response, error) {
	res, err := c.userLikeEmailOrNameFinder.Run(ctx, &req)

	if err != nil {
		return models.CreateResponseFromError(err), nil
	}

	return models.CreateResponse(
		http.StatusOK,
		res.FetchedData,
	), nil
}
