package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/charly_team_api/internal/handlers/catalog"
	"github.com/manicar2093/charly_team_api/pkg/models"
)

type GetCatalogsAWSLambda struct {
	catalogsGetter catalog.CatalogGetter
}

func NewGetCatalogsAWSLambda(catalogsGetter catalog.CatalogGetter) *GetCatalogsAWSLambda {
	return &GetCatalogsAWSLambda{catalogsGetter: catalogsGetter}
}

func (c *GetCatalogsAWSLambda) Handler(ctx context.Context, catalogs catalog.CatalogGetterRequest) (*models.Response, error) {
	res, err := c.catalogsGetter.Run(ctx, &catalogs)

	if err != nil {
		return models.CreateResponseFromError(err), nil
	}

	return models.CreateResponse(
		http.StatusOK,
		res.Catalogs,
	), nil
}
