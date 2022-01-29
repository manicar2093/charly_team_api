package usercreator

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/aws"
	"github.com/manicar2093/charly_team_api/internal/config"
	"github.com/manicar2093/charly_team_api/internal/logger"
	"github.com/manicar2093/charly_team_api/services"
	"github.com/manicar2093/charly_team_api/validators"
	"github.com/manicar2093/charly_team_api/validators/nullsql"
)

type UserCreator interface {
	Run(ctx context.Context, user *UserCreatorRequest) (*UserCreatorResponse, error)
}

type userCreatorImpl struct {
	authProvider aws.CongitoClient
	passGen      services.PassGen
	userRepo     repositories.UserRepository
	uuidGen      services.UUIDGenerator
	validator    validators.ValidatorService
}

func NewUserCreatorImpl(
	authProvider aws.CongitoClient,
	passGen services.PassGen,
	userRepo repositories.UserRepository,
	uuidGen services.UUIDGenerator,
	validator validators.ValidatorService,
) *userCreatorImpl {
	return &userCreatorImpl{
		authProvider: authProvider,
		passGen:      passGen,
		userRepo:     userRepo,
		uuidGen:      uuidGen,
		validator:    validator,
	}
}

func (c *userCreatorImpl) Run(ctx context.Context, user *UserCreatorRequest) (*UserCreatorResponse, error) {
	logger.Info(user)
	if err := c.isValidRequest(user); err != nil {
		logger.Error(err)
		return nil, err
	}

	pass, err := c.passGen.Generate()
	if err != nil {
		logger.Error(err)
		return nil, errGenerationPass
	}

	requestData := cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:        &config.CognitoPoolID,
		Username:          &user.Email,
		TemporaryPassword: &pass,
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  &emailAttributeName,
				Value: &user.Email,
			},
		},
	}

	userOutput, err := c.authProvider.AdminCreateUser(&requestData)
	if err != nil {
		logger.Error(err)
		return nil, errSavingUserAWS
	}

	userEntity := entities.User{
		Name:          user.Name,
		LastName:      user.LastName,
		RoleID:        int32(user.RoleID),
		GenderID:      nullsql.ValidateIntSQLValid(int64(user.GenderID)),
		Email:         user.Email,
		Birthday:      user.Birthday,
		AvatarUrl:     fmt.Sprintf("%s%s.svg", config.AvatarURLSrc, c.uuidGen.New()),
		BiotypeID:     nullsql.ValidateIntSQLValid(int64(user.BiotypeID)),
		BoneDensityID: nullsql.ValidateIntSQLValid(int64(user.BoneDensityID)),
	}

	userEntity.UserUUID = *userOutput.User.Username

	err = c.userRepo.SaveUser(ctx, &userEntity)

	if err != nil {
		logger.Error(err)
		return nil, errSavingUserDB
	}

	return &UserCreatorResponse{UserCreated: &userEntity}, nil
}

func (c *userCreatorImpl) isValidRequest(createUserReq *UserCreatorRequest) error {
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
