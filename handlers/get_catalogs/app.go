package main

import (
	"context"

	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/models"
)

var catalogs = map[string]interface{}{
	"biotype":               &[]entities.Biotype{},
	"bone_density":          &[]entities.BoneDensity{},
	"heart_healths":         &[]entities.HeartHealth{},
	"roles":                 &[]entities.Role{},
	"weight_clasifications": &[]entities.WeightClasification{},
	"genders":               &[]entities.Gender{},
}

// CatalogFactory creates the need catalog response
func CatalogFactory(
	catalog string,
	catalogsRepository repositories.CatalogRepository,
	ctx context.Context,
) (interface{}, error) {

	handler, isRegistred := catalogs[catalog]
	if !isRegistred {
		return []interface{}{}, apperrors.NoCatalogFound{CatalogName: catalog}
	}

	return catalogsRepository.FindAllCatalogItems(ctx, handler)

}

// CatalogFactoryLoop creates all catalog response from a slice of requested catalogs
func CatalogFactoryLoop(
	catalogs models.GetCatalogsRequest,
	catalogsRepository repositories.CatalogRepository,
	ctx context.Context,
) (map[string]interface{}, error) {
	gotCatalogs := make(map[string]interface{})

	for _, catalog := range catalogs.CatalogNames {
		foundCatalog, err := CatalogFactory(catalog, catalogsRepository, ctx)
		if err != nil {
			return gotCatalogs, err
		}
		gotCatalogs[catalog] = foundCatalog
	}

	return gotCatalogs, nil
}
