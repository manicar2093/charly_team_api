package main

import (
	"context"
	"errors"
	"testing"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/filters"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/stretchr/testify/suite"
)

type BiotestFilterTest struct {
	suite.Suite
	repo                         *reltest.Repository
	paginator                    *mocks.Paginable
	ctx                          context.Context
	filterParams                 filters.FilterParameters
	notFoundError, ordinaryError error
}

func (c *BiotestFilterTest) SetupTest() {
	c.repo = reltest.New()
	c.paginator = &mocks.Paginable{}
	c.ctx = context.Background()
	c.ordinaryError = errors.New("An ordinary error :O")
	c.filterParams = filters.FilterParameters{
		Ctx:       c.ctx,
		Repo:      c.repo,
		Paginator: c.paginator,
	}
	c.notFoundError = rel.NotFoundError{}

}

func (c *BiotestFilterTest) TearDownTest() {
	c.repo.AssertExpectations(c.T())
	c.paginator.AssertExpectations(c.T())
}

func (c *BiotestFilterTest) TestFindBiotestByUUID() {

	biotestUUIDRequested := "an_uuid"

	request := map[string]interface{}{
		"biotest_uuid": biotestUUIDRequested,
	}

	c.repo.ExpectFind(
		where.Eq("biotest_uuid", biotestUUIDRequested),
	).Result(
		entities.Biotest{
			ID:          int32(1),
			BiotestUUID: biotestUUIDRequested,
		},
	)

	c.filterParams.Values = request

	got, err := FindBiotestByUUID(&c.filterParams)

	c.Nil(err, "FindBiotestByUUID return an error")

	biotetsGot, ok := got.(entities.Biotest)
	c.True(ok, "unexpected answare type")
	c.Equal(biotetsGot.BiotestUUID, biotestUUIDRequested, "unexpected user id")

}

func (c *BiotestFilterTest) TestFindBiotestByUUID_ValidationError() {

	biotestUUIDRequested := "an_uuid"

	request := map[string]interface{}{
		"biotest_bad": biotestUUIDRequested,
	}

	c.filterParams.Values = request

	_, err := FindBiotestByUUID(&c.filterParams)

	c.NotNil(err, "FindBiotestByUUID not return an error")

	_, ok := err.(apperrors.ValidationError)
	c.True(ok, "unexpected answare type")

}

func (c *BiotestFilterTest) TestFindBiotestByUUID_NotFound() {

	biotestUUIDRequested := "an_uuid"

	request := map[string]interface{}{
		"biotest_uuid": biotestUUIDRequested,
	}

	c.repo.ExpectFind(
		where.Eq("biotest_uuid", biotestUUIDRequested),
	).Return(c.notFoundError)

	c.filterParams.Values = request

	_, err := FindBiotestByUUID(&c.filterParams)

	c.NotNil(err, "FindBiotestByUUID not return an error")

	_, ok := err.(apperrors.NotFoundError)
	c.True(ok, "unexpected answare type")

	c.Contains(err.Error(), biotestUUIDRequested, "error is not representative")

}

func (c *BiotestFilterTest) TestFindBiotestByUUID_UnhandledError() {

	biotestUUIDRequested := "an_uuid"

	request := map[string]interface{}{
		"biotest_uuid": biotestUUIDRequested,
	}

	c.repo.ExpectFind(
		where.Eq("biotest_uuid", biotestUUIDRequested),
	).Return(c.ordinaryError)

	c.filterParams.Values = request

	_, err := FindBiotestByUUID(&c.filterParams)

	c.NotNil(err, "FindBiotestByUUID not return an error")

}

func TestUserFilter(t *testing.T) {
	suite.Run(t, new(BiotestFilterTest))
}
