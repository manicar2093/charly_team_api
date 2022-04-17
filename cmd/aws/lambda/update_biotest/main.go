package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/internal/biotest"
	"github.com/manicar2093/charly_team_api/internal/config"
	"github.com/manicar2093/charly_team_api/internal/db/connections"
	"github.com/manicar2093/charly_team_api/pkg/validators"
)

func main() {
	config.StartConfig()
	service := biotest.NewBiotestUpdater(
		connections.PostgressConnection(),
		validators.NewStructValidator(),
	)
	lambda.Start(
		NewUpdateBiotestAWSLambda(service).Handler,
	)
}
