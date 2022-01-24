package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/db/filters"
	"github.com/manicar2093/charly_team_api/db/paginator"
	"github.com/manicar2093/charly_team_api/handlers/get_biotest_by_filter/biotestfilters"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/validators"
)

var biotestFiltersRegistered = []filters.FilterRegistrationData{
	{Name: "get_biotest_comparision", Func: biotestfilters.GetBiotestComparision},
	{Name: "get_all_user_biotests", Func: biotestfilters.GetAllUserBiotest},
}

func main() {
	config.StartConfig()
	repo := connections.PostgressConnection()
	paginator := paginator.NewPaginable(repo)
	validator := validators.NewStructValidator()
	biotestFilter := filters.NewFilter(&filters.FilterParameters{
		Repo:      repo,
		Paginator: paginator,
		Validator: validator,
	}, biotestFiltersRegistered...)
	lambda.Start(
		CreateLambdaHandlerWDependencies(validator, biotestFilter),
	)

}

func CreateLambdaHandlerWDependencies(
	validator validators.ValidatorService,
	biotestFilter filters.Filterable,
) func(ctx context.Context, req models.FilterRequest) (*models.Response, error) {

	return func(ctx context.Context, req models.FilterRequest) (*models.Response, error) {

		isValid, response := validators.CheckValidationErrors(validator.Validate(req))
		if !isValid {
			return response, nil
		}

		err := biotestFilter.GetFilter(req.FilterName)
		if err != nil {
			return models.CreateResponseFromError(err), nil
		}

		biotestFilter.SetContext(ctx)
		biotestFilter.SetValues(req.Values)

		items, err := biotestFilter.Run()
		if err != nil {
			return models.CreateResponseFromError(err), nil
		}

		return models.CreateResponse(http.StatusOK, items), nil
	}

}
