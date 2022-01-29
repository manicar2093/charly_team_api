package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/db/paginator"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/handlers/userfilters/userbyuuidfinder"
	"github.com/manicar2093/charly_team_api/internal/config"
	"github.com/manicar2093/charly_team_api/validators"
)

func main() {
	config.StartConfig()
	conn := connections.PostgressConnection()
	paginator := paginator.NewPaginable(conn)
	userRepo := repositories.NewUserRepositoryRel(
		conn,
		paginator,
	)
	service := userbyuuidfinder.NewUserByUUIDFinderImpl(
		userRepo,
		validators.NewStructValidator(),
	)
	lambda.Start(
		NewGetUserByUUIDAWSLambda(service).Handler,
	)
}
