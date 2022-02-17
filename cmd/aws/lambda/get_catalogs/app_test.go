package main

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/manicar2093/charly_team_api/internal/handlers/catalog"
	"github.com/manicar2093/charly_team_api/pkg/apperrors"
	"github.com/manicar2093/charly_team_api/pkg/models"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(GetCatalogsAWSLambdaTests))
}

type GetCatalogsAWSLambdaTests struct {
	suite.Suite
	ctx                  context.Context
	catalogGetter        *catalog.MockCatalogGetter
	getCatalogsAWSLambda *GetCatalogsAWSLambda
	ordinaryError        error
}

func (c *GetCatalogsAWSLambdaTests) SetupTest() {
	c.ctx = context.Background()
	c.catalogGetter = &catalog.MockCatalogGetter{}
	c.getCatalogsAWSLambda = NewGetCatalogsAWSLambda(c.catalogGetter)
	c.ordinaryError = errors.New("An ordinary error :O")

}

func (c *GetCatalogsAWSLambdaTests) TearDownTest() {
	c.catalogGetter.AssertExpectations(c.T())
}

func (c *GetCatalogsAWSLambdaTests) TestHandler() {
	request := catalog.CatalogGetterRequest{
		CatalogNames: []string{"biotest", "roles"},
	}
	response := map[string]interface{}{
		"biotest": []interface{}{},
		"roles":   []interface{}{},
	}
	c.catalogGetter.On("Run", c.ctx, &request).Return(
		&catalog.CatalogGetterResponse{Catalogs: response},
		nil,
	)

	res, err := c.getCatalogsAWSLambda.Handler(c.ctx, request)

	c.Nil(err, "should not return an error")
	c.Equal(http.StatusOK, res.StatusCode, "status code not correct")
	c.IsType(map[string]interface{}{}, res.Body, "body is not correct type")
}

func (c *GetCatalogsAWSLambdaTests) TestHandler_ValidationError() {
	request := catalog.CatalogGetterRequest{
		CatalogNames: []string{"biotest", "roles"},
	}
	validationErrors := apperrors.ValidationErrors{
		{Field: "name", Validation: "required"},
		{Field: "last_name", Validation: "required"},
	}
	c.catalogGetter.On("Run", c.ctx, &request).Return(
		nil,
		validationErrors,
	)

	res, err := c.getCatalogsAWSLambda.Handler(c.ctx, request)

	c.Nil(err, "should not return an error")
	c.Equal(http.StatusBadRequest, res.StatusCode, "status code not correct")
	bodyAsErrorResponse := res.Body.(models.ErrorReponse)
	c.Len(bodyAsErrorResponse.Error.(apperrors.ValidationErrors), 2, "not correct errors returned")
}

func (c *GetCatalogsAWSLambdaTests) TestHandler_UnhandledError() {
	request := catalog.CatalogGetterRequest{
		CatalogNames: []string{"biotest", "roles"},
	}
	c.catalogGetter.On("Run", c.ctx, &request).Return(
		nil,
		c.ordinaryError,
	)

	res, err := c.getCatalogsAWSLambda.Handler(c.ctx, request)

	bodyAsErrorResponse := res.Body.(models.ErrorReponse)
	c.Nil(err, "should not return an error")
	c.Equal(http.StatusInternalServerError, res.StatusCode, "status code not correct")
	c.Equal(c.ordinaryError.Error(), bodyAsErrorResponse.Error, "not correct error returned")
}

func (c *GetCatalogsAWSLambdaTests) TestHandler_NoCatalogsFoundErr() {
	notExistCatalog := "no_exists"
	request := catalog.CatalogGetterRequest{
		CatalogNames: []string{"biotest", notExistCatalog},
	}
	errorReturned := catalog.NoCatalogFound{CatalogName: notExistCatalog}
	c.catalogGetter.On("Run", c.ctx, &request).Return(
		nil,
		errorReturned,
	)

	res, err := c.getCatalogsAWSLambda.Handler(c.ctx, request)

	bodyAsErrorResponse := res.Body.(models.ErrorReponse)
	c.Nil(err, "should not return an error")
	c.Equal(http.StatusNotFound, res.StatusCode, "status code not correct")
	c.Equal(errorReturned.Error(), bodyAsErrorResponse.Error, "not correct error returned")
}
