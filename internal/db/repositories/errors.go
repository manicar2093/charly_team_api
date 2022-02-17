package repositories

import (
	"fmt"
	"net/http"
)

type NotFoundError struct {
	Entity     string
	Identifier interface{}
}

func (c NotFoundError) Error() string {
	return fmt.Sprintf("%s with identifier %s not found", c.Entity, c.Identifier)
}

func (c NotFoundError) StatusCode() int {
	return http.StatusNotFound
}
