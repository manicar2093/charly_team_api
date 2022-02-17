package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/db/paginator"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/config"
	"github.com/manicar2093/charly_team_api/internal/handlers/biotestcreator"
	"github.com/manicar2093/charly_team_api/internal/services"
	"github.com/manicar2093/charly_team_api/pkg/validators"
)

func main() {
	config.StartConfig()
	conn := connections.PostgressConnection()
	paginator := paginator.NewPaginable(conn)
	biotestRepo := repositories.NewBiotestRepositoryRel(conn, paginator, services.UUIDGeneratorImpl{})
	service := biotestcreator.NewBiotestCreator(
		biotestRepo,
		validators.NewStructValidator(),
		services.UUIDGeneratorImpl{},
	)
	lambda.Start(
		NewCreateBiotestAWSLambda(service).Handler,
	)
}
