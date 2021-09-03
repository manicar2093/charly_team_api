package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/validators"
)

func main() {
	lambda.Start(
		CreateLambdaHandlerWDependencies(
			connections.PostgressConnection(),
			validators.NewStructValidator(),
		),
	)
}

func CreateLambdaHandlerWDependencies(
	repo rel.Repository,
	validator validators.ValidatorService,
) func(ctx context.Context, req UserFilter) *models.Response {

	return func(ctx context.Context, req UserFilter) *models.Response {
		return &models.Response{}
	}
}
