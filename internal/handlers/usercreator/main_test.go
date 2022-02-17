package usercreator

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/jaswdr/faker"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/aws"
	"github.com/manicar2093/charly_team_api/internal/config"
	"github.com/manicar2093/charly_team_api/internal/services"
	"github.com/manicar2093/charly_team_api/pkg/apperrors"

	"github.com/manicar2093/charly_team_api/pkg/validators"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(UserCreatorTests))
}

type UserCreatorTests struct {
	suite.Suite
	authProviderMock                                                   *aws.MockCongitoClient
	validator                                                          validators.MockValidatorService
	passGenMock                                                        *services.MockPassGen
	uuidGen                                                            *services.MockUUIDGenerator
	userRepo                                                           *repositories.MockUserRepository
	userCreator                                                        *userCreatorImpl
	ctx                                                                context.Context
	userRequestBase                                                    UserCreatorRequest
	emailAttribute, emailVerifiedAttribute, emailVerifiedAttributValue string
	userEmail                                                          string
	faker                                                              faker.Faker
}

func (c *UserCreatorTests) SetupTest() {
	c.authProviderMock = &aws.MockCongitoClient{}
	c.validator = validators.MockValidatorService{}
	c.passGenMock = &services.MockPassGen{}
	c.uuidGen = &services.MockUUIDGenerator{}
	c.userRepo = &repositories.MockUserRepository{}
	c.userCreator = NewUserCreatorImpl(c.authProviderMock, c.passGenMock, c.userRepo, c.uuidGen, &c.validator)
	c.ctx = context.Background()
	c.faker = faker.New()
	c.userEmail = "testing@main-func.com"
	c.emailAttribute = "email"
	c.emailVerifiedAttribute = "email_verified"
	c.emailVerifiedAttributValue = "true"
	c.userRequestBase = UserCreatorRequest{
		Name:     "testing",
		LastName: "main",
		Email:    c.userEmail,
		Birthday: time.Now(),
		RoleID:   1,
	}

}

func (c *UserCreatorTests) TearDownTest() {
	T := c.T()
	c.authProviderMock.AssertExpectations(T)
	c.validator.AssertExpectations(T)
	c.passGenMock.AssertExpectations(T)
	c.uuidGen.AssertExpectations(T)
	c.userRepo.AssertExpectations(T)
}

func (c *UserCreatorTests) TestUserCreator_Admin() {
	userUUID := c.faker.UUID().V4()
	userAvatar := c.faker.UUID().V4()
	passReturn := c.faker.Lorem().Word()
	c.validator.On("Validate", &c.userRequestBase).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.passGenMock.On("Generate").Return(passReturn, nil)
	adminUserCreateReturn := cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:        &config.CognitoPoolID,
		Username:          &c.userEmail,
		TemporaryPassword: &passReturn,
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{Name: &c.emailAttribute, Value: &c.userEmail},
			{Name: &c.emailVerifiedAttribute, Value: &c.emailVerifiedAttributValue},
		},
	}
	c.authProviderMock.On(
		"AdminCreateUser",
		&adminUserCreateReturn,
	).Return(
		&cognitoidentityprovider.AdminCreateUserOutput{
			User: &cognitoidentityprovider.UserType{
				Username: &userUUID,
			},
		},
		nil,
	)
	c.uuidGen.On("New").Return(userAvatar)
	c.userRepo.On("SaveUser", c.ctx, mock.AnythingOfType("*entities.User")).Return(nil)

	res, err := c.userCreator.Run(c.ctx, &c.userRequestBase)

	c.Nil(err)
	c.NotNil(res)
	c.Equal(userUUID, res.UserCreated.UserUUID)
	c.Contains(res.UserCreated.AvatarUrl, userAvatar)

}

