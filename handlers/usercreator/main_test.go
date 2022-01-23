package usercreator

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/jaswdr/faker"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/aws"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/services"

	"github.com/manicar2093/charly_team_api/validators"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	suite.Run(t, new(UserCreatorTests))
}

type UserCreatorTests struct {
	suite.Suite
	authProviderMock *aws.MockCongitoClient
	validator        validators.MockValidatorService
	passGenMock      *services.MockPassGen
	uuidGen          *services.MockUUIDGenerator
	userRepo         *repositories.MockUserRepository
	userCreator      *userCreatorImpl
	ctx              context.Context
	userCreated      *entities.User
	faker            faker.Faker
}

func (c *UserCreatorTests) SetupTest() {
	c.authProviderMock = &aws.MockCongitoClient{}
	c.validator = validators.MockValidatorService{}
	c.passGenMock = &services.MockPassGen{}
	c.uuidGen = &services.MockUUIDGenerator{}
	c.userRepo = &repositories.MockUserRepository{}
	c.userCreator = NewUserCreatorImpl(c.authProviderMock, c.passGenMock, c.userRepo, c.uuidGen, &c.validator)
	c.ctx = context.Background()
	c.userCreated = &entities.User{ID: int32(1)}
	c.faker = faker.New()

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
	userRequest := UserCreatorRequest{
		Name:     "testing",
		LastName: "main",
		Email:    "testing@main-func.com",
		Birthday: time.Now(),
		RoleID:   1,
	}
	c.validator.On("Validate", &userRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
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
	c.userRepo.On("SaveUser", c.ctx, mock.AnythingOfType("*entities.User")).Return(nil)

	res, err := c.userCreator.Run(c.ctx, &userRequest)

	c.Nil(err)
	c.NotNil(res)
	c.Equal(userUUID, res.UserCreated.UserUUID)
	c.Contains(res.UserCreated.AvatarUrl, userAvatar)

}

func (c *UserCreatorTests) TestUserCreator_Customer() {
	userUUID := c.faker.UUID().V4()
	userAvatar := c.faker.UUID().V4()
	passReturn := c.faker.Lorem().Word()
	userRequest := UserCreatorRequest{
		Name:          "testing",
		LastName:      "main",
		Email:         "testing@main-func.com",
		Birthday:      time.Now(),
		RoleID:        3,
		GenderID:      1,
		BoneDensityID: 1,
		BiotypeID:     1,
	}
	c.validator.On("Validate", &userRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.validator.On("Validate", userRequest.GetCustomerCreationValidations()).Return(validators.ValidateOutput{IsValid: true, Err: nil})
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
	c.userRepo.On("SaveUser", c.ctx, mock.AnythingOfType("*entities.User")).Return(nil)

	res, err := c.userCreator.Run(c.ctx, &userRequest)

	c.Nil(err)
	c.NotNil(res)
	c.Equal(userUUID, res.UserCreated.UserUUID)
	c.Contains(res.UserCreated.AvatarUrl, userAvatar)

}

func (c *UserCreatorTests) TestUserCreator_ValidationError() {
	userRequest := UserCreatorRequest{
		Name:     "testing",
		LastName: "main",
		Email:    "testing@main-func.com",
		Birthday: time.Now(),
		RoleID:   1,
	}
	validationErrors := apperrors.ValidationErrors{
		{Field: "name", Validation: "required"},
	}
	c.validator.On("Validate", &userRequest).Return(
		validators.ValidateOutput{IsValid: false, Err: validationErrors},
	)

	res, err := c.userCreator.Run(c.ctx, &userRequest)

	c.NotNil(err)
	c.Nil(res)

}

func (c *UserCreatorTests) TestUserCreator_PassGenError() {
	userRequest := UserCreatorRequest{
		Name:     "testing",
		LastName: "main",
		Email:    "testing@main-func.com",
		Birthday: time.Now(),
		RoleID:   1,
	}
	c.validator.On("Validate", &userRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.passGenMock.On("Generate").Return("", fmt.Errorf("unexpected error"))

	res, err := c.userCreator.Run(c.ctx, &userRequest)

	c.NotNil(err)
	c.Nil(res)
	c.Equal(generationPassError, err)

}

func (c *UserCreatorTests) TestUserCreator_AWSCognitoError() {
	passReturn := c.faker.Lorem().Word()
	userRequest := UserCreatorRequest{
		Name:     "testing",
		LastName: "main",
		Email:    "testing@main-func.com",
		Birthday: time.Now(),
		RoleID:   1,
	}
	c.validator.On("Validate", &userRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
	c.passGenMock.On("Generate").Return(passReturn, nil)
	c.authProviderMock.On(
		"AdminCreateUser",
		mock.AnythingOfType("*cognitoidentityprovider.AdminCreateUserInput"),
	).Return(
		nil,
		fmt.Errorf("unexpected error"),
	)

	res, err := c.userCreator.Run(c.ctx, &userRequest)

	c.NotNil(err)
	c.Nil(res)
	c.Equal(savingUserAWSError, err)

}

func (c *UserCreatorTests) TestUserCreator_SaveUserDBError() {
	userUUID := c.faker.UUID().V4()
	userAvatar := c.faker.UUID().V4()
	passReturn := c.faker.Lorem().Word()
	userRequest := UserCreatorRequest{
		Name:     "testing",
		LastName: "main",
		Email:    "testing@main-func.com",
		Birthday: time.Now(),
		RoleID:   1,
	}
	c.validator.On("Validate", &userRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
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

	res, err := c.userCreator.Run(c.ctx, &userRequest)

	c.NotNil(err)
	c.Nil(res)
	c.Equal(savingUserDBError, err)

}
