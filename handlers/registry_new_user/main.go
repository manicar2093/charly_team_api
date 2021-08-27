package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/aws"
	"github.com/manicar2093/charly_team_api/connections"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/services"
	"github.com/manicar2093/charly_team_api/validators"
)

var validator = validators.NewStructValidator()
var db = connections.PostgressConnection()
var passGen = services.PasswordGenerator{}
var userService services.UserService

func main() {

	userService = services.NewUserServiceCognito(
		aws.NewCognitoClient(),
		db,
		passGen,
	)

	lambda.Start(LambdaHandler)
}

func LambdaHandler(req models.CreateUserRequest) (models.Response, error) {
	isValid, response := validators.CheckValidationErrors(validator.Validate(req))

	if !isValid {
		return response, nil
	}

	err := userService.CreateUser(req)

	if err != nil {
		return models.Response{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Body:   err.Error(),
		}, err
	}
	return models.Response{}, nil

}
