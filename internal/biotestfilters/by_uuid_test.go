package biotestfilters_test

import (
	"context"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/manicar2093/health_records/internal/biotestfilters"
	"github.com/manicar2093/health_records/internal/db/entities"
	"github.com/manicar2093/health_records/internal/db/repositories"
	"github.com/manicar2093/health_records/mocks"
	"github.com/manicar2093/health_records/pkg/apperrors"
	"github.com/manicar2093/health_records/pkg/validators"
	"github.com/stretchr/testify/suite"
)

func TestByUUID(t *testing.T) {
	suite.Run(t, new(BiotestByUUIDTests))
}

type BiotestByUUIDTests struct {
	suite.Suite
	ctx           context.Context
	biotestRepo   *mocks.BiotestRepository
	validator     *mocks.ValidatorService
	biotestByUUID *biotestfilters.BiotestByUUIDImpl
	fake          faker.Faker
}

func (c *BiotestByUUIDTests) SetupTest() {
	c.ctx = context.Background()
	c.biotestRepo = &mocks.BiotestRepository{}
	c.validator = &mocks.ValidatorService{}
	c.biotestByUUID = biotestfilters.NewBiotestByUUIDImpl(c.biotestRepo, c.validator)
	c.fake = faker.New()
}

func (c *BiotestByUUIDTests) TearDownTest() {
	c.biotestRepo.AssertExpectations(c.T())
	c.validator.AssertExpectations(c.T())
}

func (c *BiotestByUUIDTests) TestBiotestByUUID() {
	biotestUUID := c.fake.UUID().V4()
	biotestResponse := entities.Biotest{BiotestUUID: biotestUUID}
	request := biotestfilters.BiotestByUUIDRequest{UUID: biotestUUID}
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{Err: nil, IsValid: true})
	c.biotestRepo.On("FindBiotestByUUID", c.ctx, biotestUUID).Return(&biotestResponse, nil)

	got, err := c.biotestByUUID.Run(c.ctx, &request)

	c.Nil(err)
	c.NotNil(got)
	c.Equal(biotestUUID, got.Biotest.BiotestUUID)
}

func (c *BiotestByUUIDTests) TestBiotestByUUID_NotFound() {
	biotestUUID := c.fake.UUID().V4()
	request := biotestfilters.BiotestByUUIDRequest{UUID: biotestUUID}
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{Err: nil, IsValid: true})
	c.biotestRepo.On("FindBiotestByUUID", c.ctx, biotestUUID).Return(nil, repositories.NotFoundError{})

	got, err := c.biotestByUUID.Run(c.ctx, &request)

	c.NotNil(err)
	c.Nil(got)
	c.IsType(repositories.NotFoundError{}, err)
}

func (c *BiotestByUUIDTests) TestBiotestByUUID_ValidationError() {
	biotestUUID := c.fake.UUID().V4()
	request := biotestfilters.BiotestByUUIDRequest{UUID: biotestUUID}
	validationErrors := apperrors.ValidationErrors{{Field: "uuid", Validation: "required"}}
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{Err: validationErrors, IsValid: false})

	got, err := c.biotestByUUID.Run(c.ctx, &request)

	c.NotNil(err)
	c.Nil(got)
	c.IsType(apperrors.ValidationErrors{}, err)
}
