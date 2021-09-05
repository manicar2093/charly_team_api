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
