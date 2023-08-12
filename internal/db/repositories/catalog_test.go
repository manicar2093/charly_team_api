package repositories_test

import (
	"testing"

	"github.com/manicar2093/health_records/internal/db/entities"
	"github.com/manicar2093/health_records/internal/db/repositories"
	"github.com/stretchr/testify/assert"
)

type NotAnEntity struct{}

func TestFindAllCatalogItems(t *testing.T) {

	repository := repositories.NewCatalogRepositoryImpl(DB)

	data, err := repository.FindAllCatalogItems(Ctx, &[]entities.Biotype{})

	assert.Nil(t, err, "should not present error")
	dataList, ok := data.(*[]entities.Biotype)
	assert.True(t, ok, "error parsing data to list")

	assert.Greater(t, len(*dataList), 1, "no items in data list")

}

func TestFindAllCatalogItemsError(t *testing.T) {

	repository := repositories.NewCatalogRepositoryImpl(DB)

	_, err := repository.FindAllCatalogItems(Ctx, &[]NotAnEntity{})

	assert.NotNil(t, err, "should present error")

}
