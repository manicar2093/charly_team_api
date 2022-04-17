package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/charly_team_api/internal/biotest"
	"github.com/manicar2093/charly_team_api/pkg/models"
)

type SaveBiotestImagesAWSLambda struct {
	biotestImageSaver biotest.BiotestImagesSaver
}

func NewSaveBiotestImagesAWSLambda(biotestImageSaver biotest.BiotestImagesSaver) *SaveBiotestImagesAWSLambda {
	return &SaveBiotestImagesAWSLambda{biotestImageSaver: biotestImageSaver}
}

func (c *SaveBiotestImagesAWSLambda) Handler(ctx context.Context, biotestImages biotest.BiotestImagesSaverRequest) (*models.Response, error) {
	_, err := c.biotestImageSaver.Run(ctx, &biotestImages)

	if err != nil {
		return models.CreateResponseFromError(err), nil
	}

	return models.CreateResponse(
		http.StatusOK,
		nil,
	), nil
}
