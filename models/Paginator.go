package models

type Paginator struct {
	TotalPages   int         `json:"total_pages,omitempty"`
	CurrendPage  int         `json:"currend_page,omitempty"`
	PreviousPage int         `json:"previous_page,omitempty"`
	NextPage     int         `json:"next_page,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}
