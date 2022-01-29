package biotestupdater

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/internal/logger"
	"github.com/manicar2093/charly_team_api/validators"
)

type BiotestUpdater interface {
	Run(ctx context.Context, req *entities.Biotest) (*BiotestUpdaterResponse, error)
}

type BiotestUpdaterImpl struct {
	repo      rel.Repository
	validator validators.ValidatorService
}

func NewBiotestUpdater(
	repo rel.Repository,
	validator validators.ValidatorService,
) *BiotestUpdaterImpl {
	return &BiotestUpdaterImpl{repo: repo, validator: validator}
}

func (c *BiotestUpdaterImpl) Run(ctx context.Context, req *entities.Biotest) (*BiotestUpdaterResponse, error) {
	logger.Info(req)
	if !validators.IsUpdateRequestValid(req) {
		logger.Error("identifier missed to continue")
		return nil, apperrors.ValidationErrors{{Field: "identifier", Validation: "required"}}
	}

	validation := c.validator.Validate(req)

	if !validation.IsValid {
		logger.Error(validation.Err)
		return nil, validation.Err
	}

	err := c.repo.Update(ctx, req)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &BiotestUpdaterResponse{BiotestUpdated: req}, nil
}
