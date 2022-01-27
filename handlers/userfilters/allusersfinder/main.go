package allusersfinder

import (
	"context"

	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/validators"
)

type AllUsersFinder interface {
	Run(ctx context.Context, req *AllUsersFinderRequest) (*AllUsersFinderResponse, error)
}

type allUsersFinderImpl struct {
	userRepo  repositories.UserRepository
	validator validators.ValidatorService
}

func NewAllUsersFinderImpl(
	userRepo repositories.UserRepository,
	validator validators.ValidatorService,
) *allUsersFinderImpl {
	return &allUsersFinderImpl{userRepo: userRepo, validator: validator}
}

func (c *allUsersFinderImpl) Run(ctx context.Context, req *AllUsersFinderRequest) (*AllUsersFinderResponse, error) {
	if validation := c.validator.Validate(req); !validation.IsValid {
		return nil, validation.Err
	}

	usersPage, err := c.userRepo.FindAllUsers(ctx, &req.PageSort)

	if err != nil {
		return nil, err
	}

	return &AllUsersFinderResponse{UsersFound: usersPage}, nil
}
