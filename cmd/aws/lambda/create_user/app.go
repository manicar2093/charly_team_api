package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/charly_team_api/internal/user"
	"github.com/manicar2093/charly_team_api/pkg/models"
)

type CreateUserAWSLambda struct {
	userCreator user.UserCreator
}

func NewUserCreatorAWSLambda(userCreator user.UserCreator) *CreateUserAWSLambda {
	return &CreateUserAWSLambda{userCreator: userCreator}
}

func (c *CreateUserAWSLambda) Handler(ctx context.Context, req user.UserCreatorRequest) (*models.Response, error) {
	res, err := c.userCreator.Run(ctx, &req)

	if err != nil {
		return models.CreateResponseFromError(err), nil
	}

	return models.CreateResponse(
		http.StatusCreated,
		res.UserCreated,
	), nil
}
