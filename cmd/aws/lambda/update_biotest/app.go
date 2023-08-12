package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/health_records/internal/biotest"
	"github.com/manicar2093/health_records/internal/db/entities"
	"github.com/manicar2093/health_records/pkg/models"
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
