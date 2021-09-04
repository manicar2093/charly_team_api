package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/db/paginator"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/validators"
)

func main() {
	repo := connections.PostgressConnection()
	lambda.Start(
		CreateLambdaHandlerWDependencies(
			repo,
			validators.NewStructValidator(),
			paginator.NewPaginable(repo),
			NewUserFilterService(),
		),
	)
}

func CreateLambdaHandlerWDependencies(
	repo rel.Repository,
	validator validators.ValidatorService,
	paginator paginator.Paginable,
	userFilters UserFilterService,
) func(ctx context.Context, req UserFilter) *models.Response {

	return func(ctx context.Context, req UserFilter) *models.Response {

		isValid, response := validators.CheckValidationErrors(validator.Validate(req))
		if !isValid {
			return response
		}

		filter, isFilterExists := userFilters.GetUserFilter(req.FilterName)
		if !isFilterExists {
			return models.CreateResponse(http.StatusBadRequest, models.ErrorReponse{Error: "requested filter does not exists"})
		}

		items, err := filter(ctx, repo, req.Values, paginator)
		if err != nil {
			return models.CreateResponseFromError(err)
		}

		return models.CreateResponse(http.StatusOK, items)
	}
}
