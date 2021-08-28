package apperrors

import (
	"fmt"
	"net/http"
)

type NoCatalogFound struct {
	CatalogName string
}

func (n NoCatalogFound) Error() string {
	return fmt.Sprintf("catalog '%s' does not exists", n.CatalogName)
}

func (n NoCatalogFound) StatusCode() int {
	return http.StatusNotFound
}
