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
	"gopkg.in/guregu/null.v4"
)

type MainTests struct {
	suite.Suite
	repo                         *reltest.Repository
	validator                    *mocks.ValidatorService
	paginator                    *mocks.Paginable
	userFilter                   *mocks.FilterService
	userRunable                  *mocks.FilterRunable
	ctx                          context.Context
	ordinaryError, notFoundError error
}

func (c *MainTests) SetupTest() {
	c.repo = reltest.New()
	c.validator = &mocks.ValidatorService{}
	c.paginator = &mocks.Paginable{}
	c.userFilter = &mocks.FilterService{}
	c.userRunable = &mocks.FilterRunable{}
	c.ctx = context.Background()
	c.ordinaryError = errors.New("An ordinary error :O")
	c.notFoundError = rel.NotFoundError{}

}

func (c *MainTests) TearDownTest() {
	c.repo.AssertExpectations(c.T())
	c.validator.AssertExpectations(c.T())
	c.paginator.AssertExpectations(c.T())
	c.userFilter.AssertExpectations(c.T())
}

func (c *MainTests) TestGetUserByFilter() {
	filterName := "get_user_by_id"
	userFilter := UserFilter{FilterName: filterName, Values: "values"}
	expectedRunnerParams := filters.FilterParameters{
		Ctx:       c.ctx,
		Repo:      c.repo,
		Values:    userFilter,
		Paginator: c.paginator,
	}
	userRunnerReturn := entities.User{
		ID:            1,
		BiotypeID:     null.IntFrom(1),
		BoneDensityID: null.IntFrom(1),
		RoleID:        1,
		GenderID:      null.IntFrom(1),
		Name:          "testing",
		LastName:      "testing",
		Email:         "testing@email.com",
		Birthday:      time.Now(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	c.validator.On(
		"Validate",
		userFilter,
	).Return(
		validators.ValidateOutput{
			IsValid: true,
			Err:     nil,
		},
	)
	c.userFilter.On("GetUserFilter", userFilter.FilterName).Return(c.userRunable)
	c.userRunable.On("IsFound").Return(true)

	c.userRunable.On("Run", &expectedRunnerParams).Return(userRunnerReturn, nil)

	res := CreateLambdaHandlerWDependencies(c.repo, c.validator, c.paginator, c.userFilter)(c.ctx, userFilter)

	c.Equal(res.StatusCode, http.StatusOK, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusOK), "http status is not correct")

	userResponse := res.Body.(entities.User)

	c.NotEmpty(userResponse.ID, "unexpected user id response")

}

func (c *MainTests) TestGetUserByFilter_NoFilter() {
	filterName := "get_user_by_id"
	userFilter := UserFilter{FilterName: filterName, Values: "values"}

	c.validator.On(
		"Validate",
		userFilter,
	).Return(
		validators.ValidateOutput{
			IsValid: true,
			Err:     nil,
		},
	)
	c.userFilter.On("GetUserFilter", userFilter.FilterName).Return(c.userRunable)
	c.userRunable.On("IsFound").Return(false)

	res := CreateLambdaHandlerWDependencies(c.repo, c.validator, c.paginator, c.userFilter)(c.ctx, userFilter)

	c.Equal(res.StatusCode, http.StatusBadRequest, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusBadRequest), "http status is not correct")

	bodyError := res.Body.(models.ErrorReponse)

	c.Contains(bodyError.Error, "not exists", "error message should not be empty")

}

func (c *MainTests) TestGetUserByFilter_RunError() {
	filterName := "get_user_by_id"
	userFilter := UserFilter{FilterName: filterName, Values: "values"}
	expectedRunnerParams := filters.FilterParameters{
		Ctx:       c.ctx,
		Repo:      c.repo,
		Values:    userFilter,
		Paginator: c.paginator,
	}

	c.validator.On(
		"Validate",
		userFilter,
	).Return(
		validators.ValidateOutput{
			IsValid: true,
			Err:     nil,
		},
	)
	c.userFilter.On("GetUserFilter", userFilter.FilterName).Return(c.userRunable)
	c.userRunable.On("IsFound").Return(true)

	c.userRunable.On("Run", &expectedRunnerParams).Return(nil, c.ordinaryError)

	res := CreateLambdaHandlerWDependencies(c.repo, c.validator, c.paginator, c.userFilter)(c.ctx, userFilter)

	c.Equal(res.StatusCode, http.StatusInternalServerError, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusInternalServerError), "http status is not correct")

	bodyError := res.Body.(models.ErrorReponse)

	c.Equal(bodyError.Error, c.ordinaryError.Error(), "error message should not be empty")

}

func (c *MainTests) TestGetUserByFilter_ValidationError() {
	filterName := "get_user_by_id"
	userFilter := UserFilter{FilterName: filterName, Values: "values"}

	validationErrors := apperrors.ValidationErrors{
		{Field: "filter_name", Validation: "required"},
	}

	c.validator.On(
		"Validate",
		userFilter,
	).Return(
		validators.ValidateOutput{
			IsValid: false,
			Err:     validationErrors,
		},
	)

	res := CreateLambdaHandlerWDependencies(c.repo, c.validator, c.paginator, c.userFilter)(c.ctx, userFilter)

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
