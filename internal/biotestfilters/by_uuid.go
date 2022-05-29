package biotestfilters

import (
	"context"

	"github.com/manicar2093/charly_team_api/internal/db/repositories"
	"github.com/manicar2093/charly_team_api/pkg/logger"
	"github.com/manicar2093/charly_team_api/pkg/validators"
)

type BiotestByUUID interface {
	Run(ctx context.Context, req *BiotestByUUIDRequest) (*BiotestByUUIDResponse, error)
}

type BiotestByUUIDImpl struct {
	biotestRepo repositories.BiotestRepository
	validator   validators.ValidatorService
}

func NewBiotestByUUIDImpl(biotestRepo repositories.BiotestRepository, validator validators.ValidatorService) *BiotestByUUIDImpl {
	return &BiotestByUUIDImpl{biotestRepo: biotestRepo, validator: validator}
}

func (c *BiotestByUUIDImpl) Run(ctx context.Context, req *BiotestByUUIDRequest) (*BiotestByUUIDResponse, error) {
	logger.Info(req)
	validation := c.validator.Validate(req)

	if !validation.IsValid {
		logger.Error(validation.Err)
		return nil, validation.Err
	}

	biotest, err := c.biotestRepo.FindBiotestByUUID(ctx, req.UUID)

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &BiotestByUUIDResponse{Biotest: biotest}, nil
}
