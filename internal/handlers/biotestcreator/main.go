package biotestcreator

import (
	"context"

	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/services"
	"github.com/manicar2093/charly_team_api/pkg/logger"
	"github.com/manicar2093/charly_team_api/pkg/validators"
)

type BiotestCreator interface {
	Run(ctx context.Context, req *entities.Biotest) (*BiotestCreatorResponse, error)
}

type BiotestCreatorImpl struct {
	biotestRepo repositories.BiotestRepository
	validator   validators.ValidatorService
	uuidGen     services.UUIDGenerator
}

func NewBiotestCreator(
	biotestRepo repositories.BiotestRepository,
	validator validators.ValidatorService,
	uuidGen services.UUIDGenerator,
) *BiotestCreatorImpl {
	return &BiotestCreatorImpl{
		biotestRepo: biotestRepo,
		validator:   validator,
		uuidGen:     uuidGen,
	}
}

func (c *BiotestCreatorImpl) Run(ctx context.Context, req *entities.Biotest) (*BiotestCreatorResponse, error) {
	logger.Info(req)
	validation := c.validator.Validate(req)

	if !validation.IsValid {
		logger.Error(validation.Err)
		return nil, validation.Err
	}

	err := c.biotestRepo.SaveBiotest(ctx, req)

	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return &BiotestCreatorResponse{BiotestCreated: req}, nil
}
