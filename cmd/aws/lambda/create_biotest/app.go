package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/health_records/internal/biotest"
	"github.com/manicar2093/health_records/internal/db/entities"
	"github.com/manicar2093/health_records/pkg/models"
)

type CreateBiotestAWSLambda struct {
	biotestCreator biotest.BiotestCreator
}

func NewCreateBiotestAWSLambda(
	biotestCreator biotest.BiotestCreator,
) *CreateBiotestAWSLambda {
	return &CreateBiotestAWSLambda{
		biotestCreator: biotestCreator,
	}
}

func (c *CreateBiotestAWSLambda) Handler(
	ctx context.Context,
	biotest entities.Biotest,
) (*models.Response, error) {
	res, err := c.biotestCreator.Run(ctx, &biotest)

	if err != nil {
		return models.CreateResponseFromError(err), nil
	}

	return models.CreateResponse(
		http.StatusCreated,
		&CreateBiotestResponse{
			BiotestID:   res.BiotestCreated.ID,
			BiotestUUID: res.BiotestCreated.BiotestUUID,
		},
	), nil
}
