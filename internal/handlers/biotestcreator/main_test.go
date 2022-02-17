package biotestcreator

import (
	"context"
	"errors"
	"testing"

	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/services"
	"github.com/manicar2093/charly_team_api/pkg/apperrors"
	"github.com/manicar2093/charly_team_api/pkg/validators"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(MainTests))
}

type MainTests struct {
	suite.Suite
	biotestRepo    *repositories.MockBiotestRepository
	validator      validators.MockValidatorService
	uuidGen        services.MockUUIDGenerator
	ctx            context.Context
	biotestUpdater *BiotestCreatorImpl
	ordinaryError  error
}

func (c *MainTests) SetupTest() {
	c.biotestRepo = &repositories.MockBiotestRepository{}
	c.validator = validators.MockValidatorService{}
	c.uuidGen = services.MockUUIDGenerator{}
	c.uuidGen.On("New").Return("an generated uuid")
	c.ctx = context.Background()
	c.biotestUpdater = NewBiotestCreator(c.biotestRepo, &c.validator, &c.uuidGen)
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
