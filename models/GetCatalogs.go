package models

type GetCatalogsRequest struct {
	CatalogNames []string `json:"catalog_names" validator:"required, gt=0, dive, required"`
}
