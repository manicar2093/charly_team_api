package main

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/handlers/biotestcreator"
	"github.com/manicar2093/charly_team_api/internal/apperrors"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(CreateBiotestAWSLambdaTests))
}

type CreateBiotestAWSLambdaTests struct {
	suite.Suite
	biotest             entities.Biotest
	ctx                 context.Context
	biotestCreator      *biotestcreator.MockBiotestCreator
	updateUserAWSLambda *CreateBiotestAWSLambda
	ordinaryError       error
}

func (c *CreateBiotestAWSLambdaTests) SetupTest() {
	c.biotest = entities.Biotest{}
	c.ctx = context.Background()
	c.biotestCreator = &biotestcreator.MockBiotestCreator{}
	c.updateUserAWSLambda = NewCreateBiotestAWSLambda(c.biotestCreator)
	c.ordinaryError = errors.New("An ordinary error :O")

}

func (c *CreateBiotestAWSLambdaTests) TearDownTest() {
	c.biotestCreator.AssertExpectations(c.T())
}

func (c *CreateBiotestAWSLambdaTests) TestHandler() {
	biotestID := int32(12)
	biotestUUID := "uuid"
	c.biotest.ID = biotestID
	c.biotest.BiotestUUID = biotestUUID
	c.biotestCreator.On("Run", c.ctx, &c.biotest).Return(
		&biotestcreator.BiotestCreatorResponse{BiotestCreated: &c.biotest},
		nil,
	)

	res, err := c.updateUserAWSLambda.Handler(c.ctx, c.biotest)

	c.Nil(err, "should not return an error")
	c.Equal(http.StatusCreated, res.StatusCode, "status code not correct")
	c.IsType(&CreateBiotestResponse{}, res.Body, "body is not correct type")
}

func (c *CreateBiotestAWSLambdaTests) TestHandler_ValidationError() {
	validationErrors := apperrors.ValidationErrors{
		{Field: "name", Validation: "required"},
		{Field: "last_name", Validation: "required"},
	}
	c.biotestCreator.On("Run", c.ctx, &c.biotest).Return(
		nil,
		validationErrors,
	)

	res, err := c.updateUserAWSLambda.Handler(c.ctx, c.biotest)

	bodyAsErrorResponse := res.Body.(models.ErrorReponse)
	c.Nil(err, "should not return an error")
	c.Equal(http.StatusBadRequest, res.StatusCode, "status code not correct")
	c.Len(bodyAsErrorResponse.Error.(apperrors.ValidationErrors), 2, "not correct errors returned")
}

func (c *CreateBiotestAWSLambdaTests) TestHandler_UnhandledError() {
	c.biotestCreator.On("Run", c.ctx, &c.biotest).Return(
		nil,
		c.ordinaryError,
	)

	res, err := c.updateUserAWSLambda.Handler(c.ctx, c.biotest)

	bodyAsErrorResponse := res.Body.(models.ErrorReponse)
	c.Nil(err, "should not return an error")
	c.Equal(http.StatusInternalServerError, res.StatusCode, "status code not correct")
	c.Equal(c.ordinaryError.Error(), bodyAsErrorResponse.Error, "not correct error returned")
}
