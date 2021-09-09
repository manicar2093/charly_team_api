package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/db/filters"
	"github.com/manicar2093/charly_team_api/db/paginator"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/validators"
)

func main() {
	config.StartConfig()
	repo := connections.PostgressConnection()
	paginator := paginator.NewPaginable(repo)
	lambda.Start(
		CreateLambdaHandlerWDependencies(
			repo,
			validators.NewStructValidator(),
			paginator,
			NewUserFilterService(repo, paginator),
		),
	)
}

func CreateLambdaHandlerWDependencies(
	repo rel.Repository,
	validator validators.ValidatorService,
	paginator paginator.Paginable,
	userFilters filters.FilterService,
) func(ctx context.Context, req UserFilter) (*models.Response, error) {

	return func(ctx context.Context, req UserFilter) (*models.Response, error) {

		isValid, response := validators.CheckValidationErrors(validator.Validate(req))
		if !isValid {
			return response, nil
		}

		filterRunner := userFilters.GetUserFilter(req.FilterName)
		if !filterRunner.IsFound() {
			return models.CreateResponse(http.StatusBadRequest, models.ErrorReponse{Error: "requested filter does not exists"}), nil
		}

		filterParams := filters.FilterParameters{
			Ctx:       ctx,
			Repo:      repo,
			Values:    req,
			Paginator: paginator,
		}

		items, err := filterRunner.Run(&filterParams)
		if err != nil {
			return models.CreateResponseFromError(err), nil
		}

		return models.CreateResponse(http.StatusOK, items), nil
	}
}