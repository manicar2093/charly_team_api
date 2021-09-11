package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/validators"
)

func main() {
	config.StartConfig()
	lambda.Start(CreateLambdaHandlerWDependencies(
		connections.PostgressConnection(),
		validators.NewStructValidator(),
	))
}

func CreateLambdaHandlerWDependencies(
	repo rel.Repository,
	validator validators.ValidatorService,
) func(ctx context.Context, req entities.User) (*models.Response, error) {

	return func(ctx context.Context, req entities.User) (*models.Response, error) {

		if req.ID == 0 {
			return models.CreateResponse(
				http.StatusBadRequest,
				apperrors.ValidationErrors{{Field: "id", Validation: "required"}},
			), nil
		}

		isValid, response := validators.CheckValidationErrors(validator.Validate(req))

		if !isValid {
			return response, nil
		}

		err := repo.Update(ctx, &req)
		if err != nil {
			return models.CreateResponseFromError(err), nil
		}

		return models.CreateResponse(
			http.StatusOK,
			nil,
		), nil

	}

}
