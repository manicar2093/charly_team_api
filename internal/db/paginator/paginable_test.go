package paginator

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/go-rel/reltest"
	"github.com/manicar2093/health_records/internal/config"
	"github.com/manicar2093/health_records/internal/db/entities"
	"github.com/manicar2093/health_records/pkg/testfunc"
	"github.com/stretchr/testify/assert"
)

func TestPaginator(t *testing.T) {

	t.Skip("Implementation deprecated")

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
	assert.Equal(t, page.CurrentPage, pageRequested, "current page is not correct")
	assert.Equal(t, page.TotalEntries, config.PageSize, "total entries has no correct info")

}

func TestPaginator_LastPage(t *testing.T) {
	t.Skip("Implementation deprecated")
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
	assert.Equal(t, pageRequested, page.CurrentPage, "current page is not correct")
	assert.Equal(t, page.TotalEntries, config.PageSize, "total entries has no correct info")

}

func TestPaginator_PageDoesNotExist(t *testing.T) {
	t.Skip("Implementation deprecated")
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
	t.Skip("Implementation deprecated")
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
	assert.Equal(t, page.CurrentPage, pageRequested, "current page is not correct")
	assert.Equal(t, page.TotalEntries, config.PageSize, "total entries has no correct info")

}

func TestMain(m *testing.M) {
	testfunc.LoadEnvFileOrPanic("../../../.env")
	os.Exit(m.Run())
}

func TestPaginableImpl_CreatePagination__WFilters(t *testing.T) {

	repo := reltest.New()
	paginator := NewPaginable(repo)

	pageRequested := 1
	previousPageExpect := 0
	nextPageExpected := 2
	totalPagesExpected := 100

	pageSort := PageSort{
		Page:     float64(pageRequested),
		SortBy:   []string{"field_one", "field_two"},
		SortDesc: []bool{true, false},
	}
	userUUIDFilter := where.Eq("user_uuid", "an_uuid")
	userCreatedAtFilter := where.Eq("created_at", time.Now())
	pageSort.SetFiltersQueries(userUUIDFilter, userCreatedAtFilter)

	dbTable := entities.UserTable

	var users []entities.User
	expected_user_find_all := []entities.User{{}, {}, {}, {}, {}, {}, {}}

	repo.ExpectCount(dbTable, userUUIDFilter, userCreatedAtFilter).Result(1000)
	repo.ExpectFindAll(
		userUUIDFilter,
		userCreatedAtFilter,
		rel.NewSortDesc("field_one"),
		rel.NewSortAsc("field_two"),
		rel.Limit(config.PageSize),
		rel.Offset(
			createOffsetValue(
				config.PageSize,
				pageRequested,
			),
		),
	).Result(expected_user_find_all)

	page, err := paginator.CreatePagination(context.Background(), dbTable, &users, &pageSort)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, page.TotalPages, "Total pages should not be empty")
	assert.Equal(t, page.TotalPages, totalPagesExpected, "total pages are incorrect")
	assert.Equal(t, page.NextPage, nextPageExpected, "next page is not correct")
	assert.Equal(t, page.PreviousPage, previousPageExpect, "previous page is not correct")
	assert.Equal(t, page.CurrentPage, pageRequested, "current page is not correct")
	assert.Equal(t, page.TotalEntries, len(expected_user_find_all), "total entries has no correct info")

}

func TestPaginableImpl_CreatePagination__WOFilters(t *testing.T) {

	repo := reltest.New()
	paginator := NewPaginable(repo)

	pageRequested := 1
	previousPageExpect := 0
	nextPageExpected := 2
	totalPagesExpected := 100

	dbTable := entities.UserTable

	var users []entities.User
	expected_users_find_all := []entities.User{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}}

	repo.ExpectCount(dbTable).Result(1000)
	repo.ExpectFindAll(
		rel.Limit(config.PageSize),
		rel.Offset(
			createOffsetValue(
				config.PageSize,
				pageRequested,
			),
		),
	).Result(expected_users_find_all)

	page, err := paginator.CreatePagination(context.Background(), dbTable, &users, &PageSort{Page: 1})
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, page.TotalPages, "Total pages should not be empty")
	assert.Equal(t, len(users), config.PageSize, "number of items does not correspond with page size")
	assert.Equal(t, page.TotalPages, totalPagesExpected, "total pages are incorrect")
	assert.Equal(t, page.NextPage, nextPageExpected, "next page is not correct")
	assert.Equal(t, page.PreviousPage, previousPageExpect, "previous page is not correct")
	assert.Equal(t, page.CurrentPage, pageRequested, "current page is not correct")
	assert.Equal(t, page.TotalEntries, len(expected_users_find_all), "total entries has no correct info")

}

