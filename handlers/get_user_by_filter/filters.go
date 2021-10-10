package main

import (
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/filters"
)

func FindUserByUUID(
	params *filters.FilterParameters,
) (interface{}, error) {

	valuesAsMap := params.Values.(map[string]interface{})

	userUUID, ok := valuesAsMap["user_uuid"].(string)
	if !ok {
		return nil, apperrors.ValidationError{Field: "user_uuid", Validation: "required"}
	}

	var userFound entities.User

	err := params.Repo.Find(params.Ctx, &userFound, where.Eq("user_uuid", userUUID))
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

	pageNumber, ok := valuesAsMap["page_number"].(float64)
	if !ok {
		return nil, apperrors.ValidationError{Field: "page_number", Validation: "required"}
	}

	var usersFound []entities.User

	return params.Paginator.CreatePaginator(
		params.Ctx,
		entities.User{}.Table(),
		&usersFound,
		int(pageNumber),
	)

}
