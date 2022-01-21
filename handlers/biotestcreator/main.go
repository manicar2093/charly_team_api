package biotestcreator

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/services"
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
	validation := c.validator.Validate(req)

	if !validation.IsValid {
		return nil, validation.Err
	}

	err := c.repo.Transaction(ctx, func(ctx context.Context) error {
		req.BiotestUUID = c.uuidGen.New()

		if err := c.repo.Insert(ctx, &req.HigherMuscleDensity); err != nil {
			return err
		}
		req.HigherMuscleDensityID = req.HigherMuscleDensity.ID

		if err := c.repo.Insert(ctx, &req.LowerMuscleDensity); err != nil {
			return err
		}
		req.LowerMuscleDensityID = req.LowerMuscleDensity.ID

		if err := c.repo.Insert(ctx, &req.SkinFolds); err != nil {
			return err
		}
		req.SkinFoldsID = req.SkinFolds.ID

		if err := c.repo.Insert(ctx, req); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return &BiotestCreatorResponse{BiotestCreated: req}, nil
}
