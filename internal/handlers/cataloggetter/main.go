package cataloggetter

import (
	"context"

	"github.com/manicar2093/charly_team_api/internal/db/repositories"
	"github.com/manicar2093/charly_team_api/pkg/logger"
	"github.com/manicar2093/charly_team_api/pkg/validators"
)

type CatalogGetter interface {
	Run(ctx context.Context, catalogs *CatalogGetterRequest) (*CatalogGetterResponse, error)
}

type catalogGetterImpl struct {
	catalogsRepository repositories.CatalogRepository
	validator          validators.ValidatorService
}

func NewCatalogGetterImpl(
	catalogsRepository repositories.CatalogRepository,
	validator validators.ValidatorService,
) *catalogGetterImpl {
	return &catalogGetterImpl{
		catalogsRepository: catalogsRepository,
		validator:          validator,
	}
}

func (c *catalogGetterImpl) Run(ctx context.Context, catalogs *CatalogGetterRequest) (*CatalogGetterResponse, error) {
	logger.Info(catalogs)
	validation := c.validator.Validate(catalogs)
	if !validation.IsValid {
		logger.Error(validation.Err)
		return nil, validation.Err
	}

	gotCatalogs, err := c.CatalogFactoryLoop(ctx, catalogs)

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return gotCatalogs, nil

}

// CatalogFactory check if catalog exists.
// If exists creates the need catalog response, otherwise return a NotCatalogFound
func (c *catalogGetterImpl) CatalogFactory(
	ctx context.Context,
	catalog string,
) (interface{}, error) {

	handler, isRegistred := registeredCatalogs[catalog]
	if !isRegistred {
		return []interface{}{}, NoCatalogFound{CatalogName: catalog}
	}

	return c.catalogsRepository.FindAllCatalogItems(ctx, handler)

}

// CatalogFactoryLoop creates all catalog response from a slice of requested catalogs
func (c *catalogGetterImpl) CatalogFactoryLoop(
	ctx context.Context,
	catalogs *CatalogGetterRequest,
) (*CatalogGetterResponse, error) {

	gotCatalogs := make(map[string]interface{})

	for _, catalog := range catalogs.CatalogNames {
		foundCatalog, err := c.CatalogFactory(ctx, catalog)
		if err != nil {
			return nil, err
		}
		gotCatalogs[catalog] = foundCatalog
	}

	return &CatalogGetterResponse{Catalogs: gotCatalogs}, nil
}
