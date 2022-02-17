package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/internal/config"
	"github.com/manicar2093/charly_team_api/internal/db/connections"
	"github.com/manicar2093/charly_team_api/internal/db/paginator"
	"github.com/manicar2093/charly_team_api/internal/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/handlers/biotestfilters/biotestsbyuseruuidfinder"
	"github.com/manicar2093/charly_team_api/internal/services"
	"github.com/manicar2093/charly_team_api/pkg/validators"
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
