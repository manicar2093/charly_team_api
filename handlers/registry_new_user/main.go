package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/aws"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/services"
	"github.com/manicar2093/charly_team_api/validators"
)

func main() {
	config.StartConfig()
	lambda.Start(CreateLambdaHandlerWDependencies(
		services.NewUserServiceCognito(
			aws.NewCognitoClient(),
			services.PasswordGenerator{},
			connections.PostgressConnection(),
			services.UUIDGeneratorImpl{},
		),
		validators.NewStructValidator(),
	))
}

func CreateLambdaHandlerWDependencies(
	userService services.UserService,
	validator validators.ValidatorService,
) func(ctx context.Context, req models.CreateUserRequest) (*models.Response, error) {

	return func(ctx context.Context, req models.CreateUserRequest) (*models.Response, error) {

		isValid, response := isValidRequest(&req, validator)

		if !isValid {
			return response, nil
		}

		userCreated, err := userService.CreateUser(ctx, &req)

		if err != nil {
			return models.CreateResponseFromError(err), nil
		}

		return models.CreateResponse(
			http.StatusCreated,
			models.CreateUserCreationResponseFromUser(userCreated),
		), nil

	}

}

func isValidRequest(req *models.CreateUserRequest, validator validators.ValidatorService) (bool, *models.Response) {
	isValid, response := validators.CheckValidationErrors(validator.Validate(req))

	if !isValid {
		return isValid, response
	}
	if req.RoleID == 3 {

		isValid, response = validators.CheckValidationErrors(validator.Validate(req.GetCustomerValidations()))

		if !isValid {
			return isValid, response
		}
	}

	return true, nil
}
