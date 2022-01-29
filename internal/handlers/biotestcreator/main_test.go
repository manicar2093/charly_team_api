package biotestcreator

import (
	"context"
	"errors"
	"testing"

	"github.com/go-rel/rel/reltest"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/internal/apperrors"
	"github.com/manicar2093/charly_team_api/internal/services"
	"github.com/manicar2093/charly_team_api/internal/validators"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(MainTests))
}

type MainTests struct {
	suite.Suite
	repo           *reltest.Repository
	validator      validators.MockValidatorService
	uuidGen        services.MockUUIDGenerator
	ctx            context.Context
	biotestUpdater *BiotestCreatorImpl
	ordinaryError  error
}

func (c *MainTests) SetupTest() {
	c.repo = reltest.New()
	c.validator = validators.MockValidatorService{}
	c.uuidGen = services.MockUUIDGenerator{}
	c.uuidGen.On("New").Return("an generated uuid")
	c.ctx = context.Background()
	c.biotestUpdater = NewBiotestCreator(c.repo, &c.validator, &c.uuidGen)
	c.ordinaryError = errors.New("An ordinary error :O")

}

func (c *MainTests) TearDownTest() {
	c.validator.AssertExpectations(c.T())
	c.repo.AssertExpectations(c.T())
}

func (c *MainTests) TestCreateNewBiotest() {
	biotestRequest := entities.Biotest{}
	c.validator.On("Validate", &biotestRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.repo.ExpectTransaction(func(r *reltest.Repository) {
		c.repo.ExpectInsert().ForType("entities.Biotest").Return(nil)
		c.repo.ExpectInsert().ForType("entities.HigherMuscleDensity").Return(nil)
		c.repo.ExpectInsert().ForType("entities.LowerMuscleDensity").Return(nil)
		c.repo.ExpectInsert().ForType("entities.SkinFolds").Return(nil)
	})

	res, err := c.biotestUpdater.Run(c.ctx, &biotestRequest)

	c.Nil(err)
	c.NotNil(res)
	c.NotEmpty(res.BiotestCreated.ID, "unexpected id biotest response")

}

func (c *MainTests) TestCreateNewBiotest_InsertError() {
	biotestRequest := entities.Biotest{}
	c.validator.On("Validate", &biotestRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.repo.ExpectTransaction(func(r *reltest.Repository) {
		c.repo.ExpectInsert().ForType("entities.Biotest").Error(c.ordinaryError)
		c.repo.ExpectInsert().ForType("entities.HigherMuscleDensity").Return(nil)
		c.repo.ExpectInsert().ForType("entities.LowerMuscleDensity").Return(nil)
		c.repo.ExpectInsert().ForType("entities.SkinFolds").Return(nil)
	})

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
