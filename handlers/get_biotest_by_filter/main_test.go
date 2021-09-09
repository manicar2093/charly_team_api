package main

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/filters"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/validators"
	"github.com/stretchr/testify/suite"
)

type MainTests struct {
	suite.Suite
	repo                         *reltest.Repository
	validator                    *mocks.ValidatorService
	paginator                    *mocks.Paginable
	biotestFilter                *mocks.FilterService
	userRunable                  *mocks.FilterRunable
	ctx                          context.Context
	ordinaryError, notFoundError error
}

func (c *MainTests) SetupTest() {
	c.repo = reltest.New()
	c.validator = &mocks.ValidatorService{}
	c.paginator = &mocks.Paginable{}
	c.biotestFilter = &mocks.FilterService{}
	c.userRunable = &mocks.FilterRunable{}
	c.ctx = context.Background()
	c.ordinaryError = errors.New("An ordinary error :O")
	c.notFoundError = rel.NotFoundError{}

}

func (c *MainTests) TearDownTest() {
	c.repo.AssertExpectations(c.T())
	c.validator.AssertExpectations(c.T())
	c.paginator.AssertExpectations(c.T())
	c.biotestFilter.AssertExpectations(c.T())
}

func (c *MainTests) TestFindBiotestByUUID() {
	filterName := "find_biotest_by_uuid"
	biotestFilter := models.FilterRequest{FilterName: filterName, Values: "values"}
	expectedRunnerParams := filters.FilterParameters{
		Ctx:       c.ctx,
		Repo:      c.repo,
		Values:    biotestFilter,
		Paginator: c.paginator,
	}
	biotestRunnerReturn := entities.Biotest{
		ID:                    1,
		HigherMuscleDensityID: 1,
		LowerMuscleDensityID:  1,
		SkinFoldsID:           1,
		WeightClasificationID: 1,
		HeartHealthID:         1,
		CustomerID:            1,
		CreatorID:             1,
		BiotestUUID:           "uuid",
		CreatedAt:             time.Now(),
	}

	c.validator.On(
		"Validate",
		biotestFilter,
	).Return(
		validators.ValidateOutput{
			IsValid: true,
			Err:     nil,
		},
	)
	c.biotestFilter.On("GetFilter", biotestFilter.FilterName).Return(c.userRunable)
	c.userRunable.On("IsFound").Return(true)

	c.userRunable.On("Run", &expectedRunnerParams).Return(biotestRunnerReturn, nil)

	res, _ := CreateLambdaHandlerWDependencies(c.repo, c.validator, c.biotestFilter, c.paginator)(c.ctx, biotestFilter)

	c.Equal(res.StatusCode, http.StatusOK, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusOK), "http status is not correct")

	userResponse := res.Body.(entities.Biotest)

	c.NotEmpty(userResponse.ID, "unexpected user id response")

}

func (c *MainTests) TestFindBiotestByUUID_NoFilter() {
	filterName := "find_biotest_by_uuid"
	biotestFilter := models.FilterRequest{FilterName: filterName, Values: "values"}

	c.validator.On(
		"Validate",
		biotestFilter,
	).Return(
		validators.ValidateOutput{
			IsValid: true,
			Err:     nil,
		},
	)
	c.biotestFilter.On("GetFilter", biotestFilter.FilterName).Return(c.userRunable)
	c.userRunable.On("IsFound").Return(false)

	res, _ := CreateLambdaHandlerWDependencies(c.repo, c.validator, c.biotestFilter, c.paginator)(c.ctx, biotestFilter)

	c.Equal(res.StatusCode, http.StatusBadRequest, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusBadRequest), "http status is not correct")

	bodyError := res.Body.(models.ErrorReponse)

	c.Contains(bodyError.Error, "not exists", "error message should not be empty")

}

func (c *MainTests) TestFindBiotestByUUID_RunError() {
	filterName := "find_biotest_by_uuid"
	biotestFilter := models.FilterRequest{FilterName: filterName, Values: "values"}
	expectedRunnerParams := filters.FilterParameters{
		Ctx:       c.ctx,
		Repo:      c.repo,
		Values:    biotestFilter,
		Paginator: c.paginator,
	}

	c.validator.On(
		"Validate",
		biotestFilter,
	).Return(
		validators.ValidateOutput{
			IsValid: true,
			Err:     nil,
		},
	)
	c.biotestFilter.On("GetFilter", biotestFilter.FilterName).Return(c.userRunable)
	c.userRunable.On("IsFound").Return(true)

	c.userRunable.On("Run", &expectedRunnerParams).Return(nil, c.ordinaryError)

	res, _ := CreateLambdaHandlerWDependencies(c.repo, c.validator, c.biotestFilter, c.paginator)(c.ctx, biotestFilter)

	c.Equal(res.StatusCode, http.StatusInternalServerError, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusInternalServerError), "http status is not correct")

	bodyError := res.Body.(models.ErrorReponse)

	c.Equal(bodyError.Error, c.ordinaryError.Error(), "error message should not be empty")

}

func (c *MainTests) TestFindBiotestByUUID_ValidationError() {
	filterName := "find_biotest_by_uuid"
	biotestFilter := models.FilterRequest{FilterName: filterName, Values: "values"}

	validationErrors := apperrors.ValidationErrors{
		{Field: "filter_name", Validation: "required"},
	}

	c.validator.On(
		"Validate",
		biotestFilter,
	).Return(
		validators.ValidateOutput{
			IsValid: false,
			Err:     validationErrors,
		},
	)

	res, _ := CreateLambdaHandlerWDependencies(c.repo, c.validator, c.biotestFilter, c.paginator)(c.ctx, biotestFilter)

	c.Equal(res.StatusCode, http.StatusBadRequest, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusBadRequest), "http status is not correct")

	bodyAsMap := res.Body.(map[string]interface{})

	errorGot, ok := bodyAsMap["error"].(apperrors.ValidationErrors)
	c.True(ok, "error parsing error message")
	c.Equal(len(errorGot), 1, "error message should not be empty")

}

func TestMain(t *testing.T) {
	suite.Run(t, new(MainTests))
}
