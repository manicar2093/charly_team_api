package userfilters

import (
	"context"

	"github.com/manicar2093/charly_team_api/internal/db/repositories"
	"github.com/manicar2093/charly_team_api/pkg/logger"
	"github.com/manicar2093/charly_team_api/pkg/validators"
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
	logger.Info(req)
	if validation := c.validator.Validate(req); !validation.IsValid {
		logger.Error(validation.Err)
		return nil, validation.Err
	}

	usersPage, err := c.userRepo.FindAllUsers(ctx, &req.PageSort)

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &AllUsersFinderResponse{UsersFound: usersPage}, nil
}
