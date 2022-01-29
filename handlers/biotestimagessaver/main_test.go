package biotestimagessaver

import (
	"context"
	"testing"

	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/validators"
	"github.com/stretchr/testify/suite"
	"gopkg.in/guregu/null.v4"
)

func TestSaveBiotestImage(t *testing.T) {
	suite.Run(t, new(SaveBiotestImagesTest))
}

type SaveBiotestImagesTest struct {
	suite.Suite
	ctx                                                       context.Context
	validator                                                 *validators.MockValidatorService
	biotestRepo                                               *repositories.MockBiotestRepository
	biotestUUID, frontImage, backImage, leftImage, rightImage string
	biotestExpectedUpdateCall, biotestFound                   entities.Biotest
	biotestImagesSaver                                        BiotestImagesSaver
	request                                                   BiotestImagesSaverRequest
	validationOutput                                          validators.ValidateOutput
}

func (c *SaveBiotestImagesTest) SetupTest() {
	c.ctx = context.Background()
	c.validator = &validators.MockValidatorService{}
	c.biotestRepo = &repositories.MockBiotestRepository{}
	c.biotestUUID = "biotest_uuid"
	c.frontImage = "front/image/path"
	c.backImage = "back/image/path"
	c.leftImage = "left/image/path"
	c.rightImage = "right/image/path"
	c.request = BiotestImagesSaverRequest{
		BiotestUUID:      c.biotestUUID,
		FrontPicture:     c.frontImage,
		BackPicture:      c.backImage,
		LeftSidePicture:  c.leftImage,
		RightSidePicture: c.rightImage,
	}
	c.biotestImagesSaver = NewBiotestImagesSaverImpl(c.biotestRepo, c.validator)
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
