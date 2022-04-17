package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/internal/config"
	"github.com/manicar2093/charly_team_api/internal/db/connections"
	"github.com/manicar2093/charly_team_api/internal/db/paginator"
	"github.com/manicar2093/charly_team_api/internal/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/services"
	"github.com/manicar2093/charly_team_api/internal/user"
	"github.com/manicar2093/charly_team_api/pkg/aws"
	"github.com/manicar2093/charly_team_api/pkg/validators"
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
