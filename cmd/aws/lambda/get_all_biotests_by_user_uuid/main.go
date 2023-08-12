package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/health_records/internal/biotestfilters"
	"github.com/manicar2093/health_records/internal/config"
	"github.com/manicar2093/health_records/internal/db/connections"
	"github.com/manicar2093/health_records/internal/db/paginator"
	"github.com/manicar2093/health_records/internal/db/repositories"
	"github.com/manicar2093/health_records/internal/services"
	"github.com/manicar2093/health_records/pkg/validators"
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
	service := biotestfilters.NewBiotestByUserUUIDImpl(
		biotestRepo,
		validators.NewStructValidator(),
	)
	lambda.Start(
		NewGetAllBiotestByUserUUIDAWSLambda(service).Handler,
	)
}
