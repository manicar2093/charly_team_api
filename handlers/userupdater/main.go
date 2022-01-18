package userupdater

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/validators"
)

type UserUpdater interface {
	Run(ctx context.Context, userData *entities.User) (*UserUpdaterResponse, error)
}

type UserUpdaterImpl struct {
	repo      rel.Repository
	validator validators.ValidatorService
}

func NewUpdateUser(repo rel.Repository, validator validators.ValidatorService) *UserUpdaterImpl {
	return &UserUpdaterImpl{
		repo:      repo,
		validator: validator,
	}
}

func (c *UserUpdaterImpl) Run(ctx context.Context, userData *entities.User) (*UserUpdaterResponse, error) {

	if !validators.IsUpdateRequestValid(userData) {
		return nil, apperrors.ValidationErrors{{Field: "identifier", Validation: "required"}}
	}

	dataValidation := c.validator.Validate(userData)

	if !dataValidation.IsValid {
		return nil, dataValidation.Err
	}

	err := c.repo.Update(ctx, userData)
	if err != nil {
		return nil, err
	}

	return &UserUpdaterResponse{userData}, nil
}
