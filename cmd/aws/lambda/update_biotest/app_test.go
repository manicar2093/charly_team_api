package main

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/handlers/biotestupdater"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/testfunc"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(UpdateBiotestAWSLambdaTests))
}

type UpdateBiotestAWSLambdaTests struct {
	suite.Suite
	biotest             entities.Biotest
	ctx                 context.Context
	biotestUpdater      *biotestupdater.MockBiotestUpdater
	updateUserAWSLambda *UpdateBiotestAWSLambda
	ordinaryError       error
}

func (c *UpdateBiotestAWSLambdaTests) SetupTest() {
	c.biotest = entities.Biotest{}
	c.ctx = context.Background()
	c.biotestUpdater = &biotestupdater.MockBiotestUpdater{}
	c.updateUserAWSLambda = NewUpdateBiotestAWSLambda(c.biotestUpdater)
	c.ordinaryError = errors.New("An ordinary error :O")

}

func (c *UpdateBiotestAWSLambdaTests) TearDownTest() {
	c.biotestUpdater.AssertExpectations(c.T())
}

func (c *UpdateBiotestAWSLambdaTests) TestHandler() {
	c.biotestUpdater.On("Run", c.ctx, &c.biotest).Return(
		&biotestupdater.BiotestUpdaterResponse{BiotestUpdated: &c.biotest},
		nil,
	)

	res, err := c.updateUserAWSLambda.Handler(c.ctx, c.biotest)
	testfunc.PrintJsonIndented(res)
	c.Nil(err, "should not return an error")
	c.Equal(http.StatusOK, res.StatusCode, "status code not correct")
	c.IsType(&entities.Biotest{}, res.Body, "body is not correct type")
}

func (c *UpdateBiotestAWSLambdaTests) TestHandler_ValidationError() {
	validationErrors := apperrors.ValidationErrors{
		{Field: "name", Validation: "required"},
		{Field: "last_name", Validation: "required"},
	}
	c.biotestUpdater.On("Run", c.ctx, &c.biotest).Return(
		nil,
		validationErrors,
	)

	res, err := c.updateUserAWSLambda.Handler(c.ctx, c.biotest)

	bodyAsErrorResponse := res.Body.(models.ErrorReponse)
	c.Nil(err, "should not return an error")
	c.Equal(http.StatusBadRequest, res.StatusCode, "status code not correct")
	c.Len(bodyAsErrorResponse.Error.(apperrors.ValidationErrors), 2, "not correct errors returned")
}

func (c *UpdateBiotestAWSLambdaTests) TestHandler_UnhandledError() {
	c.biotestUpdater.On("Run", c.ctx, &c.biotest).Return(
		nil,
		c.ordinaryError,
	)

	res, err := c.updateUserAWSLambda.Handler(c.ctx, c.biotest)

	bodyAsErrorResponse := res.Body.(models.ErrorReponse)
	c.Nil(err, "should not return an error")
	c.Equal(http.StatusInternalServerError, res.StatusCode, "status code not correct")
	c.Equal(c.ordinaryError.Error(), bodyAsErrorResponse.Error, "not correct error returned")
}
