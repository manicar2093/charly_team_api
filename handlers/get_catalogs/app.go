package main

import (
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/connections"
	"github.com/manicar2093/charly_team_api/entities"
)

func CatalogFactory(
	catalog string,
	db connections.Findable,
) (interface{}, error) {

	switch catalog {
	case "biotype":
		var data []entities.Biotype
		return FindCatalogByName(data, db)
	case "bone_density":
		var data []entities.BoneDensity
		return FindCatalogByName(data, db)
	case "heart_healths":
		var data []entities.HeartHealth
		return FindCatalogByName(data, db)
	case "roles":
		var data []entities.Role
		return FindCatalogByName(data, db)
	case "weight_classifications":
		var data []entities.WeightClasification
		return FindCatalogByName(data, db)
	default:
		return []interface{}{}, apperrors.NoCatalogFound{CatalogName: catalog}

	}
}

func FindCatalogByName(holderEntity interface{}, db connections.Findable) (interface{}, error) {
	dataSlice := db.Find(holderEntity)
	if dataSlice.Error != nil {
		return holderEntity, dataSlice.Error
	}
	return holderEntity, nil
}
