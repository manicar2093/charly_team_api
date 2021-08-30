package services

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
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
	repoMock              *mocks.UserRepository
	passGenMock           *mocks.PassGen
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
	u.repoMock = &mocks.UserRepository{}
	u.passGenMock = &mocks.PassGen{}
	u.username = "testing"
	u.temporaryPass = "12345678"
	u.anError = errors.New("An error")
	u.saveFuncMock = func(userDBReq *entities.User) func(args mock.Arguments) {
		return func(args mock.Arguments) {
			user := args[0].(*entities.User)
			user.ID = u.idUserCreated
			user.IsCreated = true

			userDBReq.ID = u.idUserCreated
			userDBReq.IsCreated = true
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
	}

	u.providerMock.On(
		"AdminCreateUser",
		&adminCreateUserReq,
	).Return(
		&cognitoidentityprovider.AdminCreateUserOutput{},
		nil,
	)
	u.passGenMock.On("Generate").Return(u.temporaryPass, nil)
	u.repoMock.On("Save", &userDBReq).Run(u.saveFuncMock(&userDBReq)).Return(nil)
	u.repoMock.On("Save", &userDBReq).Return(nil)

	userService := NewUserServiceCognito(u.providerMock, u.repoMock, u.passGenMock)

	userCreated, err := userService.CreateUser(&u.userRequest)

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
	}

	u.passGenMock.On("Generate").Return(u.temporaryPass, nil)
	u.repoMock.On("Save", &userDBReq).Run(u.saveFuncMock(&userDBReq)).Return(u.anError).Once()

	userService := NewUserServiceCognito(u.providerMock, u.repoMock, u.passGenMock)

	userCreated, err := userService.CreateUser(&u.userRequest)

	u.NotNil(err, "should return an error")
	u.Equal(userCreated, int32(0), "user id is not correct")

}

func (u *UserServiceTest) TestCreateUserPassGenError() {

	u.passGenMock.On("Generate").Return("", u.anError).Once()

	userService := NewUserServiceCognito(u.providerMock, u.repoMock, u.passGenMock)

	userGot, err := userService.CreateUser(&u.userRequest)

	u.NotNil(err)
	u.Empty(userGot, "user should not be created")
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTest))
}
