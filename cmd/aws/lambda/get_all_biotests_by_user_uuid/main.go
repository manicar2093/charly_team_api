package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/db/paginator"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/handlers/biotestfilters/biotestsbyuseruuidfinder"
	"github.com/manicar2093/charly_team_api/internal/config"
	"github.com/manicar2093/charly_team_api/services"
	"github.com/manicar2093/charly_team_api/validators"
)

func main() {
	config.StartConfig()
	conn := connections.PostgressConnection()
	paginator := paginator.NewPaginable(conn)
	biotestRepo := repositories.NewBiotestRepositoryRel(
		conn,
		paginator,
		services.UUIDGeneratorImpl{},
	)
	service := biotestsbyuseruuidfinder.NewBiotestByUserUUIDImpl(
		biotestRepo,
		validators.NewStructValidator(),
	)
	lambda.Start(
		NewGetAllBiotestByUserUUIDAWSLambda(service).Handler,
	)
}
