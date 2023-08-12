package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/health_records/internal/catalog"
	"github.com/manicar2093/health_records/internal/config"
	"github.com/manicar2093/health_records/internal/db/connections"
	"github.com/manicar2093/health_records/internal/db/repositories"
	"github.com/manicar2093/health_records/pkg/validators"
)

func main() {
	config.StartConfig()
	conn := connections.PostgressConnection()
	service := catalog.NewCatalogGetterImpl(
		repositories.NewCatalogRepositoryImpl(conn),
		validators.NewStructValidator(),
	)
	lambda.Start(
		NewGetCatalogsAWSLambda(service).Handler,
	)
}
