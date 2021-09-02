package services

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/go-rel/rel/reltest"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserServiceTest struct {
	suite.Suite
	providerMock          *mocks.CongitoClient
	passGenMock           *mocks.PassGen
	repoMock              *reltest.Repository
	username              string
	temporaryPass         string
	name, lastName, email string
	idUserCreated         int32
	birthday              time.Time
	userRequest           models.CreateUserRequest
	anError               error
	saveFuncMock          func(*entities.User) func(args mock.Arguments)
}

func (u *UserServiceTest) SetupTest() {
	u.providerMock = &mocks.CongitoClient{}
	u.repoMock = reltest.New()
	u.passGenMock = &mocks.PassGen{}
	u.username = "testing"
	u.temporaryPass = "12345678"
	u.anError = errors.New("An error")
	u.saveFuncMock = func(userDBReq *entities.User) func(args mock.Arguments) {
		return func(args mock.Arguments) {
			user := args[0].(*entities.User)
			user.ID = u.idUserCreated

			userDBReq.ID = u.idUserCreated
		}
	}

	u.name = "testing"
	u.lastName = "testing"
	u.email = strings.Join([]string{u.username, "@gmail.com"}, "")
	u.birthday = time.Date(1993, time.August, 20, 0, 0, 0, 0, time.UTC)
	u.idUserCreated = 1

	u.userRequest = models.CreateUserRequest{
		Name:     u.name,
		LastName: u.lastName,
		Email:    u.email,
		Birthday: u.birthday,
		RoleID:   3,
		GenderID: 1,
	}
}

func (u *UserServiceTest) TearDownTest() {
	t := u.T()
	u.repoMock.AssertExpectations(t)
	u.passGenMock.AssertExpectations(t)
	u.providerMock.AssertExpectations(t)
}

func (u *UserServiceTest) TestCreateUser() {

	adminCreateUserReq := cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:        &config.CognitoPoolID,
		Username:          &u.username,
		TemporaryPassword: &u.temporaryPass,
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  &emailAttributeName,
				Value: &u.userRequest.Email,
			},
		},
	}
	userDBReq := entities.User{
		Name:     u.userRequest.Name,
		LastName: u.userRequest.LastName,
		RoleID:   int32(u.userRequest.RoleID),
		Email:    u.userRequest.Email,
		Birthday: u.userRequest.Birthday,
		GenderID: int32(1),
	}

	u.providerMock.On(
		"AdminCreateUser",
		&adminCreateUserReq,
	).Return(
		&cognitoidentityprovider.AdminCreateUserOutput{},
		nil,
	)
	u.repoMock.ExpectTransaction(func(r *reltest.Repository) {
		u.passGenMock.On("Generate").Return(u.temporaryPass, nil)
		r.ExpectInsert().For(&userDBReq)
	})

	userService := NewUserServiceCognito(u.providerMock, u.passGenMock, u.repoMock)

	userCreated, err := userService.CreateUser(context.Background(), &u.userRequest)

	u.Nil(err)
	u.Equal(u.idUserCreated, userCreated, "user id is not correct")

}

func (u *UserServiceTest) TestCreateUserRepoSaveErr() {

	userDBReq := entities.User{
		Name:     u.userRequest.Name,
		LastName: u.userRequest.LastName,
		RoleID:   int32(u.userRequest.RoleID),
		Email:    u.userRequest.Email,
		Birthday: u.userRequest.Birthday,
		GenderID: int32(1),
	}

	u.repoMock.ExpectTransaction(func(r *reltest.Repository) {
		u.passGenMock.On("Generate").Return(u.temporaryPass, nil)
		r.ExpectInsert().For(&userDBReq).Return(u.anError)
	})
	userService := NewUserServiceCognito(u.providerMock, u.passGenMock, u.repoMock)

	userCreated, err := userService.CreateUser(context.Background(), &u.userRequest)

	u.NotNil(err, "should return an error")
	u.Empty(userCreated, "user was created")

}

func (u *UserServiceTest) TestCreateUserPassGenError() {

	u.repoMock.ExpectTransaction(func(r *reltest.Repository) {
		u.passGenMock.On("Generate").Return("", u.anError).Once()
	})

	userService := NewUserServiceCognito(u.providerMock, u.passGenMock, u.repoMock)

	userGot, err := userService.CreateUser(context.Background(), &u.userRequest)

	u.NotNil(err)
	u.Empty(userGot, "user should not be created")
}

func (u *UserServiceTest) TestCreateUserAdminCreateUserError() {

	adminCreateUserReq := cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:        &config.CognitoPoolID,
		Username:          &u.username,
		TemporaryPassword: &u.temporaryPass,
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  &emailAttributeName,
				Value: &u.userRequest.Email,
			},
		},
	}
	userDBReq := entities.User{
		Name:     u.userRequest.Name,
		LastName: u.userRequest.LastName,
		RoleID:   int32(u.userRequest.RoleID),
		Email:    u.userRequest.Email,
		Birthday: u.userRequest.Birthday,
		GenderID: int32(1),
	}

	u.providerMock.On(
		"AdminCreateUser",
		&adminCreateUserReq,
	).Return(
		&cognitoidentityprovider.AdminCreateUserOutput{},
		u.anError,
	)
	u.repoMock.ExpectTransaction(func(r *reltest.Repository) {
		u.passGenMock.On("Generate").Return(u.temporaryPass, nil)
		r.ExpectInsert().For(&userDBReq)
	})

	userService := NewUserServiceCognito(u.providerMock, u.passGenMock, u.repoMock)

	userCreated, err := userService.CreateUser(context.Background(), &u.userRequest)

	u.NotNil(err)
	u.Empty(userCreated, "user id is not correct")

}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTest))
}