package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/handlers/biotestupdater"
	"github.com/manicar2093/charly_team_api/models"
)

type UpdateBiotestAWSLambda struct {
	biotestUpdater biotestupdater.BiotestUpdater
}

func NewUpdateBiotestAWSLambda(biotestUpdater biotestupdater.BiotestUpdater) *UpdateBiotestAWSLambda {
	return &UpdateBiotestAWSLambda{biotestUpdater: biotestUpdater}
}

func (c *UpdateBiotestAWSLambda) Handler(ctx context.Context, biotest entities.Biotest) (*models.Response, error) {
	res, err := c.biotestUpdater.Run(ctx, &biotest)

	if err != nil {
		return models.CreateResponseFromError(err), nil
	}

	return models.CreateResponse(
		http.StatusOK,
		res.BiotestUpdated,
	), nil
}
