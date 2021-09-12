package main

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/validators"
	"github.com/stretchr/testify/suite"
	"gopkg.in/guregu/null.v4"
)

type MainTests struct {
	suite.Suite
	userFilter                         *mocks.Filterable
	validator                          *mocks.ValidatorService
	ctx                                context.Context
	ordinaryError, filterNotFoundError error
}

func (c *MainTests) SetupTest() {
	c.userFilter = &mocks.Filterable{}
	c.validator = &mocks.ValidatorService{}
	c.ctx = context.Background()
	c.ordinaryError = errors.New("An ordinary error :O")
	c.filterNotFoundError = apperrors.BadRequestError{Message: "not exists"}

}

func (c *MainTests) TearDownTest() {
	c.userFilter.AssertExpectations(c.T())
}

func (c *MainTests) TestGetUserByFilter() {
	filterName := "get_user_by_uuid"
	userFilter := models.FilterRequest{FilterName: filterName, Values: "values"}

	userRunnerReturn := entities.User{
		ID:            1,
		BiotypeID:     null.IntFrom(1),
		BoneDensityID: null.IntFrom(1),
		RoleID:        1,
		GenderID:      null.IntFrom(1),
		UserUUID:      "an-uuid",
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
	c.userFilter.On("GetFilter", userFilter.FilterName).Return(nil)
	c.userFilter.On("SetContext", c.ctx)
	c.userFilter.On("SetValues", userFilter.Values)
	c.userFilter.On("Run").Return(userRunnerReturn, nil)

	res, _ := CreateLambdaHandlerWDependencies(c.validator, c.userFilter)(c.ctx, userFilter)

	c.Equal(res.StatusCode, http.StatusOK, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusOK), "http status is not correct")

	userResponse := res.Body.(entities.User)

	c.NotEmpty(userResponse.ID, "unexpected user id response")
}

func (c *MainTests) TestGetUserByFilter_NoFilter() {
	filterName := "get_user_by_uuid"
	userFilter := models.FilterRequest{FilterName: filterName, Values: "values"}

	c.validator.On(
		"Validate",
		userFilter,
	).Return(
		validators.ValidateOutput{
			IsValid: true,
			Err:     nil,
		},
	)
	c.userFilter.On("GetFilter", userFilter.FilterName).Return(c.filterNotFoundError)

	res, _ := CreateLambdaHandlerWDependencies(c.validator, c.userFilter)(c.ctx, userFilter)

	c.Equal(res.StatusCode, http.StatusBadRequest, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusBadRequest), "http status is not correct")

	bodyError := res.Body.(models.ErrorReponse)

	c.Contains(bodyError.Error, "not exists", "error message should not be empty")

}

func (c *MainTests) TestGetUserByFilter_RunError() {
	filterName := "get_user_by_uuid"
	userFilter := models.FilterRequest{FilterName: filterName, Values: "values"}

	c.validator.On(
		"Validate",
		userFilter,
	).Return(
		validators.ValidateOutput{
			IsValid: true,
			Err:     nil,
		},
	)
	c.userFilter.On("GetFilter", userFilter.FilterName).Return(nil)
	c.userFilter.On("SetContext", c.ctx)
	c.userFilter.On("SetValues", userFilter.Values)

	c.userFilter.On("Run").Return(nil, c.ordinaryError)

	res, _ := CreateLambdaHandlerWDependencies(c.validator, c.userFilter)(c.ctx, userFilter)

	c.Equal(res.StatusCode, http.StatusInternalServerError, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusInternalServerError), "http status is not correct")

	bodyError := res.Body.(models.ErrorReponse)

	c.Equal(bodyError.Error, c.ordinaryError.Error(), "error message should not be empty")

}

func (c *MainTests) TestGetUserByFilter_ValidationError() {
	filterName := "get_user_by_uuid"
	userFilter := models.FilterRequest{FilterName: filterName, Values: "values"}

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

	res, _ := CreateLambdaHandlerWDependencies(c.validator, c.userFilter)(c.ctx, userFilter)

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
