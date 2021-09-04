package filters

import (
	"context"
	"fmt"

	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/paginator"
)

type FilterFunc func(filterParameters *FilterParameters) (interface{}, error)

type FilterParameters struct {
	Ctx       context.Context
	Repo      rel.Repository
	Values    interface{}
	Paginator paginator.Paginable
}

type FilterRunner struct {
	FilterName string
	Filter     FilterFunc
	Found      bool
}

func (c FilterRunner) Run(filterParameters *FilterParameters) (interface{}, error) {
	if !c.IsFound() {
		return nil, apperrors.BadStatusError{
			Message: fmt.Sprintf("'%s' filter does not exists",
				c.FilterName,
			),
		}
	}
	return c.Filter(filterParameters)
}

func (c FilterRunner) IsFound() bool {
	return c.Found
}
