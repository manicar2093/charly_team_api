package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/connections"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/validators"
)

func main() {
	lambda.Start(
		CreateLambdaHandlerWDependencies(
			connections.PostgressConnection(),
			validators.NewStructValidator(),
		),
	)

}

func CreateLambdaHandlerWDependencies(
	db connections.Findable,
	validator validators.ValidatorService,
) interface{} {

	return func(catalogs models.GetCatalogsRequest) *models.Response {

		isValid, response := validators.CheckValidationErrors(validator.Validate(catalogs))
		if !isValid {
			return response
		}

		gotCatalogs := make(map[string]interface{})

		for _, catalog := range catalogs.CatalogNames {
			foundCatalog, err := CatalogFactory(catalog, db)
			if err != nil {
				return validators.CreateResponseError(err)
			}
			gotCatalogs[catalog] = foundCatalog

		}

		return models.CreateResponse(http.StatusOK, gotCatalogs)

	}

}
