package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/charly_team_api/internal/biotest"
	"github.com/manicar2093/charly_team_api/internal/db/entities"
	"github.com/manicar2093/charly_team_api/pkg/models"
)

type UpdateBiotestAWSLambda struct {
	biotestUpdater biotest.BiotestUpdater
}

func NewUpdateBiotestAWSLambda(biotestUpdater biotest.BiotestUpdater) *UpdateBiotestAWSLambda {
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
