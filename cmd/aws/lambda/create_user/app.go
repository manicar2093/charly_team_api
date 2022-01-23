package main

import (
	"context"
	"net/http"

	"github.com/manicar2093/charly_team_api/handlers/usercreator"
	"github.com/manicar2093/charly_team_api/models"
)

type CreateUserAWSLambda struct {
	userCreator usercreator.UserCreator
}

func NewUserCreatorAWSLambda(userCreator usercreator.UserCreator) *CreateUserAWSLambda {
	return &CreateUserAWSLambda{userCreator: userCreator}
}

func (c *CreateUserAWSLambda) Handler(ctx context.Context, req usercreator.UserCreatorRequest) (*models.Response, error) {
	res, err := c.userCreator.Run(ctx, &req)

	if err != nil {
		return models.CreateResponseFromError(err), nil
	}

	return models.CreateResponse(
		http.StatusCreated,
		res.UserCreated,
	), nil
}
