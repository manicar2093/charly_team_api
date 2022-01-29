package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/charly_team_api/handlers/biotestfilters/biotestsbyuseruuidfinder"
	"github.com/manicar2093/charly_team_api/internal/models"
)

type GetAllBiotestByUserUUIDAWSLambda struct {
	biotestsByUserUUIDFinder biotestsbyuseruuidfinder.BiotestByUserUUID
}

func NewGetAllBiotestByUserUUIDAWSLambda(biotestsByUserUUIDFinder biotestsbyuseruuidfinder.BiotestByUserUUID) *GetAllBiotestByUserUUIDAWSLambda {
	return &GetAllBiotestByUserUUIDAWSLambda{biotestsByUserUUIDFinder: biotestsByUserUUIDFinder}
}

func (c *GetAllBiotestByUserUUIDAWSLambda) Handler(ctx context.Context, req biotestsbyuseruuidfinder.BiotestByUserUUIDRequest) (*models.Response, error) {
	res, err := c.biotestsByUserUUIDFinder.Run(ctx, &req)

	if err != nil {
		return models.CreateResponseFromError(err), nil
	}

	return models.CreateResponse(
		http.StatusOK,
		res.FoundBiotests,
	), nil
}
