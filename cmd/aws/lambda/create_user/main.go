package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/health_records/internal/config"
	"github.com/manicar2093/health_records/internal/db/connections"
	"github.com/manicar2093/health_records/internal/db/paginator"
	"github.com/manicar2093/health_records/internal/db/repositories"
	"github.com/manicar2093/health_records/internal/services"
	"github.com/manicar2093/health_records/internal/user"
	"github.com/manicar2093/health_records/pkg/aws"
	"github.com/manicar2093/health_records/pkg/validators"
)

func main() {
	config.StartConfig()
	conn := connections.PostgressConnection()
	paginator := paginator.NewPaginable(conn)
	service := user.NewUserCreatorImpl(
		aws.NewCognitoClient(),
		services.PasswordGenerator{},
		repositories.NewUserRepositoryRel(conn, paginator),
		services.UUIDGeneratorImpl{},
		validators.NewStructValidator(),
	)
	lambda.Start(
		NewUserCreatorAWSLambda(service).Handler,
	)
}
