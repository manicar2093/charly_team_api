package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/health_records/internal/biotest"
	"github.com/manicar2093/health_records/internal/config"
	"github.com/manicar2093/health_records/internal/db/connections"
	"github.com/manicar2093/health_records/pkg/validators"
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
