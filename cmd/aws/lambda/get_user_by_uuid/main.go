package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/health_records/internal/config"
	"github.com/manicar2093/health_records/internal/db/connections"
	"github.com/manicar2093/health_records/internal/db/paginator"
	"github.com/manicar2093/health_records/internal/db/repositories"
	"github.com/manicar2093/health_records/internal/userfilters"
	"github.com/manicar2093/health_records/pkg/validators"
)

func main() {
	config.StartConfig()
	conn := connections.PostgressConnection()
	paginator := paginator.NewPaginable(conn)
	userRepo := repositories.NewUserRepositoryRel(
		conn,
		paginator,
	)
	service := userfilters.NewUserByUUIDFinderImpl(
		userRepo,
		validators.NewStructValidator(),
	)
	lambda.Start(
		NewGetUserByUUIDAWSLambda(service).Handler,
	)
}
