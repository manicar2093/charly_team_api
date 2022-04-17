package biotest

import (
	"context"
	"errors"
	"testing"

	"github.com/go-rel/reltest"
	"github.com/manicar2093/charly_team_api/internal/db/entities"
	"github.com/manicar2093/charly_team_api/pkg/apperrors"
	"github.com/manicar2093/charly_team_api/pkg/validators"
	"github.com/stretchr/testify/suite"
)

func TestUpdater(t *testing.T) {
	suite.Run(t, new(MainUpdater))
}

type MainUpdater struct {
	suite.Suite
	repo           *reltest.Repository
	validator      *validators.MockValidatorService
	ctx            context.Context
	biotestUpdater *BiotestUpdaterImpl
	ordinaryError  error
}

func (c *MainUpdater) SetupTest() {
	c.repo = reltest.New()
	c.validator = &validators.MockValidatorService{}
	c.ctx = context.Background()
	c.biotestUpdater = NewBiotestUpdater(c.repo, c.validator)
	c.ordinaryError = errors.New("An ordinary error :O")

}

func (c *MainUpdater) TearDownTest() {
	c.validator.AssertExpectations(c.T())
	c.repo.AssertExpectations(c.T())
}

func (c *MainUpdater) TestUpdateBiotest() {
	biotestRequest := entities.Biotest{
		ID: 1,
	}
	c.validator.On("Validate", &biotestRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.repo.ExpectUpdate().ForType("entities.Biotest").Success()

	res, err := c.biotestUpdater.Run(c.ctx, &biotestRequest)

	c.Nil(err, "should not return an error")
	c.NotNil(res, "should return a response")
	c.Equal(res.BiotestUpdated, &biotestRequest, "response content not correct")

}

func (c *MainUpdater) TestUpdateBiotest_UpdateError() {
	biotestRequest := entities.Biotest{
		ID: 1,
	}
	c.validator.On("Validate", &biotestRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.repo.ExpectUpdate().ForType("entities.Biotest").Error(c.ordinaryError)

	res, err := c.biotestUpdater.Run(c.ctx, &biotestRequest)

	c.Nil(res, "should not return a response")
	c.NotNil(err, "should return an error")
	c.Equal(c.ordinaryError.Error(), err.Error())

}

func (c *MainUpdater) TestUpdateBiotest_NoBiotestID() {
	biotestRequest := entities.Biotest{}

	res, err := c.biotestUpdater.Run(c.ctx, &biotestRequest)

	c.Nil(res, "should not return data")
	bodyError := err.(apperrors.ValidationErrors)
	c.Equal("identifier", bodyError[0].Field, "validation error is not correct")
	c.Equal("required", bodyError[0].Validation, "validation error is not correct")

}

func (c *MainUpdater) TestUpdateBiotest_NoValidRequest() {
	biotestRequest := entities.Biotest{ID: 1}
	validationErrors := apperrors.ValidationErrors{
		{Field: "weight", Validation: "required"},
		{Field: "height", Validation: "required"},
	}
	c.validator.On("Validate", &biotestRequest).Return(
		validators.ValidateOutput{IsValid: false, Err: validationErrors},
	)

	res, err := c.biotestUpdater.Run(c.ctx, &biotestRequest)

	c.Nil(res, "should not return data")
	errorGot, ok := err.(apperrors.ValidationErrors)
	c.True(ok, "error parsing error message")
	c.Equal(len(errorGot), 2, "error message should not be empty")

}
