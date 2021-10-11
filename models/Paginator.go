package models

type Paginator struct {
	TotalPages   int         `json:"total_pages"`
	CurrendPage  int         `json:"currend_page"`
	PreviousPage int         `json:"previous_page"`
	NextPage     int         `json:"next_page"`
	Data         interface{} `json:"data"`
}
