package biotestbyuuidfinder

import (
	"context"

	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/validators"
	"github.com/manicar2093/charly_team_api/pkg/logger"
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
