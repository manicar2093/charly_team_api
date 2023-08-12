package userfilters

import (
	"context"

	"github.com/manicar2093/health_records/internal/db/repositories"
	"github.com/manicar2093/health_records/pkg/logger"
	"github.com/manicar2093/health_records/pkg/validators"
)

type AllUsersFinder interface {
	Run(ctx context.Context, req *AllUsersFinderRequest) (*AllUsersFinderResponse, error)
}

type AllUsersFinderImpl struct {
	userRepo  repositories.UserRepository
	validator validators.ValidatorService
}

func NewAllUsersFinderImpl(
	userRepo repositories.UserRepository,
	validator validators.ValidatorService,
) *AllUsersFinderImpl {
	return &AllUsersFinderImpl{userRepo: userRepo, validator: validator}
}

func (c *AllUsersFinderImpl) Run(ctx context.Context, req *AllUsersFinderRequest) (*AllUsersFinderResponse, error) {
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
