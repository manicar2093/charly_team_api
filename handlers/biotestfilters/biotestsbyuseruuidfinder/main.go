package biotestsbyuseruuidfinder

import (
	"context"

	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/validators"
)

type BiotestByUserUUID interface {
	Run(ctx context.Context, req *BiotestByUserUUIDRequest) (*BiotestByUserUUIDResponse, error)
}

type biotestByUserUUIDImpl struct {
	biotestRepo repositories.BiotestRepository
	validator   validators.ValidatorService
}

func NewBiotestByUserUUIDImpl(biotestRepo repositories.BiotestRepository, validator validators.ValidatorService) *biotestByUserUUIDImpl {
	return &biotestByUserUUIDImpl{biotestRepo: biotestRepo, validator: validator}
}

func (c *biotestByUserUUIDImpl) Run(ctx context.Context, req *BiotestByUserUUIDRequest) (*BiotestByUserUUIDResponse, error) {
	validation := c.validator.Validate(req)

	if !validation.IsValid {
		return nil, validation.Err
	}

	if req.AsCatalog {
		biotests, err := c.biotestRepo.GetAllUserBiotestByUserUUIDAsCatalog(ctx, &req.PageSort, req.UserUUID)
		if err != nil {
			return nil, err
		}
		return &BiotestByUserUUIDResponse{FoundBiotests: biotests}, nil
	}

	biotests, err := c.biotestRepo.GetAllUserBiotestByUserUUID(ctx, &req.PageSort, req.UserUUID)
	if err != nil {
		return nil, err
	}

	return &BiotestByUserUUIDResponse{FoundBiotests: biotests}, nil

}
