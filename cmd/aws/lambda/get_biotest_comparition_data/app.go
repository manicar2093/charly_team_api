package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/charly_team_api/internal/handlers/biotestfilters/biotestcomparitiondatafinder"
	"github.com/manicar2093/charly_team_api/pkg/models"
)

type GetBiotestComparitionDataAWSLambda struct {
	biotestComparitionDataFinder biotestcomparitiondatafinder.BiotestComparitionDataFinder
}

func NewGetBiotestComparitionDataAWSLambda(
	biotestComparitionDataFinder biotestcomparitiondatafinder.BiotestComparitionDataFinder,
) *GetBiotestComparitionDataAWSLambda {
	return &GetBiotestComparitionDataAWSLambda{
		biotestComparitionDataFinder: biotestComparitionDataFinder,
	}
}

func (c *GetBiotestComparitionDataAWSLambda) Handler(
	ctx context.Context,
	req biotestcomparitiondatafinder.BiotestComparitionDataFinderRequest,
) (*models.Response, error) {
	res, err := c.biotestComparitionDataFinder.Run(ctx, &req)

	if err != nil {
		return models.CreateResponseFromError(err), nil
	}

	return models.CreateResponse(
		http.StatusOK,
		res.ComparitionData,
	), nil
}
