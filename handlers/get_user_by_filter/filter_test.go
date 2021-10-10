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

func (c *UserFilterTest) TestFilterUserByUUID() {

	userUUIDRequested := "an-uuid"

	request := map[string]interface{}{
		"user_uuid": userUUIDRequested,
	}

	c.repo.ExpectFind(
		where.Eq("user_uuid", userUUIDRequested),
	).Result(
		entities.User{
			ID:       1,
			UserUUID: userUUIDRequested,
		},
	)

	c.filterParams.Values = request

	got, err := FindUserByUUID(&c.filterParams)

	c.Nil(err, "FindUserByID return an error")

	userGot, ok := got.(entities.User)
	c.True(ok, "unexpected answare type")
	c.Equal(userGot.UserUUID, userUUIDRequested, "unexpected user uuid")

}

func (c *UserFilterTest) TestFilterUserByUUID_ValidatioError() {

	request := map[string]interface{}{}

	c.filterParams.Values = request

	_, err := FindUserByUUID(&c.filterParams)

	validationError, isValidationError := err.(apperrors.ValidationError)

	c.True(isValidationError, "bad type of error ")

	c.Equal(validationError.Validation, "required")
	c.Equal(validationError.Field, "user_uuid")

}

func (c *UserFilterTest) TestFilterUserByUUID_NotFound() {

	userUUIDRequested := "an-uud"

	request := map[string]interface{}{
		"user_uuid": userUUIDRequested,
	}

	c.repo.ExpectFind(
		where.Eq("user_uuid", userUUIDRequested),
	).Return(c.notFoundError)

	c.filterParams.Values = request

	_, err := FindUserByUUID(&c.filterParams)

	_, isHandableNotFoundError := err.(apperrors.UserNotFound)

	c.True(isHandableNotFoundError, "unexpected error type")

}

func (c *UserFilterTest) TestFilterUserByUUID_UnhandledError() {

	userUUIDRequested := "an-uud"

	request := map[string]interface{}{
		"user_uuid": userUUIDRequested,
	}

	c.repo.ExpectFind(
		where.Eq("user_uuid", userUUIDRequested),
	).Return(c.ordinaryError)

	c.filterParams.Values = request

	_, err := FindUserByUUID(&c.filterParams)

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

	userPageRequested := float64(2)

	request := map[string]interface{}{
		"page_number": userPageRequested,
	}

	c.paginator.On(
		"CreatePaginator",
		c.ctx,
		entities.UserTable,
		mock.Anything,
		int(userPageRequested),
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

func TestUserFilter(t *testing.T) {
	suite.Run(t, new(UserFilterTest))
}
