package biotestbyuuidfinder

import (
	"context"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/apperrors"
	"github.com/manicar2093/charly_team_api/pkg/validators"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(BiotestByUUIDTests))
}

type BiotestByUUIDTests struct {
	suite.Suite
	ctx           context.Context
	biotestRepo   *repositories.MockBiotestRepository
	validator     *validators.MockValidatorService
	biotestByUUID *biotestByUUIDImpl
	fake          faker.Faker
}

func (c *BiotestByUUIDTests) SetupTest() {
	c.ctx = context.Background()
	c.biotestRepo = &repositories.MockBiotestRepository{}
	c.validator = &validators.MockValidatorService{}
	c.biotestByUUID = NewBiotestByUUIDImpl(c.biotestRepo, c.validator)
	c.fake = faker.New()
}

func (c *BiotestByUUIDTests) TearDownTest() {
	c.biotestRepo.AssertExpectations(c.T())
	c.validator.AssertExpectations(c.T())
}

func (c *BiotestByUUIDTests) TestBiotestByUUID() {
	biotestUUID := c.fake.UUID().V4()
	biotestResponse := entities.Biotest{BiotestUUID: biotestUUID}
	request := BiotestByUUIDRequest{UUID: biotestUUID}
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{Err: nil, IsValid: true})
	c.biotestRepo.On("FindBiotestByUUID", c.ctx, biotestUUID).Return(&biotestResponse, nil)

	got, err := c.biotestByUUID.Run(c.ctx, &request)

	c.Nil(err)
	c.NotNil(got)
	c.Equal(biotestUUID, got.Biotest.BiotestUUID)
}

func (c *BiotestByUUIDTests) TestBiotestByUUID_NotFound() {
	biotestUUID := c.fake.UUID().V4()
	request := BiotestByUUIDRequest{UUID: biotestUUID}
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{Err: nil, IsValid: true})
	c.biotestRepo.On("FindBiotestByUUID", c.ctx, biotestUUID).Return(nil, repositories.NotFoundError{})

	got, err := c.biotestByUUID.Run(c.ctx, &request)

	c.NotNil(err)
	c.Nil(got)
	c.IsType(repositories.NotFoundError{}, err)
}

func (c *BiotestByUUIDTests) TestBiotestByUUID_ValidationError() {
	biotestUUID := c.fake.UUID().V4()
	request := BiotestByUUIDRequest{UUID: biotestUUID}
	validationErrors := apperrors.ValidationErrors{{Field: "uuid", Validation: "required"}}
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{Err: validationErrors, IsValid: false})

	got, err := c.biotestByUUID.Run(c.ctx, &request)

	c.NotNil(err)
	c.Nil(got)
	c.IsType(apperrors.ValidationErrors{}, err)
}
