package main

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/validators"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MainTests struct {
	suite.Suite
	catalogRepo    mocks.CatalogRepository
	validator      mocks.ValidatorService
	ctx            context.Context
	biotypesReturn []entities.Biotype
	rolesReturn    []entities.Role
}

func (c *MainTests) SetupTest() {
	c.catalogRepo = mocks.CatalogRepository{}
	c.validator = mocks.ValidatorService{}
	c.ctx = context.Background()
	c.biotypesReturn = []entities.Biotype{
		{ID: 1, Description: "biotype1", CreatedAt: time.Now()},
		{ID: 2, Description: "biotype2", CreatedAt: time.Now()},
		{ID: 3, Description: "biotype3", CreatedAt: time.Now()},
		{ID: 4, Description: "biotype4", CreatedAt: time.Now()},
	}
	c.rolesReturn = []entities.Role{
		{ID: 1, Description: "role1", CreatedAt: time.Now()},
		{ID: 2, Description: "role2", CreatedAt: time.Now()},
		{ID: 3, Description: "role3", CreatedAt: time.Now()},
		{ID: 4, Description: "role4", CreatedAt: time.Now()},
	}
}

func (c *MainTests) TearDownTest() {
	c.catalogRepo.AssertExpectations(c.T())
	c.validator.AssertExpectations(c.T())
}

func (c *MainTests) TestGetCatalogsNotExists() {

	catalogs := models.GetCatalogsRequest{
		CatalogNames: []string{"biotype", "no_exists"},
	}

	c.validator.On("Validate", catalogs).Return(validators.ValidateOutput{IsValid: true, Err: nil}).Once()
	c.catalogRepo.On("FindAllCatalogItems", c.ctx, mock.Anything).Return(nil, nil).Once()

	res, _ := CreateLambdaHandlerWDependencies(&c.catalogRepo, &c.validator)(c.ctx, catalogs)

	c.Equal(res.StatusCode, http.StatusNotFound, "http error is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusNotFound), "http error is not correct")

	bodyError := res.Body.(models.ErrorReponse)

	c.NotEmpty(bodyError.Error, "error message should not be empty")

}

func (c *MainTests) TestGetCatalogs() {

	catalogs := models.GetCatalogsRequest{
		CatalogNames: []string{"biotype", "roles"},
	}

	c.validator.On("Validate", catalogs).Return(validators.ValidateOutput{IsValid: true, Err: nil}).Once()
	c.catalogRepo.On("FindAllCatalogItems", c.ctx, mock.Anything).Return(c.biotypesReturn, nil).Once()
	c.catalogRepo.On("FindAllCatalogItems", c.ctx, mock.Anything).Return(c.rolesReturn, nil).Once()

	res, _ := CreateLambdaHandlerWDependencies(&c.catalogRepo, &c.validator)(c.ctx, catalogs)

	c.Equal(res.StatusCode, http.StatusOK, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusOK), "http status is not correct")

	bodyAsMap := res.Body.(map[string]interface{})

	biotypes, ok := bodyAsMap["biotype"].([]entities.Biotype)
	c.True(ok, "error parsing biotypes")
	c.Greater(len(biotypes), 0, "no sufficent items received")

	roles, ok := bodyAsMap["roles"].([]entities.Role)
	c.True(ok, "error parsing biotypes")
	c.Greater(len(roles), 0, "no sufficent items received")
}

func (c *MainTests) TestGetCatalogsValidationError() {

	catalogs := models.GetCatalogsRequest{
		CatalogNames: []string{"biotype", "roles"},
	}

	c.validator.On("Validate", catalogs).Return(
		validators.ValidateOutput{
			IsValid: false,
			Err: apperrors.ValidationErrors{
				{Field: "biotype", Validation: "required"},
			},
		},
	).Once()

	res, _ := CreateLambdaHandlerWDependencies(&c.catalogRepo, &c.validator)(c.ctx, catalogs)

	c.Equal(res.StatusCode, http.StatusBadRequest, "http status is not correct")
	c.Equal(res.Status, http.StatusText(http.StatusBadRequest), "http status is not correct")

	bodyAsMap := res.Body.(map[string]interface{})

	errMessage, ok := bodyAsMap["error"].(apperrors.ValidationErrors)
	c.True(ok, "error parsing error message")
	c.Equal(len(errMessage), 1, "error message should not be empty")

}

func TestMain(t *testing.T) {
	suite.Run(t, new(MainTests))
}
