package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/db/paginator"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/aws"
	"github.com/manicar2093/charly_team_api/internal/config"
	"github.com/manicar2093/charly_team_api/internal/handlers/usercreator"
	"github.com/manicar2093/charly_team_api/internal/services"
	"github.com/manicar2093/charly_team_api/internal/validators"
)

func main() {
	config.StartConfig()
	conn := connections.PostgressConnection()
	paginator := paginator.NewPaginable(conn)
	service := usercreator.NewUserCreatorImpl(
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
