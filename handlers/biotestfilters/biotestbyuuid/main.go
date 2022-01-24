package biotestbyuuid

import (
	"context"

	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/validators"
)

type BiotestByUUID interface {
	Run(ctx context.Context, req *BiotestByUUIDRequest) (*BiotestByUUIDResponse, error)
}

type biotestByUUIDImpl struct {
	biotestRepo repositories.BiotestRepository
	validator   validators.ValidatorService
}

func NewBiotestByUUIDImpl(biotestRepo repositories.BiotestRepository, validator validators.ValidatorService) *biotestByUUIDImpl {
	return &biotestByUUIDImpl{biotestRepo: biotestRepo, validator: validator}
}

func (c *biotestByUUIDImpl) Run(ctx context.Context, req *BiotestByUUIDRequest) (*BiotestByUUIDResponse, error) {
	validation := c.validator.Validate(req)

	if !validation.IsValid {
		return nil, validation.Err
	}

	biotest, err := c.biotestRepo.FindBiotestByUUID(ctx, req.UUID)

	if err != nil {
		return nil, err
	}

	return &BiotestByUUIDResponse{Biotest: biotest}, nil
}
