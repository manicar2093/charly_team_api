package main

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/validators"
	"github.com/stretchr/testify/suite"
)

type MainTests struct {
	suite.Suite
	repo                         *reltest.Repository
	validator                    mocks.ValidatorService
	ctx                          context.Context
	ordinaryError, notFoundError error
}

func (c *MainTests) SetupTest() {
	c.repo = reltest.New()
	c.validator = mocks.ValidatorService{}
	c.ctx = context.Background()
	c.ordinaryError = errors.New("An ordinary error :O")
	c.notFoundError = rel.NotFoundError{}

}

func (c *MainTests) TearDownTest() {
	c.validator.AssertExpectations(c.T())
	c.repo.AssertExpectations(c.T())
}

func (c *MainTests) TestGetBiotstByID() {

	biotestID := GetBiotestByID{BiotestID: 1}
	biotestFindResult := entities.Biotest{
		ID: 1,
	}

	c.validator.On("Validate", biotestID).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.repo.ExpectFind(rel.Eq("id", 1)).Result(biotestFindResult)

	res := CreateLambdaHandlerWDependencies(c.repo, &c.validator)(c.ctx, biotestID)

	c.Equal(res.StatusCode, http.StatusOK, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusOK), "http status is not correct")

	biotestResponse := res.Body.(entities.Biotest)

	c.NotEmpty(biotestResponse.ID, "unexpected biotest id response")

}

func (c *MainTests) TestGetBiotstByIDNotFound() {

	biotestID := GetBiotestByID{BiotestID: 1}

	c.validator.On("Validate", biotestID).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.repo.ExpectFind(rel.Eq("id", 1)).Return(c.notFoundError)

	res := CreateLambdaHandlerWDependencies(c.repo, &c.validator)(c.ctx, biotestID)

	c.Equal(res.StatusCode, http.StatusNotFound, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusNotFound), "http status is not correct")

	bodyError := res.Body.(models.ErrorReponse)

	c.Equal(bodyError.Error, c.notFoundError.Error(), "error message should not be empty")

}

func (c *MainTests) TestGetBiotestByIDError() {

	biotestID := GetBiotestByID{BiotestID: 1}

	c.validator.On("Validate", biotestID).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.repo.ExpectFind(rel.Eq("id", 1)).Return(c.ordinaryError)

	res := CreateLambdaHandlerWDependencies(c.repo, &c.validator)(c.ctx, biotestID)

	c.Equal(res.StatusCode, http.StatusInternalServerError, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusInternalServerError), "http status is not correct")

	bodyAsErr := res.Body.(models.ErrorReponse)

	c.Equal(bodyAsErr.Error, c.ordinaryError.Error(), "error message should not be empty")

}

func (c *MainTests) TestRegistryNewUserNoValidReq() {

	validationErrors := apperrors.ValidationErrors{
		{Field: "biotest_id", Validation: "required"},
	}

	biotestID := GetBiotestByID{BiotestID: 1}

	c.validator.On("Validate", biotestID).Return(validators.ValidateOutput{IsValid: false, Err: validationErrors})

	res := CreateLambdaHandlerWDependencies(c.repo, &c.validator)(c.ctx, biotestID)

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
