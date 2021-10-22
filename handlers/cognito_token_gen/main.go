package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/db/entities"
)

func main() {
	config.StartConfig()
	lambda.Start(
		CreateLambdaHandlerWDependencies(
			connections.PostgressConnection(),
		),
	)
}

func CreateLambdaHandlerWDependencies(
	repo rel.Repository,
) func(ctx context.Context, event events.CognitoEventUserPoolsPreTokenGen) (events.CognitoEventUserPoolsPreTokenGen, error) {

	return func(ctx context.Context, event events.CognitoEventUserPoolsPreTokenGen) (events.CognitoEventUserPoolsPreTokenGen, error) {
		var userToSign entities.User

		err := repo.Find(ctx, &userToSign, where.Eq("user_uuid", event.UserName))
		if err != nil {
			log.Println(err)
			return event, errors.New("user was not found")
		}

		myClaims := map[string]string{
			"name_to_show": CreateNameToShow(userToSign.Name, userToSign.LastName),
			"avatar_url":   userToSign.AvatarUrl,
			"uuid":         userToSign.UserUUID,
		}

		event.Response.ClaimsOverrideDetails.ClaimsToAddOrOverride = myClaims

		return event, nil
	}

}

// CreateNameToShow will split names to create a full name compose by first name and paternal surename
func CreateNameToShow(name, lastName string) string {
	nameSplitted := strings.Split(name, " ")
	first, _ := nameSplitted[0], ""

	sureNameSplitted := strings.Split(lastName, " ")

	paternal, _ := sureNameSplitted[0], ""

	return fmt.Sprintf("%s %s", first, paternal)

}
