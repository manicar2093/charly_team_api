package biotestfilters

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
	"github.com/go-rel/rel/sort"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/filters"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/stretchr/testify/suite"
)

func TestGetComparision(t *testing.T) {
	suite.Run(t, new(GetComparisionTest))
}

type GetComparisionTest struct {
	suite.Suite
	repo                         *reltest.Repository
	paginator                    *mocks.Paginable
	ctx                          context.Context
	filterParams                 filters.FilterParameters
	notFoundError, ordinaryError error
}

func (c *GetComparisionTest) SetupTest() {
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

func (c *GetComparisionTest) TearDownTest() {
	c.repo.AssertExpectations(c.T())
	c.paginator.AssertExpectations(c.T())
}

func (c *GetComparisionTest) TestGetComparision() {

	userUUIDRequest := "an-uuid"
	userID := int32(1)

	request := map[string]interface{}{
		"user_uuid": userUUIDRequest,
	}
	biotestResponses := []entities.Biotest{
		{ID: 1, BiotestUUID: "uuid1", CreatedAt: time.Now()},
		{ID: 2, BiotestUUID: "uuid2", CreatedAt: time.Now()},
	}
	biotestDetails := []BiotestDetails{
		{BiotestUUID: "uuid1", CreatedAt: time.Now()},
		{BiotestUUID: "uuid2", CreatedAt: time.Now()},
	}

	c.repo.ExpectFind(
		where.Eq("user_uuid", userUUIDRequest),
	).Result(
		entities.User{
			ID:       userID,
			UserUUID: userUUIDRequest,
		},
	)

	c.repo.ExpectFindAll(
		rel.Select("biotest_uuid", "created_at").From(entities.BiotestTable),
	).Result(biotestDetails)

	c.repo.ExpectFind(
		where.Eq("customer_id", userID),
		sort.Asc("created_at"),
	).Result(biotestResponses[0])

	c.repo.ExpectFind(
		where.Eq("customer_id", userID),
		sort.Desc("created_at"),
	).Result(biotestResponses[1])

	c.filterParams.Values = request

	got, err := GetBiotestComparision(&c.filterParams)

	c.Nil(err, "return an error")

	_, ok := got.(BiotestComparisionResponse)
	c.True(ok, "unexpected answare type")

}

func (c *GetComparisionTest) TestGetComparision_NoUserUUIDOnRequest() {

	request := map[string]interface{}{}

	c.filterParams.Values = request

	_, err := GetBiotestComparision(&c.filterParams)

	c.NotNil(err, "should return an error")

	validationError, ok := err.(apperrors.ValidationError)
	c.True(ok, "unexpected answare type")
	c.Equal("user_uuid", validationError.Field)
	c.Equal("required", validationError.Validation)

}

func (c *GetComparisionTest) TestGetComparision_UserNotFound() {

	userUUIDRequest := "an-uuid"

	request := map[string]interface{}{
		"user_uuid": userUUIDRequest,
	}

	c.repo.ExpectFind(
		where.Eq("user_uuid", userUUIDRequest),
	).NotFound()

	c.filterParams.Values = request

	_, err := GetBiotestComparision(&c.filterParams)

	c.NotNil(err, "should return an error")

	notFoundError, ok := err.(apperrors.NotFoundError)
	c.True(ok, "unexpected answare type")
	c.Contains(notFoundError.Error(), "does not exist")

}

func (c *GetComparisionTest) TestGetComparision_UserError() {

	userUUIDRequest := "an-uuid"

	request := map[string]interface{}{
		"user_uuid": userUUIDRequest,
	}

	c.repo.ExpectFind(
		where.Eq("user_uuid", userUUIDRequest),
	).Error(c.ordinaryError)

	c.filterParams.Values = request

	_, err := GetBiotestComparision(&c.filterParams)

	c.NotNil(err, "should return an error")

	c.Equal(err.Error(), c.ordinaryError.Error())

}

func (c *GetComparisionTest) TestGetComparision_UserHasNoBiotest() {

	userUUIDRequest := "an-uuid"
	userID := int32(1)

	request := map[string]interface{}{
		"user_uuid": userUUIDRequest,
	}

	biotestDetails := []BiotestDetails{}

	c.repo.ExpectFind(
		where.Eq("user_uuid", userUUIDRequest),
	).Result(
		entities.User{
			ID:       userID,
			UserUUID: userUUIDRequest,
		},
	)

	c.repo.ExpectFindAll(
		rel.Select("biotest_uuid", "created_at").From(entities.BiotestTable),
	).Result(biotestDetails)

	c.filterParams.Values = request

	_, err := GetBiotestComparision(&c.filterParams)

	c.NotNil(err, "should return an error")

	notFoundError, ok := err.(apperrors.NotFoundError)
	c.True(ok, "unexpected answare type")

	c.Contains(notFoundError.Error(), "no biotests")

}

func (c *GetComparisionTest) TestGetComparision_ErrorBiotestDetails() {

	userUUIDRequest := "an-uuid"
	userID := int32(1)

	request := map[string]interface{}{
		"user_uuid": userUUIDRequest,
	}

	c.repo.ExpectFind(
		where.Eq("user_uuid", userUUIDRequest),
	).Result(
		entities.User{
			ID:       userID,
			UserUUID: userUUIDRequest,
		},
	)

	c.repo.ExpectFindAll(
		rel.Select("biotest_uuid", "created_at").From(entities.BiotestTable),
	).Error(c.ordinaryError)

	c.filterParams.Values = request

	_, err := GetBiotestComparision(&c.filterParams)

	c.NotNil(err, "should return an error")

	c.Equal(err.Error(), c.ordinaryError.Error())

}

func (c *GetComparisionTest) TestGetComparision_ErrorFirstBiotest() {

	userUUIDRequest := "an-uuid"
	userID := int32(1)

	request := map[string]interface{}{
		"user_uuid": userUUIDRequest,
	}

	biotestDetails := []BiotestDetails{
		{BiotestUUID: "uuid1", CreatedAt: time.Now()},
		{BiotestUUID: "uuid2", CreatedAt: time.Now()},
	}

	c.repo.ExpectFind(
		where.Eq("user_uuid", userUUIDRequest),
	).Result(
		entities.User{
			ID:       userID,
			UserUUID: userUUIDRequest,
		},
	)

	c.repo.ExpectFindAll(
		rel.Select("biotest_uuid", "created_at").From(entities.BiotestTable),
	).Result(biotestDetails)

	c.repo.ExpectFind(
		where.Eq("customer_id", userID),
		sort.Asc("created_at"),
	).Error(c.ordinaryError)

	c.filterParams.Values = request

	_, err := GetBiotestComparision(&c.filterParams)

	c.NotNil(err, "should return an error")

	c.Equal(err.Error(), c.ordinaryError.Error())

}

func (c *GetComparisionTest) TestGetComparision_ErrorLastBiotest() {

	userUUIDRequest := "an-uuid"
	userID := int32(1)

	request := map[string]interface{}{
		"user_uuid": userUUIDRequest,
	}
	biotestResponses := []entities.Biotest{
		{ID: 1, BiotestUUID: "uuid1", CreatedAt: time.Now()},
		{ID: 2, BiotestUUID: "uuid2", CreatedAt: time.Now()},
	}
	biotestDetails := []BiotestDetails{
		{BiotestUUID: "uuid1", CreatedAt: time.Now()},
		{BiotestUUID: "uuid2", CreatedAt: time.Now()},
	}

	c.repo.ExpectFind(
		where.Eq("user_uuid", userUUIDRequest),
	).Result(
		entities.User{
			ID:       userID,
			UserUUID: userUUIDRequest,
		},
	)

	c.repo.ExpectFindAll(
		rel.Select("biotest_uuid", "created_at").From(entities.BiotestTable),
	).Result(biotestDetails)

	c.repo.ExpectFind(
		where.Eq("customer_id", userID),
		sort.Asc("created_at"),
	).Result(biotestResponses[0])

	c.repo.ExpectFind(
		where.Eq("customer_id", userID),
		sort.Desc("created_at"),
	).Error(c.ordinaryError)

	c.filterParams.Values = request

	_, err := GetBiotestComparision(&c.filterParams)

	c.NotNil(err, "return an error")

	c.Equal(err.Error(), c.ordinaryError.Error())

}