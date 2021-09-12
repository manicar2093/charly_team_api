package filters

import (
	"context"
	"fmt"

	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/paginator"
	"github.com/manicar2093/charly_team_api/validators"
)

type FilterFunc func(filterParameters *FilterParameters) (interface{}, error)

type FilterParameters struct {
	Ctx       context.Context
	Repo      rel.Repository
	Paginator paginator.Paginable
	Validator validators.ValidatorService
	// Values can be anything but must be handled in the filter to get
	// data from it
	Values interface{}
}

type FilterRegistrationData struct {
	Name string
	Func FilterFunc
}

// Filter is an implementation of Filterable
type Filter struct {
	registeredFilters map[string]FilterFunc
	filterToRun       FilterFunc
	filterParams      FilterParameters
}

func NewFilter(params *FilterParameters, filters ...FilterRegistrationData) Filterable {
	registered := make(map[string]FilterFunc)
	for _, item := range filters {
		registered[item.Name] = item.Func
	}
	return &Filter{registeredFilters: registered, filterParams: *params}
}

// GetUserFilter looks up if the requested filter exists. If exists
// Run method will be
func (c *Filter) GetFilter(filterName string) error {
	filterFound, isFound := c.registeredFilters[filterName]
	if !isFound {
		return apperrors.BadStatusError{
			Message: fmt.Sprintf("'%s' filter does not exists",
				filterName,
			),
		}
	}
	c.filterToRun = filterFound
	return nil
}

func (c *Filter) Run() (interface{}, error) {
	return c.filterToRun(&c.filterParams)
}
