package paginator

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/sort"
	"github.com/manicar2093/health_records/internal/config"
)

// PageSort
type PageSort struct {
	Page         float64  `validate:"required,gt=0" json:"page_number,omitempty"`
	ItemsPerPage float64  `json:"itemsPerPage,omitempty"`
	SortBy       []string `json:"sortBy,omitempty"`
	SortDesc     []bool   `json:"sortDesc,omitempty"`
	GroupBy      []string `json:"groupBy,omitempty"`
	GroupDesc    []bool   `json:"groupDesc,omitempty"`
	MustSort     bool     `json:"mustSort,omitempty"`
	MultiSort    bool     `json:"multiSort,omitempty"`
	filters      []rel.Querier
}

func CreatePageSortFromMap(values map[string]interface{}) *PageSort {
	var res PageSort

	b, e := json.Marshal(values)
	if e != nil {
		panic(e)
	}
	if e := json.Unmarshal(b, &res); e != nil {
		panic(e)
	}
	return &res
}

func (c *PageSort) SetFiltersQueries(queries ...rel.Querier) {
	c.filters = queries
}

func (c *PageSort) GetFiltersQueries() []rel.Querier {
	return c.filters
}

func (c *PageSort) GetSortQueries() []rel.Querier {
	var sortQueries []rel.Querier
	for i := 0; i < len(c.SortBy); i++ {
		var sortQuery rel.Querier
		isSortDesc := c.SortDesc[i]
		if isSortDesc {
			sortQuery = sort.Desc(c.SortBy[i])
		} else {
			sortQuery = sort.Asc(c.SortBy[i])
		}

		sortQueries = append(sortQueries, sortQuery)
	}
	return sortQueries
}

func (c *PageSort) GetItemsPerPage() int {
	if config.PageSize > 0 {
		return config.PageSize
	}
	return int(c.ItemsPerPage)
}
func (c *PageSort) GetPage() int {
	return int(c.Page)
}

type PageError struct {
	PageNumber int
}

func (c PageError) Error() string {
	return fmt.Sprintf("Page %v does not exists", c.PageNumber)
}

func (c PageError) StatusCode() int {
	return http.StatusBadRequest
}

type Paginator struct {
	TotalPages   int         `json:"total_pages,omitempty"`
	CurrentPage  int         `json:"current_page,omitempty"`
	PreviousPage int         `json:"previous_page,omitempty"`
	PageSize     int         `json:"page_size,omitempty"`
	NextPage     int         `json:"next_page,omitempty"`
	TotalEntries int         `json:"total_entries,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}
