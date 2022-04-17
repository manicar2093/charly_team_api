package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/internal/config"
	"github.com/manicar2093/charly_team_api/internal/db/connections"
	"github.com/manicar2093/charly_team_api/internal/db/paginator"
	"github.com/manicar2093/charly_team_api/internal/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/userfilters"
	"github.com/manicar2093/charly_team_api/pkg/validators"
)

func main() {
	config.StartConfig()
	conn := connections.PostgressConnection()
	paginator := paginator.NewPaginable(conn)
	userRepo := repositories.NewUserRepositoryRel(
		conn,
		paginator,
	)
	service := userfilters.NewUserByUUIDFinderImpl(
		userRepo,
		validators.NewStructValidator(),
	)
	lambda.Start(
		NewGetUserByUUIDAWSLambda(service).Handler,
	)
}
