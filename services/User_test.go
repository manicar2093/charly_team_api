package services

import (
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/entities"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type UserServiceTest struct {
	suite.Suite
	providerMock          *mocks.CongitoClient
	dbMock                *mocks.Repository
	passGenMock           *mocks.PassGen
	username              string
	temporaryPass         string
	name, lastName, email string
	birthday              time.Time
	userRequest           models.CreateUserRequest
}

func (u *UserServiceTest) SetupTest() {
	u.providerMock = &mocks.CongitoClient{}
	u.dbMock = &mocks.Repository{}
	u.passGenMock = &mocks.PassGen{}
	u.username = "testing"
	u.temporaryPass = "12345678"

	u.name = "testing"
	u.lastName = "testing"
	u.email = strings.Join([]string{u.username, "@gmail.com"}, "")
	u.birthday = time.Date(1993, time.August, 20, 0, 0, 0, 0, time.UTC)

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
	u.dbMock.AssertExpectations(t)
	u.passGenMock.AssertExpectations(t)
	u.providerMock.AssertExpectations(t)
}

func (u *UserServiceTest) TestCreateUser() {

	dbReturn := &gorm.DB{
		Error: nil,
	}

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
	saveFuncMock := func(args mock.Arguments) {
		user := args[0].(*entities.User)
		user.ID = 1
		user.IsCreated = true

		userDBReq.ID = 1
		userDBReq.IsCreated = true
	}

	u.providerMock.On(
		"AdminCreateUser",
		&adminCreateUserReq,
	).Return(
		&cognitoidentityprovider.AdminCreateUserOutput{},
		nil,
	)
	u.passGenMock.On("Generate").Return(u.temporaryPass, nil)
	u.dbMock.On("Save", &userDBReq).Run(saveFuncMock).Return(dbReturn)
	u.dbMock.On("Save", &userDBReq).Return(dbReturn)

	userService := UserServiceCognito{
		provider: u.providerMock,
		db:       u.dbMock,
		passGen:  u.passGenMock,
	}

	err := userService.CreateUser(u.userRequest)

	if err != nil {
		u.T().Fatal(err)
	}
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTest))
}