func (c *UserCreatorTests) TestUserCreator_Customer() {
	userUUID := c.faker.UUID().V4()
	userAvatar := c.faker.UUID().V4()
	passReturn := c.faker.Lorem().Word()
	c.userRequestBase.RoleID = 3
	c.userRequestBase.GenderID = 1
	c.userRequestBase.BoneDensityID = 1
	c.userRequestBase.BiotypeID = 1
	c.validator.On("Validate", &c.userRequestBase).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.validator.On("Validate", c.userRequestBase.GetCustomerCreationValidations()).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.passGenMock.On("Generate").Return(passReturn, nil)
	adminUserCreateReturn := cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:        &config.CognitoPoolID,
		Username:          &c.userEmail,
		TemporaryPassword: &passReturn,
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{Name: &c.emailAttribute, Value: &c.userEmail},
			{Name: &c.emailVerifiedAttribute, Value: &c.emailVerifiedAttributValue},
		},
	}
	c.authProviderMock.On(
		"AdminCreateUser",
		&adminUserCreateReturn,
	).Return(
		&cognitoidentityprovider.AdminCreateUserOutput{
			User: &cognitoidentityprovider.UserType{
				Username: &userUUID,
			},
		},
		nil,
	)
	c.uuidGen.On("New").Return(userAvatar)
	c.userRepo.On("SaveUser", c.ctx, mock.AnythingOfType("*entities.User")).Return(nil)

	res, err := c.userCreator.Run(c.ctx, &c.userRequestBase)

	c.Nil(err)
	c.NotNil(res)
	c.Equal(userUUID, res.UserCreated.UserUUID)
	c.Contains(res.UserCreated.AvatarUrl, userAvatar)

}

func (c *UserCreatorTests) TestUserCreator_ValidationError() {
	validationErrors := apperrors.ValidationErrors{
		{Field: "name", Validation: "required"},
	}
	c.validator.On("Validate", &c.userRequestBase).Return(
		validators.ValidateOutput{IsValid: false, Err: validationErrors},
	)

	res, err := c.userCreator.Run(c.ctx, &c.userRequestBase)

	c.NotNil(err)
	c.Nil(res)

}

func (c *UserCreatorTests) TestUserCreator_PassGenError() {
	c.validator.On("Validate", &c.userRequestBase).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.passGenMock.On("Generate").Return("", fmt.Errorf("unexpected error"))

	res, err := c.userCreator.Run(c.ctx, &c.userRequestBase)

	c.NotNil(err)
	c.Nil(res)
	c.Equal(errGenerationPass, err)

}

func (c *UserCreatorTests) TestUserCreator_AWSCognitoError() {
	passReturn := c.faker.Lorem().Word()
	c.validator.On("Validate", &c.userRequestBase).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.passGenMock.On("Generate").Return(passReturn, nil)
	c.authProviderMock.On(
		"AdminCreateUser",
		mock.AnythingOfType("*cognitoidentityprovider.AdminCreateUserInput"),
	).Return(
		nil,
		fmt.Errorf("unexpected error"),
	)

	res, err := c.userCreator.Run(c.ctx, &c.userRequestBase)

	c.NotNil(err)
	c.Nil(res)
	c.Equal(errSavingUserAWS, err)

}

func (c *UserCreatorTests) TestUserCreator_SaveUserDBError() {
	userUUID := c.faker.UUID().V4()
	userAvatar := c.faker.UUID().V4()
	passReturn := c.faker.Lorem().Word()
	c.validator.On("Validate", &c.userRequestBase).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.passGenMock.On("Generate").Return(passReturn, nil)
	c.authProviderMock.On(
		"AdminCreateUser",
		mock.AnythingOfType("*cognitoidentityprovider.AdminCreateUserInput"),
	).Return(
		&cognitoidentityprovider.AdminCreateUserOutput{
			User: &cognitoidentityprovider.UserType{
				Username: &userUUID,
			},
		},
		nil,
	)
	c.uuidGen.On("New").Return(userAvatar)
	c.userRepo.On("SaveUser", c.ctx, mock.AnythingOfType("*entities.User")).Return(fmt.Errorf("unexpected error"))

	res, err := c.userCreator.Run(c.ctx, &c.userRequestBase)

	c.NotNil(err)
	c.Nil(res)
	c.Equal(errSavingUserDB, err)

}
