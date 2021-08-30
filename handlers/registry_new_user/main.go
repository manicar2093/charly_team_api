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
		services.NewUserServiceCognito(
			aws.NewCognitoClient(),
			repositories.NewUserRepositoryGorm(
				connections.PostgressConnection(),
			),
			services.PasswordGenerator{},
		),
		validators.NewStructValidator(),
	))
}

func CreateLambdaHandlerWDependencies(
	userService services.UserService,
	validator validators.ValidatorService,
) func(req models.CreateUserRequest) *models.Response {

	return func(req models.CreateUserRequest) *models.Response {

		isValid, response := validators.CheckValidationErrors(validator.Validate(req))

		if !isValid {
			return response
		}

		userCreated, err := userService.CreateUser(&req)

		if err != nil {
			return models.CreateResponseFromError(err)
		}

		return models.CreateResponse(
			http.StatusCreated,
			models.CreateUserResponse{
				UserID: userCreated,
			})

	}

}
