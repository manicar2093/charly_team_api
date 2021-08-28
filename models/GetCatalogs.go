package models

type GetCatalogsRequest struct {
	CatalogNames []string `json:"catalog_names" validator:"required, gt=0, dive, required"`
}

type GetCatalogsResponse struct {
	Error   string        `json:"error,omitempty"`
	Catalog []interface{} `json:"catalog,omitempty"`
}
