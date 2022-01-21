package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/handlers/biotestcreator"
	"github.com/manicar2093/charly_team_api/services"
	"github.com/manicar2093/charly_team_api/validators"
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
