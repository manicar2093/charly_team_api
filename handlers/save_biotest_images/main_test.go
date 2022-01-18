package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-rel/rel/reltest"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/validators"
	"github.com/stretchr/testify/suite"
	"gopkg.in/guregu/null.v4"
)

func TestSaveBiotestImage(t *testing.T) {
	suite.Run(t, new(SaveBiotestImagesTest))
}

type SaveBiotestImagesTest struct {
	suite.Suite
	validator                                                 *validators.MockValidatorService
	repo                                                      reltest.Repository
	biotestUUID, frontImage, backImage, leftImage, rightImage string
	biotestExpectedUpdateCall, biotestFound                   entities.Biotest
	request                                                   BiotestImagesRequest
	validationOutput                                          validators.ValidateOutput
}

func (c *SaveBiotestImagesTest) SetupTest() {
	c.validator = &validators.MockValidatorService{}
	c.repo = *reltest.New()
	c.biotestUUID = "biotest_uuid"
	c.frontImage = "front/image/path"
	c.backImage = "back/image/path"
	c.leftImage = "left/image/path"
	c.rightImage = "right/image/path"
	c.request = BiotestImagesRequest{
		BiotestUUID:      c.biotestUUID,
		FrontPicture:     c.frontImage,
		BackPicture:      c.backImage,
		LeftSidePicture:  c.leftImage,
		RightSidePicture: c.rightImage,
	}

	c.biotestFound = entities.Biotest{
		BiotestUUID: c.biotestUUID,
	}

	c.biotestExpectedUpdateCall = entities.Biotest{
		BiotestUUID:      c.biotestUUID,
		FrontPicture:     null.StringFromPtr(&c.frontImage),
		BackPicture:      null.StringFromPtr(&c.backImage),
		LeftSidePicture:  null.StringFromPtr(&c.leftImage),
		RightSidePicture: null.StringFromPtr(&c.rightImage),
	}

	c.validationOutput = validators.ValidateOutput{IsValid: true, Err: nil}
}

func (c *SaveBiotestImagesTest) TearDownTest() {
	t := c.T()

	c.validator.AssertExpectations(t)
	c.repo.AssertExpectations(t)
}

func (c *SaveBiotestImagesTest) TestSaveBiotestImages() {
	c.validator.On("Validate", c.request).Return(c.validationOutput)
	c.repo.ExpectTransaction(func(r *reltest.Repository) {
		r.ExpectFind(where.Eq("biotest_uuid", c.biotestUUID))
		r.ExpectUpdate().ForType("entities.Biotest")
	})

	got, err := CreateLambdaHandlerWDependencies(c.validator, &c.repo)(context.TODO(), c.request)

	c.Nil(err, "should not get an error")
	c.Equal(http.StatusOK, got.StatusCode, "wrong status code")
	c.Nil(got.Body)

}
