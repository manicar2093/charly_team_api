package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/db/paginator"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/config"
	"github.com/manicar2093/charly_team_api/internal/handlers/user"
	"github.com/manicar2093/charly_team_api/internal/validators"
)

func main() {
	config.StartConfig()
	repo := connections.PostgressConnection()
	paginator := paginator.NewPaginable(repo)
	userRepo := repositories.NewUserRepositoryRel(repo, paginator)
	service := user.NewUserAvatarUpdater(userRepo, validators.NewStructValidator())

	lambda.Start(NewUpdateUserAWSLambda(service).Handler)
}
