package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/db/paginator"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/config"
	"github.com/manicar2093/charly_team_api/internal/handlers/userfilters/userlikeemailornamefinder"
	"github.com/manicar2093/charly_team_api/internal/validators"
)

func main() {
	config.StartConfig()

	conn := connections.PostgressConnection()
	paginator := paginator.NewPaginable(conn)
	userRepo := repositories.NewUserRepositoryRel(conn, paginator)
	services := userlikeemailornamefinder.NewUserLikeEmailOrNameFinderImpl(
		userRepo,
		validators.NewStructValidator(),
	)
	lambda.Start(
		NewGetUsersLikeEmailOrNameAWSLambda(services).Handler,
	)
}
