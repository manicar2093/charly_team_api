package models

type Response struct {
	Code   int         `json:"code,omitempty"`
	Status string      `json:"status,omitempty"`
	Body   interface{} `json:"body,omitempty"`
}
