package biotestfilters

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/filters"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/validators"
	"github.com/stretchr/testify/suite"
)

func TestGetAlluserBiotest(t *testing.T) {
	suite.Run(t, new(GetAllUserBiotestTest))
}

type GetAllUserBiotestTest struct {
	suite.Suite
	repo                         *reltest.Repository
	paginator                    *mocks.Paginable
	validator                    *mocks.ValidatorService
	ctx                          context.Context
	filterParams                 filters.FilterParameters
	notFoundError, ordinaryError error
}

func (c *GetAllUserBiotestTest) SetupTest() {
	c.repo = reltest.New()
	c.paginator = &mocks.Paginable{}
	c.validator = &mocks.ValidatorService{}
	c.ctx = context.Background()
	c.ordinaryError = errors.New("An ordinary error :O")
	c.filterParams = filters.FilterParameters{
		Ctx:       c.ctx,
		Repo:      c.repo,
		Paginator: c.paginator,
		Validator: c.validator,
	}
	c.notFoundError = rel.NotFoundError{}

}

func (c *GetAllUserBiotestTest) TearDownTest() {
	c.repo.AssertExpectations(c.T())
	c.paginator.AssertExpectations(c.T())
}

func (c *GetAllUserBiotestTest) TestGetAllUserBiotest() {

	userUUID := "an-uuid"
	pageNumber := float64(1)
	pageNumberAsInt := int(pageNumber)
	userID := int32(1)

	request := map[string]interface{}{
		"user_uuid":   userUUID,
		"page_number": pageNumber,
	}

	c.filterParams.Values = request

	biotestResponse := []entities.Biotest{
		{ID: 1, BiotestUUID: "uuid1", CreatedAt: time.Now()},
		{ID: 2, BiotestUUID: "uuid2", CreatedAt: time.Now()},
	}

	pageResponse := &models.Paginator{
		TotalPages:   2,
		CurrendPage:  pageNumberAsInt,
		PreviousPage: 0,
		NextPage:     2,
		Data:         biotestResponse,
	}

	c.validator.On(
		"Validate",
		&GetAllUserBiotestsRequest{
			userUUID,
			pageNumberAsInt,
		}).Return(validators.ValidateOutput{IsValid: true, Err: nil})

	c.repo.ExpectFind(
		where.Eq("user_uuid", userUUID),
	).Result(
		entities.User{
			ID:       userID,
			UserUUID: userUUID,
		},
	)

	var biotestHolder []entities.Biotest
	c.paginator.On(
		"CreatePaginator",
		c.ctx,
		entities.BiotestTable,
		&biotestHolder,
		pageNumberAsInt,
		where.Eq("customer_id", userID),
	).Return(pageResponse, nil)

	got, err := GetAllUserBiotest(&c.filterParams)

	c.Nil(err, "return an error")

	page, ok := got.(*models.Paginator)

	c.True(ok, "unexpected answare type")
	c.Equal(2, len(page.Data.([]entities.Biotest)), "Wrong data len")

}

func (c *GetAllUserBiotestTest) TestGetAllUserBiotest_NoUserUUID() {

	userUUID := ""
	pageNumber := float64(1)
	pageNumberAsInt := int(pageNumber)
	validationErrors := apperrors.ValidationErrors{
		{Field: "user_uuid", Validation: "required"},
	}

	request := map[string]interface{}{
		"user_uuid":   userUUID,
		"page_number": pageNumber,
	}

	c.filterParams.Values = request

	c.validator.On(
		"Validate",
		&GetAllUserBiotestsRequest{
			userUUID,
			pageNumberAsInt,
		}).Return(validators.ValidateOutput{IsValid: false, Err: validationErrors})

	_, err := GetAllUserBiotest(&c.filterParams)

	c.NotNil(err, "should return an error")

	gotErrors, _ := err.(apperrors.ValidationErrors)

	c.Len(gotErrors, 1, "errors gotten wrong")

}
