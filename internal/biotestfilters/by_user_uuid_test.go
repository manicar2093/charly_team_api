package biotestfilters_test

import (
	"context"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/manicar2093/health_records/internal/biotestfilters"
	"github.com/manicar2093/health_records/internal/db/entities"
	"github.com/manicar2093/health_records/internal/db/paginator"
	"github.com/manicar2093/health_records/internal/db/repositories"
	"github.com/manicar2093/health_records/mocks"
	"github.com/manicar2093/health_records/pkg/apperrors"
	"github.com/manicar2093/health_records/pkg/validators"
	"github.com/stretchr/testify/suite"
)

func TestByUserUUID(t *testing.T) {
	suite.Run(t, new(BiotestByUserUUIDTests))
}

type BiotestByUserUUIDTests struct {
	suite.Suite
	ctx               context.Context
	biotestRepo       *mocks.BiotestRepository
	validator         *mocks.ValidatorService
	biotestByUserUUID *biotestfilters.BiotestByUserUUIDImpl
	fake              faker.Faker
}

func (c *BiotestByUserUUIDTests) SetupTest() {
	c.ctx = context.Background()
	c.biotestRepo = &mocks.BiotestRepository{}
	c.validator = &mocks.ValidatorService{}
	c.biotestByUserUUID = biotestfilters.NewBiotestByUserUUIDImpl(c.biotestRepo, c.validator)
	c.fake = faker.New()
}

func (c *BiotestByUserUUIDTests) TearDownTest() {
	c.biotestRepo.AssertExpectations(c.T())
	c.validator.AssertExpectations(c.T())
}

func (c *BiotestByUserUUIDTests) TestBiotestByUserUUID() {
	userUUID := c.fake.UUID().V4()
	paginatorResponse := paginator.Paginator{Data: &[]entities.Biotest{}}
	request := biotestfilters.BiotestByUserUUIDRequest{UserUUID: userUUID}
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{Err: nil, IsValid: true})
	c.biotestRepo.On("GetAllUserBiotestByUserUUID", c.ctx, &request.PageSort, userUUID).Return(&paginatorResponse, nil)

	got, err := c.biotestByUserUUID.Run(c.ctx, &request)

	c.Nil(err)
	c.NotNil(got)
	c.IsType(&paginator.Paginator{}, got.FoundBiotests)
	c.IsType(&[]entities.Biotest{}, got.FoundBiotests.Data)
}

func (c *BiotestByUserUUIDTests) TestBiotestByUserUUID_AsCatalog() {
	userUUID := c.fake.UUID().V4()
	paginatorResponse := paginator.Paginator{Data: &[]repositories.BiotestDetails{}}
	request := biotestfilters.BiotestByUserUUIDRequest{UserUUID: userUUID, AsCatalog: true}
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{Err: nil, IsValid: true})
	c.biotestRepo.On("GetAllUserBiotestByUserUUIDAsCatalog", c.ctx, &request.PageSort, userUUID).Return(&paginatorResponse, nil)

	got, err := c.biotestByUserUUID.Run(c.ctx, &request)

	c.Nil(err)
	c.NotNil(got)
	c.IsType(&paginator.Paginator{}, got.FoundBiotests)
	c.IsType(&[]repositories.BiotestDetails{}, got.FoundBiotests.Data)
}

func (c *BiotestByUserUUIDTests) TestBiotestByUserUUID_ValidationError() {
	userUUID := c.fake.UUID().V4()
	request := biotestfilters.BiotestByUserUUIDRequest{UserUUID: userUUID}
	validationErrors := apperrors.ValidationErrors{{Field: "uuid", Validation: "required"}}
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{Err: validationErrors, IsValid: false})

	got, err := c.biotestByUserUUID.Run(c.ctx, &request)

	c.NotNil(err)
	c.Nil(got)
	c.IsType(apperrors.ValidationErrors{}, err)
}
