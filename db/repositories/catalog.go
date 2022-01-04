package repositories

import (
	"context"

	"github.com/go-rel/rel"
)

type CatalogRepository interface {
	// FindAllCatalogItems find any catalog from DB.
	FindAllCatalogItems(context.Context, interface{}) (interface{}, error)
}

type CatalogRepositoryImpl struct {
	provider rel.Repository
}

func NewCatalogRepositoryImpl(db rel.Repository) CatalogRepository {
	return &CatalogRepositoryImpl{db}
}

// FindCatalogByName find any catalog from DB
func (c CatalogRepositoryImpl) FindAllCatalogItems(ctx context.Context, holderEntity interface{}) (interface{}, error) {
	err := c.provider.FindAll(ctx, holderEntity)
	if err != nil {
		return holderEntity, err
	}
	return holderEntity, nil
}
