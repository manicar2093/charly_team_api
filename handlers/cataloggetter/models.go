package cataloggetter

type CatalogGetterRequest struct {
	CatalogNames []string `json:"catalog_names" validate:"gt=0"`
}

type CatalogGetterResponse struct {
	Catalogs map[string]interface{}
}
