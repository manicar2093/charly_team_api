package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/handlers/userupdater"
	"github.com/manicar2093/charly_team_api/validators"
)

func main() {
	config.StartConfig()

	conn := connections.PostgressConnection()
	service := userupdater.NewUpdateUser(conn, validators.NewStructValidator())
	lambda.Start(
		NewUpdateUserAWSLambda(service).Handler,
	)
}
