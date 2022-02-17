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

type BadRequestError struct {
	Message string
}

func (c BadRequestError) Error() string {
	return c.Message
}

func (c BadRequestError) StatusCode() int {
	return http.StatusBadRequest
}
