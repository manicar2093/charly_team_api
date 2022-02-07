package user

import (
	"context"

	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/validators"
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
	if validation := c.validator.Validate(req); !validation.IsValid {
		return nil, validation.Err
	}

	user, err := c.userRepo.FindUserByUUID(ctx, req.UserUUID)
	if err != nil {
		return nil, err
	}

	user.AvatarUrl = req.AvatarURL

	if err := c.userRepo.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	return &AvatarUpdaterResponse{
		UserUpdated: user,
	}, nil
}
