package biotestsbyuseruuidfinder

import (
	"context"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/paginator"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/apperrors"
	"github.com/manicar2093/charly_team_api/validators"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(BiotestByUserUUIDTests))
}

type BiotestByUserUUIDTests struct {
	suite.Suite
	ctx               context.Context
	biotestRepo       *repositories.MockBiotestRepository
	validator         *validators.MockValidatorService
	biotestByUserUUID *biotestByUserUUIDImpl
	fake              faker.Faker
}

func (c *BiotestByUserUUIDTests) SetupTest() {
	c.ctx = context.Background()
	c.biotestRepo = &repositories.MockBiotestRepository{}
	c.validator = &validators.MockValidatorService{}
	c.biotestByUserUUID = NewBiotestByUserUUIDImpl(c.biotestRepo, c.validator)
	c.fake = faker.New()
}

func (c *BiotestByUserUUIDTests) TearDownTest() {
	c.biotestRepo.AssertExpectations(c.T())
	c.validator.AssertExpectations(c.T())
}

func (c *BiotestByUserUUIDTests) TestBiotestByUserUUID() {
	userUUID := c.fake.UUID().V4()
	paginatorResponse := paginator.Paginator{Data: &[]entities.Biotest{}}
	request := BiotestByUserUUIDRequest{UserUUID: userUUID}
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
	request := BiotestByUserUUIDRequest{UserUUID: userUUID, AsCatalog: true}
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
	request := BiotestByUserUUIDRequest{UserUUID: userUUID}
	validationErrors := apperrors.ValidationErrors{{Field: "uuid", Validation: "required"}}
	c.validator.On("Validate", &request).Return(validators.ValidateOutput{Err: validationErrors, IsValid: false})

	got, err := c.biotestByUserUUID.Run(c.ctx, &request)

	c.NotNil(err)
	c.Nil(got)
	c.IsType(apperrors.ValidationErrors{}, err)
}
