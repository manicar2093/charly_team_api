package paginator

import (
	"context"

	"github.com/go-rel/rel"
)

type Paginable interface {
	// Deprecated: CreatePaginator is deprecated. Use CreatePagination instead
	CreatePaginator(
		ctx context.Context,
		tableName string,
		holder interface{},
		pageNumber int,
		queries ...rel.Querier,
	) (*Paginator, error)
	CreatePagination(
		ctx context.Context,
		tableName string,
		holder interface{},
		pageSort *PageSort,
	) (*Paginator, error)
}
