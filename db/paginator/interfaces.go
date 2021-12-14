package paginator

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/models"
)

type Paginable interface {
	// Deprecated: CreatePaginator is deprecated. Use CreatePagination instead
	CreatePaginator(
		ctx context.Context,
		tableName string,
		holder interface{},
		pageNumber int,
		queries ...rel.Querier,
	) (*models.Paginator, error)
	CreatePagination(
		ctx context.Context,
		tableName string,
		holder interface{},
		pageSort *PageSort,
	) (*models.Paginator, error)
}
