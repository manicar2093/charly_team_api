package repositories

import (
	"testing"

	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/stretchr/testify/assert"
)

type NotAnEntity struct{}

func TestFindAllCatalogItems(t *testing.T) {

	var data []entities.Biotype
	repository := NewCatalogRepositoryImpl(DB)

	err := repository.FindAllCatalogItems(Ctx, &data)

	assert.Nil(t, err, "should not present error")

	assert.Greater(t, len(data), 1, "no items in data list")

}

func TestFindAllCatalogItemsError(t *testing.T) {

	repository := NewCatalogRepositoryImpl(DB)

	var wrongData []NotAnEntity

	err := repository.FindAllCatalogItems(Ctx, &wrongData)

	assert.NotNil(t, err, "should present error")

}
