package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/services"
	"github.com/manicar2093/charly_team_api/validators"
)

func main() {
	config.StartConfig()
	lambda.Start(
		CreateLambdaHandlerWDependencies(
			connections.PostgressConnection(),
			validators.NewStructValidator(),
			services.UUIDGeneratorImpl{},
		),
	)

}

func CreateLambdaHandlerWDependencies(
	repo rel.Repository,
	validator validators.ValidatorService,
	uuidGen services.UUIDGenerator,
) func(ctx context.Context, req entities.Biotest) (*models.Response, error) {

	return func(ctx context.Context, req entities.Biotest) (*models.Response, error) {

		isValid, response := validators.CheckValidationErrors(validator.Validate(req))

		if !isValid {
			return response, nil
		}

		req.BiotestUUID = uuidGen.New()

		err := repo.Insert(ctx, &req)
		if err != nil {
			return models.CreateResponseFromError(err), nil
		}

		return models.CreateResponse(
			http.StatusCreated,
			CreateBiotestResponse{
				BiotestID: req.ID,
			},
		), nil
	}
}
