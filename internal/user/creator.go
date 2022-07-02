package user

import (
	"context"
	"fmt"

	"github.com/manicar2093/charly_team_api/internal/config"
	"github.com/manicar2093/charly_team_api/internal/db/entities"
	"github.com/manicar2093/charly_team_api/internal/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/services"
	"github.com/manicar2093/charly_team_api/pkg/logger"
	"github.com/manicar2093/charly_team_api/pkg/validators"
	"github.com/manicar2093/charly_team_api/pkg/validators/nullsql"
)

var (
	emailAttributeName          string = "email"
	emailVerifiedAttributeName  string = "email_verified"
	emailVerifiedAttributeValue string = "true"
)

type UserCreator interface {
	Run(ctx context.Context, user *UserCreatorRequest) (*UserCreatorResponse, error)
}

type UserCreatorImpl struct {
	passGen    services.PassGen
	userRepo   repositories.UserRepository
	uuidGen    services.UUIDGenerator
	validator  validators.ValidatorService
	passHasher services.HashPassword
}

func NewUserCreatorImpl(
	passGen services.PassGen,
	userRepo repositories.UserRepository,
	uuidGen services.UUIDGenerator,
	validator validators.ValidatorService,
	passHasher services.HashPassword,
) *UserCreatorImpl {
	return &UserCreatorImpl{
		passGen:    passGen,
		userRepo:   userRepo,
		uuidGen:    uuidGen,
		validator:  validator,
		passHasher: passHasher,
	}
}

func (c *UserCreatorImpl) Run(ctx context.Context, user *UserCreatorRequest) (*UserCreatorResponse, error) {
	logger.Info(user)
	if err := c.isValidRequest(user); err != nil {
		logger.Error(err)
		return nil, err
	}

	pass, err := c.passGen.Generate()
	if err != nil {
		logger.Error(err)
		return nil, ErrGenerationPass
	}

	passDigested, err := c.passHasher.Digest(pass)
	if err != nil {
		return nil, err
	}

	userEntity := entities.User{
		Name:          user.Name,
		LastName:      user.LastName,
		RoleID:        int32(user.RoleID),
		GenderID:      nullsql.ValidateIntSQLValid(int64(user.GenderID)),
		Email:         user.Email,
		Password:      passDigested,
		Birthday:      user.Birthday,
		BiotypeID:     nullsql.ValidateIntSQLValid(int64(user.BiotypeID)),
		BoneDensityID: nullsql.ValidateIntSQLValid(int64(user.BoneDensityID)),
	}

	userEntity.UserUUID = c.uuidGen.New()
	userEntity.AvatarUrl = fmt.Sprintf("%s%s.svg", config.AvatarURLSrc, c.uuidGen.New())

	if err := c.userRepo.SaveUser(ctx, &userEntity); err != nil {
		logger.Error(err)
		return nil, ErrSavingUserDB
	}

	return &UserCreatorResponse{UserCreated: &userEntity}, nil
}

func (c *UserCreatorImpl) isValidRequest(createUserReq *UserCreatorRequest) error {
	if validation := c.validator.Validate(createUserReq); !validation.IsValid {
		return validation.Err
	}

	if createUserReq.RoleID == CustomerRole {
		if validation := c.validator.Validate(
			createUserReq.GetCustomerCreationValidations(),
		); !validation.IsValid {
			return validation.Err
		}
	}

	return nil

}
