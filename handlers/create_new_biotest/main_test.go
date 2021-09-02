package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
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

func (c *MainTests) TestCreateNewBiotest() {

	biotestRequest := entities.Biotest{}

	c.validator.On("Validate", biotestRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.repo.ExpectInsert().ForType("entities.Biotest").Return(nil)

	res := CreateLambdaHandlerWDependencies(c.repo, &c.validator)(c.ctx, biotestRequest)

	c.Equal(res.StatusCode, http.StatusCreated, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusCreated), "http status is not correct")

	createBiotestResponse := res.Body.(CreateBiotestResponse)

	c.NotEmpty(createBiotestResponse.BiotestID, "unexpected id biotest response")

}

func (c *MainTests) TestCreateNewBiotestError() {

	biotestRequest := entities.Biotest{}

	c.validator.On("Validate", biotestRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.repo.ExpectInsert().ForType("entities.Biotest").Return(c.ordinaryError)

	res := CreateLambdaHandlerWDependencies(c.repo, &c.validator)(c.ctx, biotestRequest)

	c.Equal(res.StatusCode, http.StatusInternalServerError, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusInternalServerError), "http status is not correct")

	bodyAsMap := res.Body.(map[string]interface{})

	errorGot, ok := bodyAsMap["error"].(string)
	c.True(ok, "error parsing error message")
	c.Equal(errorGot, c.ordinaryError.Error(), "error message should not be empty")

}

func (c *MainTests) TestRegistryNewUserNoValidReq() {

	biotestRequest := entities.Biotest{}

	validationErrors := apperrors.ValidationErrors{
		{Field: "weight", Validation: "required"},
		{Field: "height", Validation: "required"},
	}

	c.validator.On("Validate", biotestRequest).Return(validators.ValidateOutput{IsValid: false, Err: validationErrors})

	res := CreateLambdaHandlerWDependencies(c.repo, &c.validator)(c.ctx, biotestRequest)

	json.NewEncoder(os.Stdout).Encode(res)

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
