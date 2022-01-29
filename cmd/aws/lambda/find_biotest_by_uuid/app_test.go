package main

import (
	"context"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/handlers/biotestfilters/biotestbyuuidfinder"
	"github.com/manicar2093/charly_team_api/internal/models"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(FindBiotestByUUIDAWSLambdaTest))
}

type FindBiotestByUUIDAWSLambdaTest struct {
	suite.Suite
	ctx                        context.Context
	biotestByUUIDDFinder       *biotestbyuuidfinder.MockBiotestByUUID
	findBiotestByUUIDAWSLambda *FindBiotestByUUIDAWSLambda
	faker                      faker.Faker
}

func (c *FindBiotestByUUIDAWSLambdaTest) SetupTest() {
	c.ctx = context.Background()
	c.biotestByUUIDDFinder = &biotestbyuuidfinder.MockBiotestByUUID{}
	c.findBiotestByUUIDAWSLambda = NewFindBiotestByUUIDAWSLambda(c.biotestByUUIDDFinder)
	c.faker = faker.New()
}

func (c *FindBiotestByUUIDAWSLambdaTest) TearDownTest() {
	c.biotestByUUIDDFinder.AssertExpectations(c.T())
}

func (c *FindBiotestByUUIDAWSLambdaTest) TestHandler() {
	biotestUUID := c.faker.UUID().V4()
	req := biotestbyuuidfinder.BiotestByUUIDRequest{UUID: biotestUUID}
	biotestFound := entities.Biotest{BiotestUUID: biotestUUID}
	biotestByUUIDFinderReturn := biotestbyuuidfinder.BiotestByUUIDResponse{Biotest: &biotestFound}
	c.biotestByUUIDDFinder.On("Run", c.ctx, &req).Return(&biotestByUUIDFinderReturn, nil)

	got, err := c.findBiotestByUUIDAWSLambda.Handler(c.ctx, req)

	c.Nil(err)
	c.NotNil(got)
	c.IsType(&entities.Biotest{}, got.Body)
}

func (c *FindBiotestByUUIDAWSLambdaTest) TestHandler_NotFound() {
	biotestUUID := c.faker.UUID().V4()
	req := biotestbyuuidfinder.BiotestByUUIDRequest{UUID: biotestUUID}
	c.biotestByUUIDDFinder.On("Run", c.ctx, &req).Return(nil, repositories.NotFoundError{Entity: "Biotest", Identifier: biotestUUID})

	got, err := c.findBiotestByUUIDAWSLambda.Handler(c.ctx, req)

	c.Nil(err)
	c.NotNil(got)
	c.IsType(models.ErrorReponse{}, got.Body)
}

func (c *FindBiotestByUUIDAWSLambdaTest) TestHandler_ValidationError() {
	biotestUUID := c.faker.UUID().V4()
	req := biotestbyuuidfinder.BiotestByUUIDRequest{UUID: biotestUUID}
	c.biotestByUUIDDFinder.On("Run", c.ctx, &req).Return(nil, repositories.NotFoundError{Entity: "Biotest", Identifier: biotestUUID})

	got, err := c.findBiotestByUUIDAWSLambda.Handler(c.ctx, req)

	c.Nil(err)
	c.NotNil(got)
	c.IsType(models.ErrorReponse{}, got.Body)
}
