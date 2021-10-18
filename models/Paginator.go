package models

type Paginator struct {
	TotalPages   int         `json:"total_pages"`
	CurrentPage  int         `json:"current_page"`
	PreviousPage int         `json:"previous_page"`
	NextPage     int         `json:"next_page"`
	Data         interface{} `json:"data"`
}
