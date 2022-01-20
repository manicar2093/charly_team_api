package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/handlers/biotestupdater"
	"github.com/manicar2093/charly_team_api/validators"
)

func main() {
	config.StartConfig()
	service := biotestupdater.NewBiotestUpdater(
		connections.PostgressConnection(),
		validators.NewStructValidator(),
	)
	lambda.Start(
		NewUpdateBiotestAWSLambda(service).Handler,
	)
}
