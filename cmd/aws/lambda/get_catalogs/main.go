package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/internal/config"
	"github.com/manicar2093/charly_team_api/internal/db/connections"
	"github.com/manicar2093/charly_team_api/internal/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/handlers/cataloggetter"
	"github.com/manicar2093/charly_team_api/pkg/validators"
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
