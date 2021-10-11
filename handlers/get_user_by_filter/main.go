package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/db/filters"
	"github.com/manicar2093/charly_team_api/db/paginator"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/validators"
)

var userFilters = []filters.FilterRegistrationData{
	{Name: "find_user_by_uuid", Func: FindUserByUUID},
	{Name: "find_all_users", Func: FindAllUsers},
	{Name: "find_user_like_email_or_name_or_last_name", Func: FindUserLikeEmailOrNameOrLastName},
}

func main() {
	config.StartConfig()
	repo := connections.PostgressConnection()
	paginator := paginator.NewPaginable(repo)
	validator := validators.NewStructValidator()
	userFilter := filters.NewFilter(&filters.FilterParameters{
		Repo:      repo,
		Paginator: paginator,
		Validator: validator,
	}, userFilters...)
	lambda.Start(
		CreateLambdaHandlerWDependencies(
			validator,
			userFilter,
		),
	)
}

func CreateLambdaHandlerWDependencies(
	validator validators.ValidatorService,
	userFilter filters.Filterable,
) func(ctx context.Context, req models.FilterRequest) (*models.Response, error) {

	return func(ctx context.Context, req models.FilterRequest) (*models.Response, error) {

		isValid, response := validators.CheckValidationErrors(validator.Validate(req))
		if !isValid {
			return response, nil
		}

		err := userFilter.GetFilter(req.FilterName)
		if err != nil {
			return models.CreateResponseFromError(err), nil
		}

		userFilter.SetContext(ctx)
		userFilter.SetValues(req.Values)
		items, err := userFilter.Run()
		if err != nil {
			return models.CreateResponseFromError(err), nil
		}

		return models.CreateResponse(http.StatusOK, items), nil
	}
}
