package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/validators"
	"github.com/manicar2093/charly_team_api/validators/nullsql"
)

func main() {
	config.StartConfig()
	lambda.Start(CreateLambdaHandlerWDependencies(
		validators.NewStructValidator(),
		connections.PostgressConnection(),
	))

}

func CreateLambdaHandlerWDependencies(
	validator validators.ValidatorService,
	repo rel.Repository,
) func(ctx context.Context, req BiotestImagesRequest) (*models.Response, error) {

	return func(ctx context.Context, req BiotestImagesRequest) (*models.Response, error) {

		isValid, response := validators.CheckValidationErrors(validator.Validate(req))

		if !isValid {
			return response, nil
		}

		err := repo.Transaction(ctx, func(ctx context.Context) error {
			var biotest entities.Biotest

			err := repo.Find(ctx, &biotest, where.Eq("biotest_uuid", req.BiotestUUID))
			if err != nil {
				return err
			}

			biotest.FrontPicture = nullsql.ValidateStringSQLValid(req.FrontPicture)
			biotest.BackPicture = nullsql.ValidateStringSQLValid(req.BackPicture)
			biotest.LeftSidePicture = nullsql.ValidateStringSQLValid(req.LeftSidePicture)
			biotest.RightSidePicture = nullsql.ValidateStringSQLValid(req.RightSidePicture)

			err = repo.Update(ctx, &biotest)
			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			return models.CreateResponseFromError(err), nil
		}
		return models.CreateResponse(
			http.StatusOK,
			nil,
		), nil
	}
}
