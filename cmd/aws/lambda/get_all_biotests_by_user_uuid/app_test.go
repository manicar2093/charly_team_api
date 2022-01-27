package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/paginator"
	"github.com/manicar2093/charly_team_api/handlers/biotestfilters/biotestsbyuseruuidfinder"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(GetAllBiotestByUserUUIDAWSLambdaTest))
}

type GetAllBiotestByUserUUIDAWSLambdaTest struct {
	suite.Suite
	ctx                               context.Context
	biotestsByUserUUIDFinder          *biotestsbyuseruuidfinder.MockBiotestByUserUUID
	getAllBiotestsByUserUUIDAWSLambda *GetAllBiotestByUserUUIDAWSLambda
	faker                             faker.Faker
}

func (c *GetAllBiotestByUserUUIDAWSLambdaTest) SetupTest() {
	c.ctx = context.Background()
	c.biotestsByUserUUIDFinder = &biotestsbyuseruuidfinder.MockBiotestByUserUUID{}
	c.getAllBiotestsByUserUUIDAWSLambda = NewGetAllBiotestByUserUUIDAWSLambda(c.biotestsByUserUUIDFinder)
	c.faker = faker.New()
}

func (c *GetAllBiotestByUserUUIDAWSLambdaTest) TearDownTest() {
	c.biotestsByUserUUIDFinder.AssertExpectations(c.T())
}

func (c *GetAllBiotestByUserUUIDAWSLambdaTest) TestsHandler() {
	userUUID := c.faker.UUID().V4()
	request := biotestsbyuseruuidfinder.BiotestByUserUUIDRequest{PageSort: paginator.PageSort{}, UserUUID: userUUID}
	pageReturned := paginator.Paginator{Data: &[]entities.Biotest{{}, {}}}
	response := biotestsbyuseruuidfinder.BiotestByUserUUIDResponse{FoundBiotests: &pageReturned}
	c.biotestsByUserUUIDFinder.On("Run", c.ctx, &request).Return(&response, nil)

	got, err := c.getAllBiotestsByUserUUIDAWSLambda.Handler(c.ctx, request)

	c.Nil(err)
	c.NotNil(got)
	c.Equal(http.StatusOK, got.StatusCode)
	c.IsType(&paginator.Paginator{}, got.Body)

}

func (c *GetAllBiotestByUserUUIDAWSLambdaTest) TestsHandler_ValidationError() {
	userUUID := c.faker.UUID().V4()
	request := biotestsbyuseruuidfinder.BiotestByUserUUIDRequest{PageSort: paginator.PageSort{}, UserUUID: userUUID}
	validationErrors := apperrors.ValidationErrors{
		{Field: "name", Validation: "required"},
		{Field: "last_name", Validation: "required"},
	}
	c.biotestsByUserUUIDFinder.On("Run", c.ctx, &request).Return(nil, validationErrors)

	got, err := c.getAllBiotestsByUserUUIDAWSLambda.Handler(c.ctx, request)

	c.Nil(err)
	c.NotNil(got)
	c.Equal(http.StatusBadRequest, got.StatusCode)
	c.IsType(models.ErrorReponse{}, got.Body)

}
