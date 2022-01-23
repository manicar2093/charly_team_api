package cataloggetter

import "github.com/manicar2093/charly_team_api/db/entities"

var registeredCatalogs = map[string]interface{}{
	"biotype":               &[]entities.Biotype{},
	"bone_density":          &[]entities.BoneDensity{},
	"heart_healths":         &[]entities.HeartHealth{},
	"roles":                 &[]entities.Role{},
	"weight_clasifications": &[]entities.WeightClasification{},
	"genders":               &[]entities.Gender{},
}
