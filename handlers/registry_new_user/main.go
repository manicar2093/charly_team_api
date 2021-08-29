package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/aws"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/services"
	"github.com/manicar2093/charly_team_api/validators"
)

func main() {
	lambda.Start(CreateLambdaHandlerWDependencies(
		repositories.NewUserRepositoryGorm(
			connections.PostgressConnection(),
		),
		services.PasswordGenerator{},
		validators.NewStructValidator(),
	))
}

func CreateLambdaHandlerWDependencies(
	repository repositories.UserRepository,
	passGen services.PassGen,
	validator validators.ValidatorService,
) interface{} {

	return func(req models.CreateUserRequest) (*models.Response, error) {

		isValid, response := validators.CheckValidationErrors(validator.Validate(req))

		if !isValid {
			return response, nil
		}

		userService := services.NewUserServiceCognito(
			aws.NewCognitoClient(),
			repository,
			passGen,
		)

		userCreated, err := userService.CreateUser(req)

		return models.CreateResponse(
			http.StatusCreated,
			models.CreateUserResponse{
				UserID: userCreated,
			}), err

	}

}
