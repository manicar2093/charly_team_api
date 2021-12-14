package models

type Paginator struct {
	TotalPages   int         `json:"total_pages,omitempty"`
	CurrentPage  int         `json:"current_page,omitempty"`
	PreviousPage int         `json:"previous_page,omitempty"`
	NextPage     int         `json:"next_page,omitempty"`
	TotalEntries int         `json:"total_entries,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}
