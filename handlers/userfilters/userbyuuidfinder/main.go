package userbyuuidfinder

import (
	"context"

	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/validators"
)

type UserByUUIDFinder interface {
	Run(ctx context.Context, req *UserByUUIDFinderRequest) (*UserByUUIDFinderResponse, error)
}

type userByUUIDFinderImpl struct {
	userRepo  repositories.UserRepository
	validator validators.ValidatorService
}

func NewUserByUUIDFinderImpl(
	userRepo repositories.UserRepository,
	validator validators.ValidatorService,
) *userByUUIDFinderImpl {
	return &userByUUIDFinderImpl{userRepo: userRepo, validator: validator}
}

func (c *userByUUIDFinderImpl) Run(ctx context.Context, req *UserByUUIDFinderRequest) (*UserByUUIDFinderResponse, error) {
	if validation := c.validator.Validate(req); !validation.IsValid {
		return nil, validation.Err
	}

	user, err := c.userRepo.FindUserByUUID(ctx, req.UserUUID)

	if err != nil {
		return nil, err
	}

	return &UserByUUIDFinderResponse{UserFound: user}, nil
}
