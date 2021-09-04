package filters

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/db/paginator"
)

type FilterParameters struct {
	Ctx       context.Context
	Repo      rel.Repository
	Values    interface{}
	Paginator paginator.Paginable
}

type FilterFunc func(filterParameters *FilterParameters) (interface{}, error)

type FilterService interface {
	// GetUserFilter looks up if the requested filter exists. If exists
	// Run method will be
	GetUserFilter(string) FilterRunner
}

type FilterRunner struct {
	Filter  FilterFunc
	IsFound bool
}

func (c FilterRunner) Run(filterParameters *FilterParameters) (interface{}, error) {
	return c.Filter(filterParameters)
}
