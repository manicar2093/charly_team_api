package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/health_records/internal/config"
	"github.com/manicar2093/health_records/internal/db/connections"
	"github.com/manicar2093/health_records/internal/user"
	"github.com/manicar2093/health_records/pkg/validators"
)

func main() {
	config.StartConfig()

	conn := connections.PostgressConnection()
	service := user.NewUpdateUser(conn, validators.NewStructValidator())
	lambda.Start(
		NewUpdateUserAWSLambda(service).Handler,
	)
}
