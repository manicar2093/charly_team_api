package services

import (
	"errors"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/manicar2093/charly_team_api/aws"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/connections"
	"github.com/manicar2093/charly_team_api/entities"
	"github.com/manicar2093/charly_team_api/models"
)

type RoleType int32

var (
	emailAttributeName string = "email"
)

const (
	ADMIN RoleType = iota + 1
	COACH
	CUSTOMER
)

type UserService interface {
	CreateUser(user models.CreateUserRequest) (int32, error)
}

// UserServiceCognito is a middleware to Cognito Services. PoolID is taken from config package.
type UserServiceCognito struct {
	provider aws.CongitoClient
	db       connections.Repository
	passGen  PassGen
}

func NewUserServiceCognito(
	provider aws.CongitoClient,
	db connections.Repository,
	passGen PassGen,
) *UserServiceCognito {
	return &UserServiceCognito{}
}

func (u UserServiceCognito) CreateUser(
	user models.CreateUserRequest,
) (int32, error) {

	userEntity := entities.User{
		Name:     user.Name,
		LastName: user.LastName,
		RoleID:   int32(user.RoleID),
		Email:    user.Email,
		Birthday: user.Birthday,
	}

	dbResponse := u.db.Save(&userEntity)

	if dbResponse.Error != nil {
		log.Println(dbResponse.Error)
		return 0, errors.New("error saving user into de db")
	}

	pass, err := u.passGen.Generate()
	if err != nil {
		log.Println(err)
		return 0, errors.New("error generating temporary password")
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

	_, err = u.provider.AdminCreateUser(&requestData)
	// TODO: Add aws error handling
	if err != nil {
		return 0, err
	}

	userEntity.IsCreated = true

	dbResponse = u.db.Save(&userEntity)
	if dbResponse.Error != nil {
		log.Println(dbResponse.Error)
		return 0, errors.New("error confirming user creation")
	}

	return userEntity.ID, nil

}

func (u UserServiceCognito) getUserUsername(email *string) *string {
	metadata := strings.Split(*email, "@")
	return &metadata[0]
}
