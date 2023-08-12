package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/health_records/internal/biotestfilters"
	"github.com/manicar2093/health_records/pkg/models"
)

type GetBiotestComparitionDataAWSLambda struct {
	biotestComparitionDataFinder biotestfilters.BiotestComparitionDataFinder
}

func NewGetBiotestComparitionDataAWSLambda(
	biotestComparitionDataFinder biotestfilters.BiotestComparitionDataFinder,
) *GetBiotestComparitionDataAWSLambda {
	return &GetBiotestComparitionDataAWSLambda{
		biotestComparitionDataFinder: biotestComparitionDataFinder,
	}
}

func (c *GetBiotestComparitionDataAWSLambda) Handler(
	ctx context.Context,
	req biotestfilters.BiotestComparitionDataFinderRequest,
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
