package services

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/aws"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/models"
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
	) (int32, error)
}

// UserServiceCognito is a middleware to Cognito Services. PoolID is taken from config package.
type UserServiceCognito struct {
	provider aws.CongitoClient
	passGen  PassGen
	repo     rel.Repository
}

func NewUserServiceCognito(
	provider aws.CongitoClient,
	passGen PassGen,
	repo rel.Repository,
) UserService {
	return &UserServiceCognito{
		provider: provider,
		passGen:  passGen,
		repo:     repo,
	}
}

func (u UserServiceCognito) CreateUser(
	ctx context.Context,
	user *models.CreateUserRequest,
) (int32, error) {

	userEntity := entities.User{
		Name:     user.Name,
		LastName: user.LastName,
		RoleID:   int32(user.RoleID),
		Email:    user.Email,
		Birthday: user.Birthday,
	}

	err := u.repo.Transaction(ctx, func(ctx context.Context) error {

		pass, err := u.passGen.Generate()
		if err != nil {
			log.Println(err)
			return errorGeneratePass
		}

		requestData := cognitoidentityprovider.AdminCreateUserInput{
			UserPoolId:        &config.CognitoPoolID,
			Username:          u.getUserUsername(&user.Email),
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
		return 0, err
	}

	return userEntity.ID, err

}

func (u UserServiceCognito) getUserUsername(email *string) *string {
	metadata := strings.Split(*email, "@")
	return &metadata[0]
}
