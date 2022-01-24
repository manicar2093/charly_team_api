package biotestimagessaver

import (
	"context"

	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/validators"
	"github.com/manicar2093/charly_team_api/validators/nullsql"
)

type BiotestImagesSaver interface {
	Run(ctx context.Context, biotestImages *BiotestImagesSaverRequest) (*BiotestImagesSaverResponse, error)
}

type biotestImagesSaverImpl struct {
	biotestRepo repositories.BiotestRepository
	validator   validators.ValidatorService
}

func NewBiotestImagesSaverImpl(biotestRepo repositories.BiotestRepository, validator validators.ValidatorService) *biotestImagesSaverImpl {
	return &biotestImagesSaverImpl{biotestRepo: biotestRepo, validator: validator}
}

func (c *biotestImagesSaverImpl) Run(ctx context.Context, biotestImages *BiotestImagesSaverRequest) (*BiotestImagesSaverResponse, error) {
	validation := c.validator.Validate(biotestImages)

	if !validation.IsValid {
		return nil, validation.Err
	}

	biotest, err := c.biotestRepo.FindBiotestByUUID(ctx, biotestImages.BiotestUUID)

	if err != nil {
		return nil, err
	}

	biotest.FrontPicture = nullsql.ValidateStringSQLValid(biotestImages.FrontPicture)
	biotest.BackPicture = nullsql.ValidateStringSQLValid(biotestImages.BackPicture)
	biotest.LeftSidePicture = nullsql.ValidateStringSQLValid(biotestImages.LeftSidePicture)
	biotest.RightSidePicture = nullsql.ValidateStringSQLValid(biotestImages.RightSidePicture)

	if err := c.biotestRepo.UpdateBiotest(ctx, biotest); err != nil {
		return nil, err
	}

	return &BiotestImagesSaverResponse{BiotestImagesSaved: biotestImages}, nil
}
