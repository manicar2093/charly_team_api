package biotest_test

import (
	"context"
	"errors"
	"testing"

	"github.com/manicar2093/health_records/internal/biotest"
	"github.com/manicar2093/health_records/internal/db/entities"
	"github.com/manicar2093/health_records/mocks"
	"github.com/manicar2093/health_records/pkg/apperrors"
	"github.com/manicar2093/health_records/pkg/validators"
	"github.com/stretchr/testify/suite"
)

func TestCreator(t *testing.T) {
	suite.Run(t, new(MainTests))
}

type MainTests struct {
	suite.Suite
	biotestRepo    *mocks.BiotestRepository
	validator      *mocks.ValidatorService
	uuidGen        *mocks.UUIDGenerator
	ctx            context.Context
	biotestUpdater *biotest.BiotestCreatorImpl
	ordinaryError  error
}

func (c *MainTests) SetupTest() {
	c.biotestRepo = &mocks.BiotestRepository{}
	c.validator = &mocks.ValidatorService{}
	c.uuidGen = &mocks.UUIDGenerator{}
	c.uuidGen.On("New").Return("an generated uuid")
	c.ctx = context.Background()
	c.biotestUpdater = biotest.NewBiotestCreator(c.biotestRepo, c.validator, c.uuidGen)
	c.ordinaryError = errors.New("An ordinary error :O")

}

func (c *MainTests) TearDownTest() {
	c.validator.AssertExpectations(c.T())
	c.biotestRepo.AssertExpectations(c.T())
}

func (c *MainTests) TestCreateNewBiotest() {
	biotestRequest := entities.Biotest{}
	c.validator.On("Validate", &biotestRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.biotestRepo.On("SaveBiotest", c.ctx, &biotestRequest).Return(nil)

	res, err := c.biotestUpdater.Run(c.ctx, &biotestRequest)

	c.Nil(err)
	c.NotNil(res)

}

func (c *MainTests) TestCreateNewBiotest_InsertError() {
	biotestRequest := entities.Biotest{}
	c.validator.On("Validate", &biotestRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.biotestRepo.On("SaveBiotest", c.ctx, &biotestRequest).Return(c.ordinaryError)

	res, err := c.biotestUpdater.Run(c.ctx, &biotestRequest)

	c.NotNil(err)
	c.Nil(res)

}

func (c *MainTests) TestCreateNewBiotest_NoValidReq() {

	biotestRequest := entities.Biotest{}

	validationErrors := apperrors.ValidationErrors{
		{Field: "weight", Validation: "required"},
		{Field: "height", Validation: "required"},
	}

	c.validator.On("Validate", &biotestRequest).Return(validators.ValidateOutput{IsValid: false, Err: validationErrors})

	res, err := c.biotestUpdater.Run(c.ctx, &biotestRequest)

	c.NotNil(err)
	c.Nil(res)
	c.IsType(apperrors.ValidationErrors{}, err)
	c.Len(err.(apperrors.ValidationErrors), 2)

}
