package repositories

import (
	"context"
	"testing"
	"time"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
	"github.com/go-rel/rel/sort"
	"github.com/go-rel/rel/where"
	"github.com/jaswdr/faker"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/paginator"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/stretchr/testify/suite"
)

func TestBiotestRepository(t *testing.T) {
	suite.Run(t, new(BiotestRepositoryTest))
}

type BiotestRepositoryTest struct {
	suite.Suite
	paginator         *mocks.Paginable
	repo              *reltest.Repository
	biotestRepository BiotestRepository
	ctx               context.Context
	faker             faker.Faker
}

func (c *BiotestRepositoryTest) SetupTest() {
	c.repo = reltest.New()
	c.paginator = &mocks.Paginable{}
	c.biotestRepository = NewBiotestRepositoryRel(c.repo, c.paginator)
	c.ctx = context.TODO()
	c.faker = faker.New()
}

func (c *BiotestRepositoryTest) TearDownTest() {
	t := c.T()
	c.repo.AssertExpectations(t)
	c.paginator.AssertExpectations(t)
}

func (c *BiotestRepositoryTest) TestFindBiotestByUUID() {
	expectedBiotestUUID := c.faker.UUID().V4()
	expectedBiotest := entities.Biotest{}
	expectedBiotest.BiotestUUID = expectedBiotestUUID
	c.repo.ExpectFind(
		where.Eq("biotest_uuid", expectedBiotestUUID),
	).Result(expectedBiotest)

	got, err := c.biotestRepository.FindBiotestByUUID(c.ctx, expectedBiotestUUID)

	c.Nil(err, "FindBiotestByUUID should not return an error")
	c.Equal(expectedBiotestUUID, got.BiotestUUID, "unexpected biotest uuid")

}

func (c *BiotestRepositoryTest) TestFindBiotestByUUID_NotFound() {
	expectedBiotestUUID := c.faker.UUID().V4()
	expectedBiotest := entities.Biotest{}
	expectedBiotest.BiotestUUID = expectedBiotestUUID
	c.repo.ExpectFind(
		where.Eq("biotest_uuid", expectedBiotestUUID),
	).Error(rel.ErrNotFound)

	got, err := c.biotestRepository.FindBiotestByUUID(c.ctx, expectedBiotestUUID)

	c.Nil(got, "FindBiotestByUUID a entities.Biotest instance")
	_, ok := err.(rel.NotFoundError)
	c.True(ok, "unexpected answare type")

}

func (c *BiotestRepositoryTest) TestGetAllUserBiotestByUserUUID() {
	expectedUserUUID := c.faker.UUID().V4()
	expectedUserID := c.faker.Int32()
	userFindReturn := entities.User{
		ID:       expectedUserID,
		UserUUID: expectedUserUUID,
	}
	c.repo.ExpectFind(
		where.Eq("user_uuid", expectedUserUUID),
	).Result(userFindReturn)
	pageNumber := c.faker.Float64(2, 1, 2)
	pageNumberAsInt := int(pageNumber)
	pageSort := paginator.PageSort{
		Page: pageNumber,
	}
	pageSort.SetFiltersQueries(where.Eq("customer_id", expectedUserID), sort.Asc("created_at"))
	biotestPaginatorResponse := []entities.Biotest{
		{ID: 1, BiotestUUID: "uuid1", CreatedAt: time.Now()},
		{ID: 2, BiotestUUID: "uuid2", CreatedAt: time.Now()},
	}
	pageResponse := &paginator.Paginator{
		TotalPages:   2,
		CurrentPage:  pageNumberAsInt,
		PreviousPage: 0,
		NextPage:     2,
		Data:         biotestPaginatorResponse,
	}
	biotestHolder := []entities.Biotest{}
	c.paginator.On(
		"CreatePagination",
		c.ctx,
		entities.BiotestTable,
		&biotestHolder,
		&pageSort,
	).Return(pageResponse, nil)

	got, err := c.biotestRepository.GetAllUserBiotestByUserUUID(c.ctx, &pageSort, expectedUserUUID)

	c.Nil(err, "should not return an error")
	c.Equal(2, len(got.Data.([]entities.Biotest)), "wrong data length")

}

func (c *BiotestRepositoryTest) TestGetAllUserBiotestByUserUUID_UserNotFound() {
	expectedUserUUID := c.faker.UUID().V4()
	c.repo.ExpectFind(
		where.Eq("user_uuid", expectedUserUUID),
	).Error(rel.ErrNotFound)
	pageNumber := c.faker.Float64(2, 1, 2)
	pageSort := paginator.PageSort{
		Page: pageNumber,
	}

	got, err := c.biotestRepository.GetAllUserBiotestByUserUUID(c.ctx, &pageSort, expectedUserUUID)

	c.NotNil(err, "should not return an error")
	c.IsType(err, NotFoundError{}, "error type is not correct")
	c.Nil(got, "should return a nil paginator")
}

func (c *BiotestRepositoryTest) TestGetAllUserBiotestByUserUUIDAsCatalog() {
	expectedUserUUID := c.faker.UUID().V4()
	expectedUserID := c.faker.Int32()
	userFindReturn := entities.User{
		ID:       expectedUserID,
		UserUUID: expectedUserUUID,
	}
	c.repo.ExpectFind(
		where.Eq("user_uuid", expectedUserUUID),
	).Result(userFindReturn)
	pageNumber := c.faker.Float64(2, 1, 2)
	pageNumberAsInt := int(pageNumber)
	pageSort := paginator.PageSort{
		Page: pageNumber,
	}
	pageSort.SetFiltersQueries(
		where.Eq("customer_id", expectedUserID),
		sort.Asc("created_at"),
		rel.Select("biotest_uuid", "created_at").From(entities.BiotestTable),
	)
	biotestPaginatorResponse := []entities.Biotest{
		{ID: 1, BiotestUUID: "uuid1", CreatedAt: time.Now()},
		{ID: 2, BiotestUUID: "uuid2", CreatedAt: time.Now()},
	}
	pageResponse := &paginator.Paginator{
		TotalPages:   2,
		CurrentPage:  pageNumberAsInt,
		PreviousPage: 0,
		NextPage:     2,
		Data:         biotestPaginatorResponse,
	}
	biotestHolder := []entities.Biotest{}
	c.paginator.On(
		"CreatePagination",
		c.ctx,
		entities.BiotestTable,
		&biotestHolder,
		&pageSort,
	).Return(pageResponse, nil)

	got, err := c.biotestRepository.GetAllUserBiotestByUserUUIDAsCatalog(c.ctx, &pageSort, expectedUserUUID)

	c.Nil(err, "should not return an error")
	c.Equal(2, len(got.Data.([]entities.Biotest)), "wrong data length")

}

func (c *BiotestRepositoryTest) TestGetAllUserBiotestByUserUUIDAsCatalog_UserNotFound() {
	expectedUserUUID := c.faker.UUID().V4()
	c.repo.ExpectFind(
		where.Eq("user_uuid", expectedUserUUID),
	).Error(rel.ErrNotFound)
	pageNumber := c.faker.Float64(2, 1, 2)
	pageSort := paginator.PageSort{
		Page: pageNumber,
	}

	got, err := c.biotestRepository.GetAllUserBiotestByUserUUIDAsCatalog(c.ctx, &pageSort, expectedUserUUID)

	c.NotNil(err, "should not return an error")
	c.IsType(err, NotFoundError{}, "error type is not correct")
	c.Nil(got, "should return a nil paginator")
}