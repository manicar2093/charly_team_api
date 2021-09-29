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

const (
	Customer_Role_ID = 3
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

		isValid, response := assignDefaultRoleToRequestAndValidate(&req, validator)

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

// assignDefaultRoleToRequestAndValidate assign Customer_Role_ID to request and do req validation
func assignDefaultRoleToRequestAndValidate(
	req *models.CreateUserRequest,
	validator validators.ValidatorService,
) (bool, *models.Response) {

	req.RoleID = Customer_Role_ID
	return validators.CheckValidationErrors(validator.Validate(req))
}
