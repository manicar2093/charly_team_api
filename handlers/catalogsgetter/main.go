package cataloggetter

import (
	"context"

	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/validators"
)

type CatalogGetter interface {
	Run(ctx context.Context, catalogs *CatalogGetterRequest) (*CatalogGetterResponse, error)
}

type catalogGetteImpl struct {
	catalogsRepository repositories.CatalogRepository
	validator          validators.ValidatorService
}

func NewCatalogGetterImpl(
	catalogsRepository repositories.CatalogRepository,
	validator validators.ValidatorService,
) *catalogGetteImpl {
	return &catalogGetteImpl{
		catalogsRepository: catalogsRepository,
		validator:          validator,
	}
}

func (c *catalogGetteImpl) Run(ctx context.Context, catalogs *CatalogGetterRequest) (*CatalogGetterResponse, error) {
	validation := c.validator.Validate(catalogs)
	if !validation.IsValid {
		return nil, validation.Err
	}

	gotCatalogs, err := c.CatalogFactoryLoop(ctx, catalogs)

	if err != nil {
		return nil, err
	}

	return gotCatalogs, nil

}

// CatalogFactory check if catalog exists.
// If exists creates the need catalog response, otherwise return a NotCatalogFound
func (c *catalogGetteImpl) CatalogFactory(
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
func (c *catalogGetteImpl) CatalogFactoryLoop(
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
