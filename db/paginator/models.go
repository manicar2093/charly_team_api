package paginator

import (
	"fmt"
	"net/http"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/sort"
	"github.com/manicar2093/charly_team_api/config"
)

// PageSort
type PageSort struct {
	Page         float64  `json:"page,omitempty"`
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
	page, _ := values["page_number"].(float64)
	items, _ := values["itemsPerPage"].(float64)
	sortBy, _ := values["sortBy"].([]string)
	sortDesc, _ := values["sortDesc"].([]bool)
	groupBy, _ := values["groupBy"].([]string)
	groupDesc, _ := values["groupDesc"].([]bool)
	mustSort, _ := values["mustSort"].(bool)
	multiSort, _ := values["multiSort"].(bool)

	return &PageSort{
		Page:         page,
		ItemsPerPage: items,
		SortBy:       sortBy,
		SortDesc:     sortDesc,
		GroupBy:      groupBy,
		GroupDesc:    groupDesc,
		MustSort:     mustSort,
		MultiSort:    multiSort,
	}
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
