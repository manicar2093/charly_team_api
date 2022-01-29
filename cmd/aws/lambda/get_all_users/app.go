package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/charly_team_api/handlers/userfilters/allusersfinder"
	"github.com/manicar2093/charly_team_api/internal/models"
)

type GetAllUsersAWSLambda struct {
	allUsersFinder allusersfinder.AllUsersFinder
}

func NewGetAllUsersAWSLambda(allUsersFinder allusersfinder.AllUsersFinder) *GetAllUsersAWSLambda {
	return &GetAllUsersAWSLambda{allUsersFinder: allUsersFinder}
}

func (c *GetAllUsersAWSLambda) Handler(
	ctx context.Context,
	req allusersfinder.AllUsersFinderRequest,
) (*models.Response, error) {
	res, err := c.allUsersFinder.Run(ctx, &req)

	if err != nil {
		return models.CreateResponseFromError(err), nil
	}

	return models.CreateResponse(
		http.StatusOK,
		res.UsersFound,
	), nil
}
