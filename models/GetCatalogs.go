package models

type GetCatalogsRequest struct {
	CatalogNames []string `json:"catalog_names" validate:"gt=0"`
}
