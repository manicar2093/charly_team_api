package models

type Paginator struct {
	TotalPages int         `json:"total_pages,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
