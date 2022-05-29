package userfilters

import (
	"context"

	"github.com/manicar2093/charly_team_api/internal/db/repositories"
	"github.com/manicar2093/charly_team_api/pkg/logger"
	"github.com/manicar2093/charly_team_api/pkg/validators"
)

type UserLikeEmailOrNameFinder interface {
	Run(ctx context.Context, req *UserLikeEmailOrNameFinderRequest) (*UserLikeEmailOrNameFinderResponse, error)
}

type UserLikeEmailOrNameFinderImpl struct {
	userRepo  repositories.UserRepository
	validator validators.ValidatorService
}

func NewUserLikeEmailOrNameFinderImpl(
	userRepo repositories.UserRepository,
	validator validators.ValidatorService,
) *UserLikeEmailOrNameFinderImpl {
	return &UserLikeEmailOrNameFinderImpl{userRepo: userRepo, validator: validator}
}

func (c *UserLikeEmailOrNameFinderImpl) Run(
	ctx context.Context,
	req *UserLikeEmailOrNameFinderRequest,
) (*UserLikeEmailOrNameFinderResponse, error) {
	logger.Info(req)
	if validation := c.validator.Validate(req); !validation.IsValid {
		logger.Error(validation.Err)
		return nil, validation.Err
	}

	usersFound, err := c.userRepo.FindUserLikeEmailOrNameOrLastName(ctx, req.FilterData)

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &UserLikeEmailOrNameFinderResponse{FetchedData: usersFound}, nil
}
