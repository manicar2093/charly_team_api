package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/aws"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/services"
	"github.com/manicar2093/charly_team_api/validators/nullsql"
)

type RoleType int32

var (
	emailAttributeName string = "email"
	errorGeneratePass         = errors.New("error generating temporary password")
	errorSavingUser           = errors.New("error saving user into de db")
)

type UserService interface {
	CreateUser(
		ctx context.Context,
		user *models.CreateUserRequest,
	) (*entities.User, error)
}

// UserServiceCognito is a middleware to Cognito Services. PoolID is taken from config package.
type UserServiceCognito struct {
	provider aws.CongitoClient
	passGen  services.PassGen
	repo     rel.Repository
	uuidGen  services.UUIDGenerator
}

func NewUserServiceCognito(
	provider aws.CongitoClient,
	passGen services.PassGen,
	repo rel.Repository,
	uuidGen services.UUIDGenerator,
) UserService {
	return &UserServiceCognito{
		provider: provider,
		passGen:  passGen,
		repo:     repo,
		uuidGen:  uuidGen,
	}
}

func (u UserServiceCognito) CreateUser(
	ctx context.Context,
	user *models.CreateUserRequest,
) (*entities.User, error) {

	var userEntity entities.User

	err := u.repo.Transaction(ctx, func(ctx context.Context) error {
		userEntity = entities.User{
			Name:          user.Name,
			LastName:      user.LastName,
			RoleID:        int32(user.RoleID),
			GenderID:      nullsql.ValidateIntSQLValid(int64(user.GenderID)),
			Email:         user.Email,
			Birthday:      user.Birthday,
			AvatarUrl:     fmt.Sprintf("%s%s.svg", config.AvatarURLSrc, u.uuidGen.New()),
			BiotypeID:     nullsql.ValidateIntSQLValid(int64(user.BiotypeID)),
			BoneDensityID: nullsql.ValidateIntSQLValid(int64(user.BoneDensityID)),
		}

		pass, err := u.passGen.Generate()
		if err != nil {
			log.Println(err)
			return errorGeneratePass
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

		userOutput, err := u.provider.AdminCreateUser(&requestData)
		if err != nil {
			log.Println(err)
			return err
		}

		userEntity.UserUUID = *userOutput.User.Username
		log.Println(userEntity)

		err = u.repo.Insert(ctx, &userEntity)

		if err != nil {
			log.Println(err)
			return errorSavingUser
		}

		return nil
	})

	if err != nil {
		log.Println(err)
		return &entities.User{}, err
	}

	return &userEntity, err

}
