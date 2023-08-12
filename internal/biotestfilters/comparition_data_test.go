package biotestfilters_test

import (
	"context"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/manicar2093/health_records/internal/biotestfilters"
	"github.com/manicar2093/health_records/internal/db/repositories"
	"github.com/manicar2093/health_records/mocks"
	"github.com/manicar2093/health_records/pkg/apperrors"
	"github.com/manicar2093/health_records/pkg/validators"
	"github.com/stretchr/testify/suite"
)

func TestComparitionData(t *testing.T) {
	suite.Run(t, new(BiotestComparitionDataFinderTest))
}

type BiotestComparitionDataFinderTest struct {
	suite.Suite
	ctx                          context.Context
	biotestRepo                  *mocks.BiotestRepository
	validator                    *mocks.ValidatorService
	biotestComparitionDataFinder *biotestfilters.BiotestComparitionDataFinderImpl
	faker                        faker.Faker
}

func (c *BiotestComparitionDataFinderTest) SetupTest() {
	c.ctx = context.Background()
	c.biotestRepo = &mocks.BiotestRepository{}
	c.validator = &mocks.ValidatorService{}
	c.biotestComparitionDataFinder = biotestfilters.NewBiotestComparitionDataFinderImpl(c.biotestRepo, c.validator)
	c.faker = faker.New()
}

func (c *BiotestComparitionDataFinderTest) TearDownTest() {
	c.biotestRepo.AssertExpectations(c.T())
	c.validator.AssertExpectations(c.T())
}

func (c *BiotestComparitionDataFinderTest) TestHandler() {
	userUUID := c.faker.UUID().V4()
	request := biotestfilters.BiotestComparitionDataFinderRequest{UserUUID: userUUID}
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
	request := biotestfilters.BiotestComparitionDataFinderRequest{UserUUID: userUUID}
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
