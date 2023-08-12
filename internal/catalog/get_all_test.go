package catalog_test

import (
	"context"
	"testing"
	"time"

	"github.com/manicar2093/health_records/internal/catalog"
	"github.com/manicar2093/health_records/internal/db/entities"
	"github.com/manicar2093/health_records/mocks"
	"github.com/manicar2093/health_records/pkg/apperrors"
	"github.com/manicar2093/health_records/pkg/validators"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MainTests struct {
	suite.Suite
	catalogRepo    *mocks.CatalogRepository
	validator      *mocks.ValidatorService
	catalogsGetter catalog.CatalogGetter
	ctx            context.Context
	biotypesReturn []entities.Biotype
	rolesReturn    []entities.Role
}

func (c *MainTests) SetupTest() {
	c.catalogRepo = &mocks.CatalogRepository{}
	c.validator = &mocks.ValidatorService{}
	c.ctx = context.Background()
	c.catalogsGetter = catalog.NewCatalogGetterImpl(c.catalogRepo, c.validator)
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

func (c *MainTests) TestGetCatalogs() {
	catalogs := catalog.CatalogGetterRequest{
		CatalogNames: []string{"biotype", "roles"},
	}
	c.validator.On("Validate", &catalogs).Return(validators.ValidateOutput{IsValid: true, Err: nil}).Once()
	c.catalogRepo.On("FindAllCatalogItems", c.ctx, mock.Anything).Return(c.biotypesReturn, nil).Once()
	c.catalogRepo.On("FindAllCatalogItems", c.ctx, mock.Anything).Return(c.rolesReturn, nil).Once()

	res, err := c.catalogsGetter.Run(c.ctx, &catalogs)

	c.Nil(err)
	c.NotNil(res)
	c.Equal(c.biotypesReturn, res.Catalogs["biotype"])
	c.Equal(c.rolesReturn, res.Catalogs["roles"])

}

func (c *MainTests) TestGetCatalogs_NotExists() {
	catalogs := catalog.CatalogGetterRequest{
		CatalogNames: []string{"biotype", "no_exists"},
	}
	c.validator.On("Validate", &catalogs).Return(validators.ValidateOutput{IsValid: true, Err: nil}).Once()
	c.catalogRepo.On("FindAllCatalogItems", c.ctx, mock.Anything).Return(nil, nil).Once()

	res, err := c.catalogsGetter.Run(c.ctx, &catalogs)

	c.NotNil(err)
	c.Nil(res)
	c.IsType(catalog.NoCatalogFound{}, err)

}

func (c *MainTests) TestGetCatalogs_ValidationError() {
	catalogs := catalog.CatalogGetterRequest{
		CatalogNames: []string{"biotype", "roles"},
	}
	c.validator.On("Validate", &catalogs).Return(
		validators.ValidateOutput{
			IsValid: false,
			Err: apperrors.ValidationErrors{
				{Field: "biotype", Validation: "required"},
			},
		},
	).Once()

	res, err := c.catalogsGetter.Run(c.ctx, &catalogs)

	c.NotNil(err)
	c.Nil(res)
	c.IsType(apperrors.ValidationErrors{}, err)
	c.Len(err.(apperrors.ValidationErrors), 1)

}

func TestMain(t *testing.T) {
	suite.Run(t, new(MainTests))
}
