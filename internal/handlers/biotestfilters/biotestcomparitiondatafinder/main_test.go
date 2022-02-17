package biotestcomparitiondatafinder

import (
	"context"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/manicar2093/charly_team_api/internal/db/repositories"
	"github.com/manicar2093/charly_team_api/pkg/apperrors"
	"github.com/manicar2093/charly_team_api/pkg/validators"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(BiotestComparitionDataFinderTest))
}

type BiotestComparitionDataFinderTest struct {
	suite.Suite
	ctx                          context.Context
	biotestRepo                  *repositories.MockBiotestRepository
	validator                    *validators.MockValidatorService
	biotestComparitionDataFinder *biotestComparitionDataFinderImpl
	faker                        faker.Faker
}

func (c *BiotestComparitionDataFinderTest) SetupTest() {
	c.ctx = context.Background()
	c.biotestRepo = &repositories.MockBiotestRepository{}
	c.validator = &validators.MockValidatorService{}
	c.biotestComparitionDataFinder = NewBiotestComparitionDataFinderImpl(c.biotestRepo, c.validator)
	c.faker = faker.New()
}

func (c *BiotestComparitionDataFinderTest) TearDownTest() {
	c.biotestRepo.AssertExpectations(c.T())
	c.validator.AssertExpectations(c.T())
}

func (c *BiotestComparitionDataFinderTest) TestHandler() {
	userUUID := c.faker.UUID().V4()
	request := BiotestComparitionDataFinderRequest{UserUUID: userUUID}
	comparitionDataReturn := repositories.BiotestComparisionResponse{}
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{Err: nil, IsValid: true})
	c.biotestRepo.On("GetComparitionDataByUserUUID", c.ctx, userUUID).Return(&comparitionDataReturn, nil)

	got, err := c.biotestComparitionDataFinder.Run(c.ctx, &request)

	c.Nil(err)
	c.NotNil(got)
	c.Equal(&comparitionDataReturn, got.ComparitionData)
}

func (c *BiotestComparitionDataFinderTest) TestHandler_ValidationError() {
	userUUID := c.faker.UUID().V4()
	request := BiotestComparitionDataFinderRequest{UserUUID: userUUID}
	validationErrors := apperrors.ValidationErrors{
		{Field: "user_uuid", Validation: "required"},
	}
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{Err: validationErrors, IsValid: false})

	got, err := c.biotestComparitionDataFinder.Run(c.ctx, &request)

	c.NotNil(err)
	c.Nil(got)
	c.IsType(apperrors.ValidationErrors{}, err)
	c.Equal(validationErrors, err)
}
