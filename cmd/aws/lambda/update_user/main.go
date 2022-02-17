package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/internal/config"
	"github.com/manicar2093/charly_team_api/internal/handlers/userupdater"
	"github.com/manicar2093/charly_team_api/pkg/validators"
)

func main() {
	config.StartConfig()

	conn := connections.PostgressConnection()
	service := userupdater.NewUpdateUser(conn, validators.NewStructValidator())
	lambda.Start(
		NewUpdateUserAWSLambda(service).Handler,
	)
}
