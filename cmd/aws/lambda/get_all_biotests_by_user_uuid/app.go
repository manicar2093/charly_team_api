package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/charly_team_api/internal/handlers/biotestfilters"
	"github.com/manicar2093/charly_team_api/pkg/models"
)

type GetAllBiotestByUserUUIDAWSLambda struct {
	biotestsByUserUUIDFinder biotestfilters.BiotestByUserUUID
}

func NewGetAllBiotestByUserUUIDAWSLambda(biotestsByUserUUIDFinder biotestfilters.BiotestByUserUUID) *GetAllBiotestByUserUUIDAWSLambda {
	return &GetAllBiotestByUserUUIDAWSLambda{biotestsByUserUUIDFinder: biotestsByUserUUIDFinder}
}

func (c *GetAllBiotestByUserUUIDAWSLambda) Handler(ctx context.Context, req biotestfilters.BiotestByUserUUIDRequest) (*models.Response, error) {
	res, err := c.biotestsByUserUUIDFinder.Run(ctx, &req)

	if err != nil {
		return models.CreateResponseFromError(err), nil
	}

	return models.CreateResponse(
		http.StatusOK,
		res.FoundBiotests,
	), nil
}
