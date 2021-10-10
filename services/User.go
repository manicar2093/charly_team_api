package services

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
	"gopkg.in/guregu/null.v4"
)

type RoleType int32

var (
	emailAttributeName string = "email"
	errorGeneratePass         = errors.New("error generating temporary password")
	errorSavingUser           = errors.New("error saving user into de db")
)

const (
	ADMIN RoleType = iota + 1
	COACH
	CUSTOMER
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
	passGen  PassGen
	repo     rel.Repository
	uuidGen  UUIDGenerator
}

func NewUserServiceCognito(
	provider aws.CongitoClient,
	passGen PassGen,
	repo rel.Repository,
	uuidGen UUIDGenerator,
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

	userEntity := entities.User{
		Name:          user.Name,
		LastName:      user.LastName,
		RoleID:        int32(user.RoleID),
		GenderID:      null.IntFrom(int64(user.GenderID)),
		UserUUID:      u.uuidGen.New(),
		Email:         user.Email,
		Birthday:      user.Birthday,
		AvatarUrl:     fmt.Sprintf("%s%s", config.AvatarURLSrc, u.uuidGen.New()),
		BiotypeID:     null.IntFrom(int64(user.BiotypeID)),
		BoneDensityID: null.IntFrom(int64(user.BoneDensityID)),
	}

	err := u.repo.Transaction(ctx, func(ctx context.Context) error {

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

		err = u.repo.Insert(ctx, &userEntity)

		if err != nil {
			log.Println(err)
			return errorSavingUser
		}

		_, err = u.provider.AdminCreateUser(&requestData)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return &entities.User{}, err
	}

	return &userEntity, err

}
