package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/charly_team_api/internal/handlers/userfilters"
	"github.com/manicar2093/charly_team_api/pkg/models"
)

type GetAllUsersAWSLambda struct {
	allUsersFinder userfilters.AllUsersFinder
}

func NewGetAllUsersAWSLambda(allUsersFinder userfilters.AllUsersFinder) *GetAllUsersAWSLambda {
	return &GetAllUsersAWSLambda{allUsersFinder: allUsersFinder}
}

func (c *GetAllUsersAWSLambda) Handler(
	ctx context.Context,
	req userfilters.AllUsersFinderRequest,
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
