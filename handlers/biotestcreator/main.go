package biotestcreator

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/internal/logger"
	"github.com/manicar2093/charly_team_api/internal/services"
	"github.com/manicar2093/charly_team_api/validators"
)

type BiotestCreator interface {
	Run(ctx context.Context, req *entities.Biotest) (*BiotestCreatorResponse, error)
}

type BiotestCreatorImpl struct {
	repo      rel.Repository
	validator validators.ValidatorService
	uuidGen   services.UUIDGenerator
}

func NewBiotestCreator(
	repo rel.Repository,
	validator validators.ValidatorService,
	uuidGen services.UUIDGenerator,
) *BiotestCreatorImpl {
	return &BiotestCreatorImpl{
		repo:      repo,
		validator: validator,
		uuidGen:   uuidGen,
	}
}

func (c *BiotestCreatorImpl) Run(ctx context.Context, req *entities.Biotest) (*BiotestCreatorResponse, error) {
	logger.Info(req)
	validation := c.validator.Validate(req)

	if !validation.IsValid {
		logger.Error(validation.Err)
		return nil, validation.Err
	}

	err := c.repo.Transaction(ctx, func(ctx context.Context) error {
		req.BiotestUUID = c.uuidGen.New()

		if err := c.repo.Insert(ctx, &req.HigherMuscleDensity); err != nil {
			logger.Error(err)
			return err
		}
		req.HigherMuscleDensityID = req.HigherMuscleDensity.ID

		if err := c.repo.Insert(ctx, &req.LowerMuscleDensity); err != nil {
			logger.Error(err)
			return err
		}
		req.LowerMuscleDensityID = req.LowerMuscleDensity.ID

		if err := c.repo.Insert(ctx, &req.SkinFolds); err != nil {
			logger.Error(err)
			return err
		}
		req.SkinFoldsID = req.SkinFolds.ID

		if err := c.repo.Insert(ctx, req); err != nil {
			logger.Error(err)
			return err
		}

		return nil
	})

	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return &BiotestCreatorResponse{BiotestCreated: req}, nil
}
