package biotest_test

import (
	"context"
	"testing"

	"github.com/manicar2093/health_records/internal/biotest"
	"github.com/manicar2093/health_records/internal/db/entities"
	"github.com/manicar2093/health_records/mocks"
	"github.com/manicar2093/health_records/pkg/validators"
	"github.com/stretchr/testify/suite"
	"gopkg.in/guregu/null.v4"
)

func TestSaveBiotestImage(t *testing.T) {
	suite.Run(t, new(SaveBiotestImagesTest))
}

type SaveBiotestImagesTest struct {
	suite.Suite
	ctx                                                       context.Context
	validator                                                 *mocks.ValidatorService
	biotestRepo                                               *mocks.BiotestRepository
	biotestUUID, frontImage, backImage, leftImage, rightImage string
	biotestExpectedUpdateCall, biotestFound                   entities.Biotest
	biotestImagesSaver                                        biotest.BiotestImagesSaver
	request                                                   biotest.BiotestImagesSaverRequest
	validationOutput                                          validators.ValidateOutput
}

func (c *SaveBiotestImagesTest) SetupTest() {
	c.ctx = context.Background()
	c.validator = &mocks.ValidatorService{}
	c.biotestRepo = &mocks.BiotestRepository{}
	c.biotestUUID = "biotest_uuid"
	c.frontImage = "front/image/path"
	c.backImage = "back/image/path"
	c.leftImage = "left/image/path"
	c.rightImage = "right/image/path"
	c.request = biotest.BiotestImagesSaverRequest{
		BiotestUUID:      c.biotestUUID,
		FrontPicture:     c.frontImage,
		BackPicture:      c.backImage,
		LeftSidePicture:  c.leftImage,
		RightSidePicture: c.rightImage,
	}
	c.biotestImagesSaver = biotest.NewBiotestImagesSaverImpl(c.biotestRepo, c.validator)
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
	c.biotestRepo.AssertExpectations(t)
}

func (c *SaveBiotestImagesTest) TestSaveBiotestImages() {
	c.validator.On("Validate", &c.request).Return(c.validationOutput)
	c.biotestRepo.On("FindBiotestByUUID", c.ctx, c.biotestUUID).Return(&c.biotestFound, nil)
	c.biotestRepo.On("UpdateBiotest", c.ctx, &c.biotestExpectedUpdateCall).Return(nil)

	got, err := c.biotestImagesSaver.Run(c.ctx, &c.request)

	c.Nil(err)
	c.NotNil(got)

}
