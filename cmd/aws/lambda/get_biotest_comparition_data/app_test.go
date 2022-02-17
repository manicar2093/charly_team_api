package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/manicar2093/charly_team_api/internal/db/entities"
	"github.com/manicar2093/charly_team_api/internal/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/handlers/biotestfilters"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(GetBiotestComparitionDataAWSLambdaTests))
}

type GetBiotestComparitionDataAWSLambdaTests struct {
	suite.Suite
	ctx                                context.Context
	biotestComparitionDataFinder       *biotestfilters.MockBiotestComparitionDataFinder
	getBiotestComparitionDataAWSLambda *GetBiotestComparitionDataAWSLambda
	faker                              faker.Faker
}

func (c *GetBiotestComparitionDataAWSLambdaTests) SetupTest() {
	c.ctx = context.Background()
	c.biotestComparitionDataFinder = &biotestfilters.MockBiotestComparitionDataFinder{}
	c.getBiotestComparitionDataAWSLambda = NewGetBiotestComparitionDataAWSLambda(c.biotestComparitionDataFinder)
	c.faker = faker.New()
}

func (c *GetBiotestComparitionDataAWSLambdaTests) TearDownTest() {
	c.biotestComparitionDataFinder.AssertExpectations(c.T())
}

func (c *GetBiotestComparitionDataAWSLambdaTests) TestHandler() {
	userUUID := c.faker.UUID().V4()
	request := biotestfilters.BiotestComparitionDataFinderRequest{UserUUID: userUUID}
	biotestComparitionDataReturn := biotestfilters.BiotestComparitionDataFinderResponse{ComparitionData: &repositories.BiotestComparisionResponse{FirstBiotest: &entities.Biotest{}, LastBiotest: &entities.Biotest{}, AllBiotestsDetails: &[]repositories.BiotestDetails{}}}
	c.biotestComparitionDataFinder.On("Run", c.ctx, &request).Return(
		&biotestComparitionDataReturn,
		nil,
	)

	got, err := c.getBiotestComparitionDataAWSLambda.Handler(c.ctx, request)

	c.Nil(err)
	c.NotNil(got)
}

func (c *GetBiotestComparitionDataAWSLambdaTests) TestHandler_UnhandledError() {
	userUUID := c.faker.UUID().V4()
	request := biotestfilters.BiotestComparitionDataFinderRequest{UserUUID: userUUID}
	ordinaryError := fmt.Errorf("ordinary error")
	c.biotestComparitionDataFinder.On("Run", c.ctx, &request).Return(
		nil,
		ordinaryError,
	)

	got, err := c.getBiotestComparitionDataAWSLambda.Handler(c.ctx, request)

	c.Nil(err)
	c.NotNil(got)
}
