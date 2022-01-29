package models

import (
	"net/http"

	"github.com/manicar2093/charly_team_api/internal/apperrors"
)

const ValidationErrorMessage = "Request body does not satisfy needs. Please check documentation"

type Response struct {
	StatusCode int         `json:"code,omitempty"`
	Status     string      `json:"status,omitempty"`
	Body       interface{} `json:"body,omitempty"`
}

// CreateResponse takes all inputs and create a Response.
// Please consider use http status constants to httpStatus
func CreateResponse(httpStatus int, body interface{}) *Response {
	return &Response{
		StatusCode: httpStatus,
		Status:     http.StatusText(httpStatus),
		Body:       body,
	}
}

// CreateResponseFromError validates if error implements HandableErrors interfaz
// func to build error. If does not returns a InternalServerError http status
func CreateResponseFromError(err error) *Response {
	var (
		statusCode   int         = http.StatusInternalServerError
		errorMessage interface{} = err.Error()
	)

	if handledErr, ok := err.(apperrors.HandableErrors); ok {
		statusCode = handledErr.StatusCode()
	}

	if handledError, ok := err.(apperrors.ValidationErrors); ok {
		errorMessage = handledError
	}

	return CreateResponse(
		statusCode,
		ErrorReponse{
			Error: errorMessage,
		},
	)
}
