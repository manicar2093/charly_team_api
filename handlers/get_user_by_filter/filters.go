package main

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/models"
)

// FilterFunc represents a result getter
type FilterFunc func(ctx context.Context, repo rel.Repository, values interface{}) (interface{}, error)

var userFilterRegistered map[string]FilterFunc

func FindUserByID(ctx context.Context, repo rel.Repository, values interface{}) (interface{}, error) {

	userID, ok := values.(int)
	if !ok {
		return nil, apperrors.ValidationError{Field: "user_id", Validation: "required"}
	}

	var userFound entities.User

	err := repo.Find(ctx, &userFound, where.Eq("id", userID))
	if err != nil {
		return nil, err
	}
	return userFound, nil
}

func FindUsersByEmail(ctx context.Context, repo rel.Repository, values interface{}) (interface{}, error) {

	userEmail, ok := values.(string)
	if !ok {
		return nil, apperrors.ValidationError{Field: "email", Validation: "required"}
	}

	totalUsers, err := repo.Count(ctx, entities.User{}.Table())
	if err != nil {
		return nil, err
	}

	var usersFound []entities.User

	err = repo.Find(ctx, &usersFound, where.Like("email", "%"+userEmail+"%"))
	if err != nil {
		return nil, err
	}

	paginator := models.Paginator{
		TotalPages: totalUsers / config.PageSize,
		Data:       usersFound,
	}

	return paginator, nil
}
