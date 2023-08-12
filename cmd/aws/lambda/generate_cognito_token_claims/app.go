package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/manicar2093/health_records/internal/token"
)

type GenerateCognitoTokenClaimsAWSLambda struct {
	tokenClaimsGenerator token.TokenClaimsGenerator
}

func NewGenerateCognitoTokenClaimsAWSLambda(tokenClaimsGenerator token.TokenClaimsGenerator) *GenerateCognitoTokenClaimsAWSLambda {
	return &GenerateCognitoTokenClaimsAWSLambda{tokenClaimsGenerator: tokenClaimsGenerator}
}

func (c *GenerateCognitoTokenClaimsAWSLambda) Handler(ctx context.Context, event events.CognitoEventUserPoolsPreTokenGen) (events.CognitoEventUserPoolsPreTokenGen, error) {

	response, err := c.tokenClaimsGenerator.Run(
		ctx,
		&token.TokenClaimsGeneratorRequest{
			UserUUID: event.UserName,
		},
	)

	if err != nil {
		return event, nil
	}

	event.Response.ClaimsOverrideDetails.ClaimsToAddOrOverride = response.Claims

	return event, nil
}
