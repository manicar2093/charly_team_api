package apperrors

import "net/http"

type NotFoundError struct {
	Message string
}

func (c NotFoundError) Error() string {
	return c.Message
}

func (c NotFoundError) StatusCode() int {
	return http.StatusNotFound
}
