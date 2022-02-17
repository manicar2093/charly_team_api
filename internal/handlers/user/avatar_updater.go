package user

import (
	"context"

	"github.com/manicar2093/charly_team_api/internal/db/entities"
	"github.com/manicar2093/charly_team_api/internal/db/repositories"
	"github.com/manicar2093/charly_team_api/pkg/logger"
	"github.com/manicar2093/charly_team_api/pkg/validators"
)

type AvatarUpdaterRequest struct {
	UserUUID  string `validation:"required" json:"user_uuid,omitempty"`
	AvatarURL string `validation:"required" json:"avatar_url,omitempty"`
}

type AvatarUpdaterResponse struct {
	UserUpdated *entities.User
}

type AvatarUpdater interface {
	Run(ctx context.Context, req *AvatarUpdaterRequest) (*AvatarUpdaterResponse, error)
}

type userAvatarUpdaterImpl struct {
	userRepo  repositories.UserRepository
	validator validators.ValidatorService
}

func NewUserAvatarUpdater(
	userRepo repositories.UserRepository,
	validator validators.ValidatorService) AvatarUpdater {
	return &userAvatarUpdaterImpl{userRepo: userRepo, validator: validator}
}

func (c *userAvatarUpdaterImpl) Run(ctx context.Context, req *AvatarUpdaterRequest) (*AvatarUpdaterResponse, error) {
	logger.Info(req)
	if validation := c.validator.Validate(req); !validation.IsValid {
		return nil, validation.Err
	}

	user, err := c.userRepo.FindUserByUUID(ctx, req.UserUUID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	user.AvatarUrl = req.AvatarURL

	if err := c.userRepo.UpdateUser(ctx, user); err != nil {
		logger.Error(err)
		return nil, err
	}

	return &AvatarUpdaterResponse{
		UserUpdated: user,
	}, nil
}