func TestPaginator_CreatePagination__LastPage(t *testing.T) {

	repo := reltest.New()
	paginator := NewPaginable(repo)

	pageRequested := 100
	previousPageExpect := 99
	nextPageExpected := 1
	totalPagesExpected := 100

	dbTable := entities.UserTable

	var users []entities.User
	expected_users_find_all := []entities.User{{}, {}, {}, {}, {}, {}}

	repo.ExpectCount(dbTable).Result(1000)
	repo.ExpectFindAll(
		rel.Limit(config.PageSize),
		rel.Offset(
			createOffsetValue(
				config.PageSize,
				pageRequested,
			),
		),
	).Result(expected_users_find_all)

	page, err := paginator.CreatePagination(context.Background(), dbTable, &users, &PageSort{Page: float64(pageRequested)})
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, page.TotalPages, "Total pages should not be empty")
	assert.GreaterOrEqual(t, config.PageSize, len(users), "number of items does not correspond with page size")
	assert.Equal(t, totalPagesExpected, page.TotalPages, "total pages are incorrect")
	assert.Equal(t, nextPageExpected, page.NextPage, "next page is not correct")
	assert.Equal(t, previousPageExpect, page.PreviousPage, "previous page is not correct")
	assert.Equal(t, pageRequested, page.CurrentPage, "current page is not correct")
	assert.Equal(t, page.TotalEntries, len(expected_users_find_all), "total entries has no correct info")

}

func TestPaginator_CreatePagination__PageDoesNotExist(t *testing.T) {

	repo := reltest.New()
	paginator := NewPaginable(repo)

	pageRequested := 101

	dbTable := entities.UserTable

	var users []entities.User
	expected_users_find_all := []entities.User{{}, {}, {}, {}, {}, {}}

	repo.ExpectCount(dbTable).Result(1000)
	repo.ExpectFindAll(
		rel.Limit(config.PageSize),
		rel.Offset(
			createOffsetValue(
				config.PageSize,
				pageRequested,
			),
		),
	).Result(expected_users_find_all)

	_, err := paginator.CreatePagination(context.Background(), dbTable, &users, &PageSort{Page: float64(pageRequested)})

	_, ok := err.(PageError)
	assert.True(t, ok, "Error returned wrong")

}

func TestPaginator_CreatePagination__ReturnOnePageIfEntriesLessThanPageSize(t *testing.T) {

	repo := reltest.New()
	paginator := NewPaginable(repo)

	pageRequested := 1
	previousPageExpect := 0
	nextPageExpected := 1
	totalPagesExpected := 1

	dbTable := entities.UserTable

	var users []entities.User
	expected_users_find_all := []entities.User{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}}

	repo.ExpectCount(dbTable).Result(2)
	repo.ExpectFindAll(
		rel.Limit(config.PageSize),
		rel.Offset(
			createOffsetValue(
				config.PageSize,
				pageRequested,
			),
		),
	).Result(expected_users_find_all)

	page, err := paginator.CreatePagination(context.Background(), dbTable, &users, &PageSort{Page: float64(pageRequested)})
	assert.Nil(t, err, "should not be an error")

	assert.NotEmpty(t, page.TotalPages, "Total pages should not be empty")
	assert.Equal(t, len(users), config.PageSize, "number of items does not correspond with page size")
	assert.Equal(t, page.TotalPages, totalPagesExpected, "total pages are incorrect")
	assert.Equal(t, page.NextPage, nextPageExpected, "next page is not correct")
	assert.Equal(t, page.PreviousPage, previousPageExpect, "previous page is not correct")
	assert.Equal(t, page.CurrentPage, pageRequested, "current page is not correct")
	assert.Equal(t, page.TotalEntries, len(expected_users_find_all), "total entries has no correct info")

}

// TestPaginator_CreatePagination__TotalEntriesAndPageSizeAreCorrect validates that TotalEntries and PageSize are filled
// with the correct data
func TestPaginator_CreatePagination__TotalEntriesAndPageSizeAreCorrect(t *testing.T) {

	repo := reltest.New()
	paginator := NewPaginable(repo)

	pageRequested := 1
	previousPageExpect := 0
	nextPageExpected := 1
	totalPagesExpected := 1

	dbTable := entities.UserTable

	var users []entities.User
	expected_users_find_all := []entities.User{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}}

	repo.ExpectCount(dbTable).Result(2)
	repo.ExpectFindAll(
		rel.Limit(config.PageSize),
		rel.Offset(
			createOffsetValue(
				config.PageSize,
				pageRequested,
			),
		),
	).Result(expected_users_find_all)

	page, err := paginator.CreatePagination(context.Background(), dbTable, &users, &PageSort{Page: float64(pageRequested)})
	assert.Nil(t, err, "should not be an error")

	assert.NotEmpty(t, page.TotalPages, "Total pages should not be empty")
	assert.Equal(t, len(users), config.PageSize, "number of items does not correspond with page size")
	assert.Equal(t, page.TotalPages, totalPagesExpected, "total pages are incorrect")
	assert.Equal(t, page.NextPage, nextPageExpected, "next page is not correct")
	assert.Equal(t, page.PreviousPage, previousPageExpect, "previous page is not correct")
	assert.Equal(t, page.CurrentPage, pageRequested, "current page is not correct")
	assert.Equal(t, page.TotalEntries, len(expected_users_find_all), "total entries has no correct info")
	assert.Equal(t, page.PageSize, config.PageSize, "total entries has no correct info")

}
