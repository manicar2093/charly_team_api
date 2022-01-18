package main

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/filters"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/validators"
	"github.com/stretchr/testify/suite"
)

type MainTests struct {
	suite.Suite
	biotestFilter                      *filters.MockFilterable
	validator                          *validators.MockValidatorService
	ctx                                context.Context
	ordinaryError, filterNotFoundError error
}

func (c *MainTests) SetupTest() {
	c.biotestFilter = &filters.MockFilterable{}
	c.validator = &validators.MockValidatorService{}
	c.ctx = context.Background()
	c.ordinaryError = errors.New("An ordinary error :O")
	c.filterNotFoundError = apperrors.BadRequestError{Message: "not exists"}

}

func (c *MainTests) TearDownTest() {
	c.biotestFilter.AssertExpectations(c.T())
}

func (c *MainTests) TestFindBiotestByUUID() {
	filterName := "find_biotest_by_uuid"
	biotestFilter := models.FilterRequest{FilterName: filterName, Values: "values"}

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

	c.biotestFilter.On("SetContext", c.ctx).Return(nil)
	c.biotestFilter.On("SetValues", biotestFilter.Values).Return(nil)
	c.biotestFilter.On("GetFilter", biotestFilter.FilterName).Return(nil)

	c.biotestFilter.On("Run").Return(biotestRunnerReturn, nil)

	res, _ := CreateLambdaHandlerWDependencies(c.validator, c.biotestFilter)(c.ctx, biotestFilter)

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
	c.biotestFilter.On("GetFilter", biotestFilter.FilterName).Return(c.filterNotFoundError)

	res, _ := CreateLambdaHandlerWDependencies(c.validator, c.biotestFilter)(c.ctx, biotestFilter)

	c.Equal(http.StatusBadRequest, res.StatusCode, "http status is not correct")
	c.Equal(http.StatusText(http.StatusBadRequest), res.Status, "http status is not correct")

	bodyError := res.Body.(models.ErrorReponse)

	c.Contains(bodyError.Error, "not exists", "error message should not be empty")

}

func (c *MainTests) TestFindBiotestByUUID_RunError() {
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

	c.biotestFilter.On("GetFilter", biotestFilter.FilterName).Return(nil)
	c.biotestFilter.On("SetContext", c.ctx).Return(nil)
	c.biotestFilter.On("SetValues", biotestFilter.Values).Return(nil)
	c.biotestFilter.On("Run").Return(nil, c.ordinaryError)

	res, _ := CreateLambdaHandlerWDependencies(c.validator, c.biotestFilter)(c.ctx, biotestFilter)

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

	res, _ := CreateLambdaHandlerWDependencies(c.validator, c.biotestFilter)(c.ctx, biotestFilter)

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
