package repositories

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/go-rel/rel/reltest"
	"github.com/go-rel/rel/where"
	"github.com/jaswdr/faker"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/stretchr/testify/suite"
)

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositoryTest))
}

type UserRepositoryTest struct {
	suite.Suite
	paginator      *mocks.Paginable
	uuidGen        *mocks.UUIDGenerator
	repo           *reltest.Repository
	userRepository UserRepository
	ctx            context.Context
	faker          faker.Faker
}

func (c *UserRepositoryTest) SetupTest() {
	c.repo = reltest.New()
	c.paginator = &mocks.Paginable{}
	c.uuidGen = &mocks.UUIDGenerator{}
	c.userRepository = NewUserRepositoryRel(c.repo)
	c.ctx = context.TODO()
	c.faker = faker.New()
}

func (c *UserRepositoryTest) TearDownTest() {
	t := c.T()
	c.repo.AssertExpectations(t)
	c.paginator.AssertExpectations(t)
	c.uuidGen.AssertExpectations(t)
}

func (c *UserRepositoryTest) TestFilterUserByUUID() {
	expectedUserUUID := c.faker.UUID().V4()
	expectedUserID := c.faker.Int32()
	userReturned := entities.User{
		ID:       expectedUserID,
		UserUUID: expectedUserUUID,
	}
	c.repo.ExpectFind(
		where.Eq("user_uuid", expectedUserUUID),
	).Result(userReturned)

	got, err := c.userRepository.FindUserByUUID(c.ctx, expectedUserUUID)

	c.Nil(err, "should not return an error")
	c.Equal(expectedUserUUID, got.UserUUID, "userUUID is not correct")

}

func (c *UserRepositoryTest) TestFilterUserByUUID_NotFound() {
	expectedUserUUID := c.faker.UUID().V4()
	c.repo.ExpectFind(
		where.Eq("user_uuid", expectedUserUUID),
	).NotFound()

	got, err := c.userRepository.FindUserByUUID(c.ctx, expectedUserUUID)

	c.IsType(NotFoundError{}, err, "error is not the correct type")
	c.Contains(err.Error(), expectedUserUUID, "error should contain the used identifier")
	c.Nil(got, "user should has no data")
}

func (c *UserRepositoryTest) TestFilterUserByUUID_UnexpectedError() {
	expectedUserUUID := c.faker.UUID().V4()
	expectedError := fmt.Errorf("an generic error")
	c.repo.ExpectFind(
		where.Eq("user_uuid", expectedUserUUID),
	).Error(expectedError)

	got, err := c.userRepository.FindUserByUUID(c.ctx, expectedUserUUID)

	c.Equal(expectedError, err, "error is not the correct")
	c.Nil(got, "user should has no data")
}

func (c *UserRepositoryTest) TestFindUserLikeEmailOrNameOrLastName() {
	expectedSearchParam := "expectedSearchParam"
	expectSearchParamLower := strings.ToLower(expectedSearchParam)
	usersReturned := []entities.User{
		{},
		{},
		{},
		{},
	}
	expectedFilter := where.Like("LOWER(email)", "%"+expectSearchParamLower+"%").OrLike("LOWER(name)", "%"+expectSearchParamLower+"%").OrLike("LOWER(last_name)", "%"+expectSearchParamLower+"%")
	c.repo.ExpectFindAll(expectedFilter).Result(usersReturned)

	got, err := c.userRepository.FindUserLikeEmailOrNameOrLastName(c.ctx, expectedSearchParam)

	c.Nil(err, "should not return an error")
	c.NotNil(got, "should return data")
	c.Len(*got, len(usersReturned), "data len is incorrect")

}

func (c *UserRepositoryTest) TestFindUserLikeEmailOrNameOrLastName_UnexpectedError() {
	expectedSearchParam := "expectedSearchParam"
	expectSearchParamLower := strings.ToLower(expectedSearchParam)

	expectedFilter := where.Like("LOWER(email)", "%"+expectSearchParamLower+"%").OrLike("LOWER(name)", "%"+expectSearchParamLower+"%").OrLike("LOWER(last_name)", "%"+expectSearchParamLower+"%")
	c.repo.ExpectFindAll(expectedFilter).Error(fmt.Errorf("a generic error"))

	got, err := c.userRepository.FindUserLikeEmailOrNameOrLastName(c.ctx, expectedSearchParam)

	c.NotNil(err, "should return an error")
	c.Nil(got, "should not return data")

}

func (c *UserRepositoryTest) TestFindAllUsers() {}
