package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/manicar2093/charly_team_api/internal/handlers/tokenclaimsgenerator"
)

type GenerateCognitoTokenClaimsAWSLambda struct {
	tokenClaimsGenerator tokenclaimsgenerator.TokenClaimsGenerator
}

func NewGenerateCognitoTokenClaimsAWSLambda(tokenClaimsGenerator tokenclaimsgenerator.TokenClaimsGenerator) *GenerateCognitoTokenClaimsAWSLambda {
	return &GenerateCognitoTokenClaimsAWSLambda{tokenClaimsGenerator: tokenClaimsGenerator}
}

func (c *GenerateCognitoTokenClaimsAWSLambda) Handler(ctx context.Context, event events.CognitoEventUserPoolsPreTokenGen) (events.CognitoEventUserPoolsPreTokenGen, error) {

	response, err := c.tokenClaimsGenerator.Run(
		ctx,
		&tokenclaimsgenerator.TokenClaimsGeneratorRequest{
			UserUUID: event.UserName,
		},
	)

	if err != nil {
		return event, nil
	}

	event.Response.ClaimsOverrideDetails.ClaimsToAddOrOverride = response.Claims

	return event, nil
}
