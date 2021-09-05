package apperrors

import (
	"net/http"
)

type UserNotFound struct{}

func (c UserNotFound) Error() string {
	return "user does not exists"
}

func (c UserNotFound) StatusCode() int {
	return http.StatusNotFound
}
