package main

import (
	"context"

	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/services"
	"github.com/manicar2093/charly_team_api/validators"
)

func main() {

}

func CreateLambdaHandlerWDependencies(
	userService services.UserService,
	validator validators.ValidatorService,
) func(ctx context.Context, req models.CreateBiotestRequest) *models.Response {

	return func(ctx context.Context, req models.CreateBiotestRequest) *models.Response {

		return &models.Response{}
	}
}
