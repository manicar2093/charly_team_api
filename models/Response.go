package models

import "net/http"

type Response struct {
	Code   int         `json:"code,omitempty"`
	Status string      `json:"status,omitempty"`
	Body   interface{} `json:"body,omitempty"`
}

// CreateResponse takes all inputs and create a Response.
// Please consider use http status constants to httpStatus
func CreateResponse(httpStatus int, body interface{}) *Response {
	return &Response{
		Code:   httpStatus,
		Status: http.StatusText(httpStatus),
		Body:   body,
	}
}
