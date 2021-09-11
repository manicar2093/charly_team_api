package main

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/go-rel/rel/reltest"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/manicar2093/charly_team_api/validators"
	"github.com/stretchr/testify/suite"
)

type MainTests struct {
	suite.Suite
	repo          *reltest.Repository
	validator     mocks.ValidatorService
	ctx           context.Context
	ordinaryError error
}

func (c *MainTests) SetupTest() {
	c.repo = reltest.New()
	c.validator = mocks.ValidatorService{}
	c.ctx = context.Background()
	c.ordinaryError = errors.New("An ordinary error :O")

}

func (c *MainTests) TearDownTest() {
	c.validator.AssertExpectations(c.T())
	c.repo.AssertExpectations(c.T())
}

func (c *MainTests) TestUpdateBiotest_UpdateError() {

	biotestRequest := entities.Biotest{
		ID: 1,
	}

	c.validator.On("Validate", biotestRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.repo.ExpectUpdate().ForType("entities.Biotest").Return(c.ordinaryError)

	res, _ := CreateLambdaHandlerWDependencies(c.repo, &c.validator)(c.ctx, biotestRequest)

	c.Equal(res.StatusCode, http.StatusInternalServerError, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusInternalServerError), "http status is not correct")

}

func (c *MainTests) TestUpdateBiotest() {

	biotestRequest := entities.Biotest{
		ID: 1,
	}

	c.validator.On("Validate", biotestRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.repo.ExpectUpdate().ForType("entities.Biotest").Return(nil)

	res, _ := CreateLambdaHandlerWDependencies(c.repo, &c.validator)(c.ctx, biotestRequest)

	c.Equal(res.StatusCode, http.StatusOK, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusOK), "http status is not correct")

}

func (c *MainTests) TestUpdateBiotest_NoBiotestID() {

	biotestRequest := entities.Biotest{}

	res, _ := CreateLambdaHandlerWDependencies(c.repo, &c.validator)(c.ctx, biotestRequest)

	c.Equal(res.StatusCode, http.StatusBadRequest, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusBadRequest), "http status is not correct")

	bodyError := res.Body.(apperrors.ValidationErrors)
	c.Equal("id", bodyError[0].Field, "validation error is not correct")
	c.Equal("required", bodyError[0].Validation, "validation error is not correct")

}

func (c *MainTests) TestUpdateBiotest_NoValidRequest() {

	biotestRequest := entities.Biotest{ID: 1}

	validationErrors := apperrors.ValidationErrors{
		{Field: "weight", Validation: "required"},
		{Field: "height", Validation: "required"},
	}

	c.validator.On("Validate", biotestRequest).Return(validators.ValidateOutput{IsValid: false, Err: validationErrors})

	res, _ := CreateLambdaHandlerWDependencies(c.repo, &c.validator)(c.ctx, biotestRequest)

	c.Equal(res.StatusCode, http.StatusBadRequest, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusBadRequest), "http status is not correct")

	bodyAsMap := res.Body.(map[string]interface{})

	errorGot, ok := bodyAsMap["error"].(apperrors.ValidationErrors)
	c.True(ok, "error parsing error message")
	c.Equal(len(errorGot), 2, "error message should not be empty")

}

func TestMain(t *testing.T) {
	suite.Run(t, new(MainTests))
}
