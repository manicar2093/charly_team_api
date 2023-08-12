package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/health_records/internal/config"
	"github.com/manicar2093/health_records/internal/db/connections"
	"github.com/manicar2093/health_records/internal/db/paginator"
	"github.com/manicar2093/health_records/internal/db/repositories"
	"github.com/manicar2093/health_records/internal/user"
	"github.com/manicar2093/health_records/pkg/validators"
)

func main() {
	config.StartConfig()
	repo := connections.PostgressConnection()
	paginator := paginator.NewPaginable(repo)
	userRepo := repositories.NewUserRepositoryRel(repo, paginator)
	service := user.NewUserAvatarUpdater(userRepo, validators.NewStructValidator())

	lambda.Start(NewUpdateUserAWSLambda(service).Handler)
}
