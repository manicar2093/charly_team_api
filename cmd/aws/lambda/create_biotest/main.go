package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/internal/config"
	"github.com/manicar2093/charly_team_api/internal/handlers/biotestcreator"
	"github.com/manicar2093/charly_team_api/internal/services"
	"github.com/manicar2093/charly_team_api/internal/validators"
)

func main() {
	config.StartConfig()
	service := biotestcreator.NewBiotestCreator(
		connections.PostgressConnection(),
		validators.NewStructValidator(),
		services.UUIDGeneratorImpl{},
	)
	lambda.Start(
		NewCreateBiotestAWSLambda(service).Handler,
	)
}
