package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/connections"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/services"
	"gorm.io/gorm"
)

var validator services.ValidatorService
var db *gorm.DB

func main() {

	validator = services.NewStructValidator()
	db = connections.PostgressConnection()

	lambda.Start(LambdaHandler)
}

func LambdaHandler(req models.CreateUserRequest) (models.Response, error) {
	isValid, response := apperrors.CheckValidationErrors(validator.Validate(req))

	if !isValid {
		return response, nil
	}

	err := createClient(req, db)
	if err != nil {
		return models.Response{Code: http.StatusInternalServerError, Status: http.StatusText(http.StatusInternalServerError)}, err
	}
	return models.Response{}, nil

}
