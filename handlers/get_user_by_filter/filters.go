package main

import (
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/filters"
	"github.com/manicar2093/charly_team_api/db/paginator"
)

type UserFilterService struct {
	filters   map[string]filters.FilterFunc
	repo      rel.Repository
	paginator paginator.Paginable
}

func NewUserFilterService(
	repo rel.Repository,
	paginator paginator.Paginable,
) filters.FilterService {
	userService := UserFilterService{
		filters:   make(map[string]filters.FilterFunc),
		repo:      repo,
		paginator: paginator,
	}
	userService.filters["find_user_by_id"] = FindUserByID
	userService.filters["find_all_users"] = FindAllUsers
	userService.filters["find_user_by_email"] = FindUserByEmail
	return &userService
}

func (c UserFilterService) GetFilter(filterName string) filters.FilterRunable {

	filterFound, isFound := c.filters[filterName]
	return filters.FilterRunner{Filter: filterFound, Found: isFound}

}

func FindUserByID(
	params *filters.FilterParameters,
) (interface{}, error) {

	valuesAsMap := params.Values.(map[string]interface{})

	userID, ok := valuesAsMap["user_id"].(int)
	if !ok {
		return nil, apperrors.ValidationError{Field: "user_id", Validation: "required"}
	}

	var userFound entities.User

	err := params.Repo.Find(params.Ctx, &userFound, where.Eq("id", userID))
	if err != nil {
		if _, ok := err.(rel.NotFoundError); ok {
			return nil, apperrors.UserNotFound{}
		}
		return nil, err
	}
	return userFound, nil
}

func FindUserByEmail(
	params *filters.FilterParameters,
) (interface{}, error) {

	valuesAsMap := params.Values.(map[string]interface{})

	userEmail, ok := valuesAsMap["email"].(string)
	if !ok {
		return nil, apperrors.ValidationError{Field: "email", Validation: "required"}
	}

	var userFound entities.User

	err := params.Repo.Find(params.Ctx, &userFound, where.Like("email", "%"+userEmail+"%"))

	if err != nil {
		if _, ok := err.(rel.NotFoundError); ok {
			return nil, apperrors.UserNotFound{}
		}
		return nil, err
	}

	return userFound, nil

}

func FindAllUsers(
	params *filters.FilterParameters,
) (interface{}, error) {

	valuesAsMap := params.Values.(map[string]interface{})

	pageNumber, ok := valuesAsMap["page_number"].(int)
	if !ok {
		return nil, apperrors.ValidationError{Field: "page_number", Validation: "required"}
	}

	var usersFound []entities.User

	return params.Paginator.CreatePaginator(
		params.Ctx,
		entities.User{}.Table(),
		&usersFound,
		pageNumber,
	)

}
