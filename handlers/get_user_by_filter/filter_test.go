package main

import (
	"context"
	"errors"
	"testing"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/filters"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserFilterTest struct {
	suite.Suite
	repo                         *reltest.Repository
	paginator                    *mocks.Paginable
	ctx                          context.Context
	filterParams                 filters.FilterParameters
	notFoundError, ordinaryError error
}

func (c *UserFilterTest) SetupTest() {
	c.repo = reltest.New()
	c.ctx = context.Background()
	c.ordinaryError = errors.New("An ordinary error :O")
	c.paginator = &mocks.Paginable{}
	c.filterParams = filters.FilterParameters{
		Ctx:       c.ctx,
		Repo:      c.repo,
		Paginator: c.paginator,
	}
	c.notFoundError = rel.NotFoundError{}

}

func (c *UserFilterTest) TearDownTest() {
	c.repo.AssertExpectations(c.T())
	c.paginator.AssertExpectations(c.T())
}

func (c *UserFilterTest) TestFilterUserByID() {

	userIDRequested := 1

	request := map[string]interface{}{
		"user_id": userIDRequested,
	}

	c.repo.ExpectFind(
		where.Eq("id", userIDRequested),
	).Result(
		entities.User{
			ID: int32(userIDRequested),
		},
	)

	c.filterParams.Values = request

	got, err := FindUserByID(&c.filterParams)

	c.Nil(err, "FindUserByID return an error")

	userGot, ok := got.(entities.User)
	c.True(ok, "unexpected answare type")
	c.Equal(userGot.ID, int32(userIDRequested), "unexpected user id")

}

func (c *UserFilterTest) TestFilterUserByIDValidatioError() {

	request := map[string]interface{}{}

	c.filterParams.Values = request

	_, err := FindUserByID(&c.filterParams)

	validationError, isValidationError := err.(apperrors.ValidationError)

	c.True(isValidationError, "bad type of error ")

	c.Equal(validationError.Validation, "required")
	c.Equal(validationError.Field, "user_id")

}

func (c *UserFilterTest) TestFilterUserByIDNotFound() {

	userIDRequested := 1

	request := map[string]interface{}{
		"user_id": userIDRequested,
	}

	c.repo.ExpectFind(
		where.Eq("id", userIDRequested),
	).Return(c.notFoundError)

	c.filterParams.Values = request

	_, err := FindUserByID(&c.filterParams)

	_, isHandableNotFoundError := err.(apperrors.UserNotFound)

	c.True(isHandableNotFoundError, "unexpected error type")

}

func (c *UserFilterTest) TestFilterUserByIDUnhandledError() {

	userIDRequested := 1

	request := map[string]interface{}{
		"user_id": userIDRequested,
	}

	c.repo.ExpectFind(
		where.Eq("id", userIDRequested),
	).Return(c.ordinaryError)

	c.filterParams.Values = request

	_, err := FindUserByID(&c.filterParams)

	c.NotNil(err, "should return error")

}

func (c *UserFilterTest) TestFindUserByEmail() {

	userEmailRequested := "testing@testing.com"

	request := map[string]interface{}{
		"email": userEmailRequested,
	}

	c.repo.ExpectFind(
		where.Like("email", "%"+userEmailRequested+"%"),
	).Result(
		entities.User{
			Email: userEmailRequested,
		},
	)

	c.filterParams.Values = request

	got, err := FindUserByEmail(&c.filterParams)

	c.Nil(err, "FindUserByID return an error")

	userGot, ok := got.(entities.User)
	c.True(ok, "unexpected answare type")
	c.Equal(userGot.Email, userEmailRequested, "unexpected user id")

}

func (c *UserFilterTest) TestFindUserByEmailValidationError() {

	request := map[string]interface{}{}

	c.filterParams.Values = request

	_, err := FindUserByEmail(&c.filterParams)

	_, isValidationError := err.(apperrors.ValidationError)
	c.True(isValidationError, "unexpected error type")

}

func (c *UserFilterTest) TestFindUserByEmailNotFoundError() {

	userEmailRequested := "testing@testing.com"

	request := map[string]interface{}{
		"email": userEmailRequested,
	}

	c.repo.ExpectFind(
		where.Like("email", "%"+userEmailRequested+"%"),
	).Return(c.notFoundError)

	c.filterParams.Values = request

	_, err := FindUserByEmail(&c.filterParams)

	_, isHandableNotFoundError := err.(apperrors.UserNotFound)

	c.True(isHandableNotFoundError, "unexpected error type")

}

func (c *UserFilterTest) TestFindUserByEmailUnhandledError() {

	userEmailRequested := "testing@testing.com"

	request := map[string]interface{}{
		"email": userEmailRequested,
	}

	c.repo.ExpectFind(
		where.Like("email", "%"+userEmailRequested+"%"),
	).Return(c.ordinaryError)

	c.filterParams.Values = request

	_, err := FindUserByEmail(&c.filterParams)

	c.NotNil(err, "should not return error")

}

func (c *UserFilterTest) TestFindAllUsers() {

	userPageRequested := 2

	request := map[string]interface{}{
		"page_number": userPageRequested,
	}

	c.paginator.On(
		"CreatePaginator",
		c.ctx,
		entities.UserTable,
		mock.Anything,
		userPageRequested,
	).Return(&models.Paginator{}, nil)

	c.filterParams.Values = request

	got, err := FindAllUsers(&c.filterParams)

	c.Nil(err, "FindUserByID return an error")

	_, ok := got.(*models.Paginator)
	c.True(ok, "unexpected answare type")

}

func (c *UserFilterTest) TestFindAllUsersValidationError() {

	request := map[string]interface{}{}

	c.filterParams.Values = request

	_, err := FindAllUsers(&c.filterParams)

	_, isValidationError := err.(apperrors.ValidationError)
	c.True(isValidationError, "unexpected error type")

}

func (c *UserFilterTest) TestNewUserFilterService() {

	userFilter := NewUserFilterService(c.repo, c.paginator)
	c.NotNil(userFilter, "user filter should not be nil")

	userFilterRunner := userFilter.GetFilter("find_user_by_id")

	c.True(userFilterRunner.IsFound(), "filter should be found")

	userFilterRunner = userFilter.GetFilter("does_not_exists")

	c.False(userFilterRunner.IsFound(), "filter should not be found")

}

func TestUserFilter(t *testing.T) {
	suite.Run(t, new(UserFilterTest))
}
