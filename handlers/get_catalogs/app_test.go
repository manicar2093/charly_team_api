package main

import (
	"testing"
	"time"

	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AppTests struct {
	suite.Suite
	catalogRepo    mocks.CatalogRepository
	biotypesReturn []entities.Biotype
	rolesReturn    []entities.Role
}

func (c *AppTests) SetupTest() {
	c.catalogRepo = mocks.CatalogRepository{}
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

func (c *AppTests) TearDownTest() {
	c.catalogRepo.AssertExpectations(c.T())
}

func (c *AppTests) TestCatalogFactoryLoopNoCatalog() {

	catalogs := models.GetCatalogsRequest{
		CatalogNames: []string{"biotype", "no_exists"},
	}

	c.catalogRepo.On("FindAllCatalogItems", mock.Anything).Return(c.biotypesReturn, nil).Once()

	_, err := CatalogFactoryLoop(catalogs, &c.catalogRepo)

	c.NotNil(err, "should not be error")

	_, isNoCatalogError := err.(apperrors.NoCatalogFound)
	c.True(isNoCatalogError, "type error not correct")

}

func (c *AppTests) TestCatalogFactoryLoop() {

	catalogs := models.GetCatalogsRequest{
		CatalogNames: []string{"biotype", "roles"},
	}

	c.catalogRepo.On("FindAllCatalogItems", mock.Anything).Return(c.biotypesReturn, nil).Once()
	c.catalogRepo.On("FindAllCatalogItems", mock.Anything).Return(c.rolesReturn, nil).Once()

	got, err := CatalogFactoryLoop(catalogs, &c.catalogRepo)

	c.Nil(err, "should not be error")

	biotypes, ok := got["biotype"].([]entities.Biotype)
	c.True(ok, "error parsing biotypes")
	c.Greater(len(biotypes), 0, "no sufficent items received")

	roles, ok := got["roles"].([]entities.Role)
	c.True(ok, "error parsing biotypes")
	c.Greater(len(roles), 0, "no sufficent items received")

}

func TestApp(t *testing.T) {
	suite.Run(t, new(AppTests))
}
