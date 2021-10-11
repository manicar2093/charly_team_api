package paginator

import (
	"context"
	"fmt"
	"net/http"

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

type PageError struct {
	PageNumber int
}

func (c PageError) Error() string {
	return fmt.Sprintf("Page %v does not exists", c.PageNumber)
}

func (c PageError) StatusCode() int {
	return http.StatusBadRequest
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

	pageLimitQuery := rel.Limit(config.PageSize)
	pageOffsetQuery := rel.Offset(createOffsetValue(config.PageSize, pageNumber))

	totalEntries, err := c.repo.Count(ctx, tableName, queries...)
	if err != nil {
		return nil, err
	}

	totalPages := totalEntries / config.PageSize

	if pageNumber > totalPages {
		return nil, PageError{PageNumber: pageNumber}
	}

	queries = append(queries, pageLimitQuery, pageOffsetQuery)

	err = c.repo.FindAll(ctx, holder, queries...)
	if err != nil {
		return nil, err
	}

	return &models.Paginator{
		TotalPages:   totalPages,
		CurrendPage:  pageNumber,
		PreviousPage: pageNumber - 1,
		NextPage:     calculateNextPage(pageNumber, totalPages),
		Data:         holder,
	}, nil
}

func createOffsetValue(pageSize, pageNumber int) int {
	return config.PageSize * (pageNumber - 1)
}

func calculateNextPage(pageNumber, totalPages int) int {
	if pageNumber == totalPages {
		return 1
	}
	return pageNumber + 1
}
