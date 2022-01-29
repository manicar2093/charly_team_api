package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/handlers/cataloggetter"
	"github.com/manicar2093/charly_team_api/internal/config"
	"github.com/manicar2093/charly_team_api/validators"
)

func main() {
	config.StartConfig()
	conn := connections.PostgressConnection()
	service := cataloggetter.NewCatalogGetterImpl(
		repositories.NewCatalogRepositoryImpl(conn),
		validators.NewStructValidator(),
	)
	lambda.Start(
		NewGetCatalogsAWSLambda(service).Handler,
	)
}
