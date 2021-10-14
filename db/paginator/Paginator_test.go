package paginator

import (
	"context"
	"testing"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/stretchr/testify/assert"
)

func TestPaginator(t *testing.T) {

	repo := reltest.New()
	paginator := NewPaginable(repo)

	pageRequested := 1
	previousPageExpect := 0
	nextPageExpected := 2
	totalPagesExpected := 100

	dbTable := entities.UserTable

	var users []entities.User

	repo.ExpectCount(dbTable).Result(1000)
	repo.ExpectFindAll(
		rel.Limit(config.PageSize),
		rel.Offset(
			createOffsetValue(
				config.PageSize,
				pageRequested,
			),
		),
	).Result([]entities.User{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}})

	page, err := paginator.CreatePaginator(context.Background(), dbTable, &users, pageRequested)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, page.TotalPages, "Total pages should not be empty")
	assert.Equal(t, len(users), config.PageSize, "number of items does not correspond with page size")
	assert.Equal(t, page.TotalPages, totalPagesExpected, "total pages are incorrect")
	assert.Equal(t, page.NextPage, nextPageExpected, "next page is not correct")
	assert.Equal(t, page.PreviousPage, previousPageExpect, "previous page is not correct")
	assert.Equal(t, page.CurrendPage, pageRequested, "current page is not correct")

}

func TestPaginator_LastPage(t *testing.T) {

	repo := reltest.New()
	paginator := NewPaginable(repo)

	pageRequested := 100
	previousPageExpect := 99
	nextPageExpected := 1
	totalPagesExpected := 100

	dbTable := entities.UserTable

	var users []entities.User

	repo.ExpectCount(dbTable).Result(1000)
	repo.ExpectFindAll(
		rel.Limit(config.PageSize),
		rel.Offset(
			createOffsetValue(
				config.PageSize,
				pageRequested,
			),
		),
	).Result([]entities.User{{}, {}, {}, {}, {}, {}})

	page, err := paginator.CreatePaginator(context.Background(), dbTable, &users, pageRequested)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, page.TotalPages, "Total pages should not be empty")
	assert.GreaterOrEqual(t, config.PageSize, len(users), "number of items does not correspond with page size")
	assert.Equal(t, totalPagesExpected, page.TotalPages, "total pages are incorrect")
	assert.Equal(t, nextPageExpected, page.NextPage, "next page is not correct")
	assert.Equal(t, previousPageExpect, page.PreviousPage, "previous page is not correct")
	assert.Equal(t, pageRequested, page.CurrendPage, "current page is not correct")

}

func TestPaginator_PageDoesNotExist(t *testing.T) {

	repo := reltest.New()
	paginator := NewPaginable(repo)

	pageRequested := 101

	dbTable := entities.UserTable

	var users []entities.User

	repo.ExpectCount(dbTable).Result(1000)
	repo.ExpectFindAll(
		rel.Limit(config.PageSize),
		rel.Offset(
			createOffsetValue(
				config.PageSize,
				pageRequested,
			),
		),
	).Result([]entities.User{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}})

	_, err := paginator.CreatePaginator(context.Background(), dbTable, &users, pageRequested)

	_, ok := err.(PageError)
	assert.True(t, ok, "Error returned wrong")

}

func TestPaginator_ReturnOnePageIfEntriesLessThanPageSize(t *testing.T) {

	repo := reltest.New()
	paginator := NewPaginable(repo)

	pageRequested := 1
	previousPageExpect := 0
	nextPageExpected := 1
	totalPagesExpected := 1

	dbTable := entities.UserTable

	var users []entities.User

	repo.ExpectCount(dbTable).Result(2)
	repo.ExpectFindAll(
		rel.Limit(config.PageSize),
		rel.Offset(
			createOffsetValue(
				config.PageSize,
				pageRequested,
			),
		),
	).Result([]entities.User{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}})

	page, err := paginator.CreatePaginator(context.Background(), dbTable, &users, pageRequested)
	assert.Nil(t, err, "should not be an error")

	assert.NotEmpty(t, page.TotalPages, "Total pages should not be empty")
	assert.Equal(t, len(users), config.PageSize, "number of items does not correspond with page size")
	assert.Equal(t, page.TotalPages, totalPagesExpected, "total pages are incorrect")
	assert.Equal(t, page.NextPage, nextPageExpected, "next page is not correct")
	assert.Equal(t, page.PreviousPage, previousPageExpect, "previous page is not correct")
	assert.Equal(t, page.CurrendPage, pageRequested, "current page is not correct")

}
