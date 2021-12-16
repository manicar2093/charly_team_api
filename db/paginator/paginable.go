package paginator

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/models"
)

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

	if totalEntries > 0 && totalEntries < config.PageSize {
		totalPages = 1
	}

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
		CurrentPage:  pageNumber,
		PreviousPage: pageNumber - 1,
		NextPage:     calculateNextPage(pageNumber, totalPages),
		Data:         holder,
	}, nil
}

func (c PaginableImpl) CreatePagination(
	ctx context.Context,
	tableName string,
	holder interface{},
	pageSort *PageSort,
) (*models.Paginator, error) {
	pageSize := pageSort.GetItemsPerPage()
	page := pageSort.GetPage()

	totalEntries, err := c.repo.Count(ctx, tableName, pageSort.GetFiltersQueries()...)
	if err != nil {
		return nil, err
	}

	totalPages := totalEntries / config.PageSize

	if totalEntries > 0 && totalEntries < config.PageSize {
		totalPages = 1
	}

	if page > totalPages {
		return nil, PageError{PageNumber: page}
	}

	pageLimitQuery := rel.Limit(pageSize)
	pageOffsetQuery := rel.Offset(createOffsetValue(pageSize, page))
	sorting := pageSort.GetSortQueries()
	queries := pageSort.GetFiltersQueries()

	queries = append(queries, sorting...)
	queries = append(queries, pageLimitQuery, pageOffsetQuery)

	err = c.repo.FindAll(ctx, holder, queries...)
	if err != nil {
		return nil, err
	}

	return &models.Paginator{
		TotalPages:   totalPages,
		CurrentPage:  page,
		PreviousPage: page - 1,
		NextPage:     calculateNextPage(page, totalPages),
		Data:         holder,
		TotalEntries: config.PageSize,
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
