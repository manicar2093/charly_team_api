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
	"github.com/manicar2093/charly_team_api/db/paginator"
	"github.com/stretchr/testify/suite"
)

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositoryTest))
}

type UserRepositoryTest struct {
	suite.Suite
	paginator      *paginator.MockPaginable
	repo           *reltest.Repository
	userRepository UserRepository
	ctx            context.Context
	faker          faker.Faker
}

func (c *UserRepositoryTest) SetupTest() {
	c.repo = reltest.New()
	c.paginator = &paginator.MockPaginable{}
	c.userRepository = NewUserRepositoryRel(c.repo, c.paginator)
	c.ctx = context.TODO()
	c.faker = faker.New()
}

func (c *UserRepositoryTest) TearDownTest() {
	t := c.T()
	c.repo.AssertExpectations(t)
	c.paginator.AssertExpectations(t)
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

func (c *UserRepositoryTest) TestFindAllUsers() {
	pageSort := paginator.PageSort{
		Page:         1,
		ItemsPerPage: 10,
	}
	usersHolder := []entities.User{}
	paginationReturn := paginator.Paginator{Data: []entities.User{{}, {}}}
	c.paginator.On(
		"CreatePagination",
		c.ctx,
		entities.UserTable,
		&usersHolder,
		&pageSort,
	).Return(&paginationReturn, nil)

	got, err := c.userRepository.FindAllUsers(c.ctx, &pageSort)

	c.Nil(err, "should not return error")
	c.NotNil(got, "should return a paginator instance")
	c.IsType([]entities.User{}, got.Data, "data has the incorrect type")
}

func (c *UserRepositoryTest) TestFindAllUsers_CreatePaginationError() {
	pageSort := paginator.PageSort{
		Page:         1,
		ItemsPerPage: 10,
	}
	usersHolder := []entities.User{}
	c.paginator.On(
		"CreatePagination",
		c.ctx,
		entities.UserTable,
		&usersHolder,
		&pageSort,
	).Return(nil, fmt.Errorf("a generic error"))

	got, err := c.userRepository.FindAllUsers(c.ctx, &pageSort)

	c.NotNil(err, "should return error")
	c.Nil(got, "should not return a paginator instance")
}

func (c *UserRepositoryTest) TestSaveUser() {
	expectedUserUUID := c.faker.UUID().V4()
	expectedUserToSave := entities.User{
		UserUUID: expectedUserUUID,
	}

	c.repo.ExpectTransaction(func(r *reltest.Repository) {
		r.ExpectInsert().For(&expectedUserToSave)
	})

	err := c.userRepository.SaveUser(c.ctx, &expectedUserToSave)

	c.Nil(err, "should not return error")

}

func (c *UserRepositoryTest) TestSaveUser_UnexpectedError() {
	expectedUserUUID := c.faker.UUID().V4()
	expectedUserToSave := entities.User{
		UserUUID: expectedUserUUID,
	}
	c.repo.ExpectTransaction(func(r *reltest.Repository) {
		r.ExpectInsert().For(&expectedUserToSave).Error(
			fmt.Errorf("a generic error"),
		)
	})

	err := c.userRepository.SaveUser(c.ctx, &expectedUserToSave)

	c.NotNil(err, "should return error")
}

func (c *UserRepositoryTest) TestUpdateUser() {
	expectedUserUUID := c.faker.UUID().V4()
	expectedUserToUpdate := entities.User{
		UserUUID: expectedUserUUID,
	}
	c.repo.ExpectTransaction(func(r *reltest.Repository) {
		r.ExpectUpdate().For(&expectedUserToUpdate)
	})

	err := c.userRepository.UpdateUser(c.ctx, &expectedUserToUpdate)

	c.Nil(err, "should not return error")
}

func (c *UserRepositoryTest) TestUpdateUser_UnexpectedError() {
	expectedUserUUID := c.faker.UUID().V4()
	expectedUserToUpdate := entities.User{
		UserUUID: expectedUserUUID,
	}
	c.repo.ExpectTransaction(func(r *reltest.Repository) {
		r.ExpectUpdate().For(&expectedUserToUpdate).Error(
			fmt.Errorf("a generic error"),
		)
	})

	err := c.userRepository.UpdateUser(c.ctx, &expectedUserToUpdate)

	c.NotNil(err, "should return error")
}
