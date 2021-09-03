package repositories

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/models"
)

type Paginable interface {
	CreatePaginator(
		ctx context.Context,
		tableName string,
		holder interface{},
		pageNumber int,
		queries ...rel.Querier,
	) (*models.Paginator, error)
}

type PaginableImpl struct {
	repo rel.Repository
}

func NewPaginable(repo rel.Repository) Paginable {
	return PaginableImpl{repo: repo}
}

func (c PaginableImpl) CreatePaginator(
	ctx context.Context,
	tableName string,
	holder interface{},
	pageNumber int,
	queries ...rel.Querier,
) (*models.Paginator, error) {
	totalEntries, err := c.repo.Count(ctx, tableName)
	if err != nil {
		return nil, err
	}

	pageLimitQuery := rel.Limit(config.PageSize)
	pageOffsetQuery := rel.Offset(config.PageSize * (pageNumber - 1))

	queries = append(queries, pageLimitQuery, pageOffsetQuery)

	err = c.repo.FindAll(ctx, &holder, queries...)
	if err != nil {
		return nil, err
	}

	return &models.Paginator{
		TotalPages: totalEntries / config.PageSize,
		Data:       holder,
	}, nil
}
