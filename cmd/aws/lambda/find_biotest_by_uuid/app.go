package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/charly_team_api/internal/handlers/biotestfilters/biotestbyuuidfinder"
	"github.com/manicar2093/charly_team_api/pkg/models"
)

type FindBiotestByUUIDAWSLambda struct {
	biotestByUUIDFinder biotestbyuuidfinder.BiotestByUUID
}

func NewFindBiotestByUUIDAWSLambda(biotestByUUIDFinder biotestbyuuidfinder.BiotestByUUID) *FindBiotestByUUIDAWSLambda {
	return &FindBiotestByUUIDAWSLambda{biotestByUUIDFinder: biotestByUUIDFinder}
}

func (c *FindBiotestByUUIDAWSLambda) Handler(ctx context.Context, req biotestbyuuidfinder.BiotestByUUIDRequest) (*models.Response, error) {
	res, err := c.biotestByUUIDFinder.Run(ctx, &req)

	if err != nil {
		return models.CreateResponseFromError(err), nil
	}

	return models.CreateResponse(
		http.StatusOK,
		res.Biotest,
	), nil
}
