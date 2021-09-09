package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/validators"
)

func main() {
	config.StartConfig()
	lambda.Start(
		CreateLambdaHandlerWDependencies(
			repositories.NewCatalogRepositoryImpl(
				connections.PostgressConnection(),
			),
			validators.NewStructValidator(),
		),
	)

}

func CreateLambdaHandlerWDependencies(
	catalogsRepository repositories.CatalogRepository,
	validator validators.ValidatorService,
) func(ctx context.Context, catalogs models.GetCatalogsRequest) (*models.Response, error) {

	return func(ctx context.Context, catalogs models.GetCatalogsRequest) (*models.Response, error) {

		isValid, response := validators.CheckValidationErrors(validator.Validate(catalogs))
		if !isValid {
			return response, nil
		}

		gotCatalogs, err := CatalogFactoryLoop(catalogs, catalogsRepository, ctx)

		if err != nil {
			return models.CreateResponseFromError(err), nil
		}

		return models.CreateResponse(http.StatusOK, gotCatalogs), nil

	}

}
