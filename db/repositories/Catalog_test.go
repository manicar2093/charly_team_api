package repositories

import (
	"testing"

	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/stretchr/testify/assert"
)

type NotAnEntity struct{}

func TestFindAllCatalogItems(t *testing.T) {

	repository := NewCatalogRepositoryGorm(DB)

	data, err := repository.FindAllCatalogItems(&[]entities.Biotype{})

	assert.Nil(t, err, "should not present error")

	dataList, ok := data.(*[]entities.Biotype)
	assert.True(t, ok, "error parsing data to list")

	assert.Greater(t, len(*dataList), 1, "no items in data list")

}

func TestFindAllCatalogItemsError(t *testing.T) {

	repository := NewCatalogRepositoryGorm(DB)

	_, err := repository.FindAllCatalogItems(&[]NotAnEntity{})

	assert.NotNil(t, err, "should present error")

}
