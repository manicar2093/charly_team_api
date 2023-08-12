package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/health_records/internal/config"
	"github.com/manicar2093/health_records/internal/db/connections"
	"github.com/manicar2093/health_records/internal/db/paginator"
	"github.com/manicar2093/health_records/internal/db/repositories"
	"github.com/manicar2093/health_records/internal/token"
)

func main() {
	config.StartConfig()
	conn := connections.PostgressConnection()
	paginator := paginator.NewPaginable(conn)
	userRepo := repositories.NewUserRepositoryRel(conn, paginator)
	service := token.NewTokenClaimsGeneratorImpl(userRepo)
	lambda.Start(
		NewGenerateCognitoTokenClaimsAWSLambda(service).Handler,
	)
}
