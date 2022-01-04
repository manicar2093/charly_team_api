package repositories

import (
	"context"
	"testing"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
	"github.com/go-rel/rel/where"
	"github.com/jaswdr/faker"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/stretchr/testify/suite"
)

func TestBiotestRepository(t *testing.T) {
	suite.Run(t, new(BiotestRepositoryTest))
}

type BiotestRepositoryTest struct {
	suite.Suite
	repo              *reltest.Repository
	biotestRepository BiotestRepository
	ctx               context.Context
	faker             faker.Faker
}

func (c *BiotestRepositoryTest) SetupTest() {
	c.repo = reltest.New()
	c.biotestRepository = NewBiotestRepositoryRel(c.repo)
	c.ctx = context.TODO()
	c.faker = faker.New()
}

func (c *BiotestRepositoryTest) TearDownTest() {
	c.repo.AssertExpectations(c.T())
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
