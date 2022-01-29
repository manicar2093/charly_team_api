package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/charly_team_api/internal/handlers/cataloggetter"
	"github.com/manicar2093/charly_team_api/internal/models"
)

type GetCatalogsAWSLambda struct {
	catalogsGetter cataloggetter.CatalogGetter
}

func NewGetCatalogsAWSLambda(catalogsGetter cataloggetter.CatalogGetter) *GetCatalogsAWSLambda {
	return &GetCatalogsAWSLambda{catalogsGetter: catalogsGetter}
}

func (c *GetCatalogsAWSLambda) Handler(ctx context.Context, catalogs cataloggetter.CatalogGetterRequest) (*models.Response, error) {
	res, err := c.catalogsGetter.Run(ctx, &catalogs)

	if err != nil {
		return models.CreateResponseFromError(err), nil
	}

	return models.CreateResponse(
		http.StatusOK,
		res.Catalogs,
	), nil
}
