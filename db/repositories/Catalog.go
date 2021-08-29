package repositories

import "gorm.io/gorm"

type CatalogRepository interface {
	// FindAllCatalogItems find any catalog from DB.
	FindAllCatalogItems(holderEntity interface{}) (interface{}, error)
}

type CatalogRepositoryGorm struct {
	provider *gorm.DB
}

func NewCatalogRepositoryGorm(db *gorm.DB) CatalogRepository {
	return &CatalogRepositoryGorm{db}
}

// FindCatalogByName find any catalog from DB
func (c CatalogRepositoryGorm) FindAllCatalogItems(holderEntity interface{}) (interface{}, error) {
	dataSlice := c.provider.Find(holderEntity)
	if dataSlice.Error != nil {
		return holderEntity, dataSlice.Error
	}
	return holderEntity, nil
}
