package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/handlers/biotestcreator"
	"github.com/manicar2093/charly_team_api/models"
)

type CreateBiotestAWSLambda struct {
	biotestCreator biotestcreator.BiotestCreator
}

func NewCreateBiotestAWSLambda(
	biotestCreator biotestcreator.BiotestCreator,
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
