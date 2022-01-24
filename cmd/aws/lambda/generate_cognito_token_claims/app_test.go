package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/manicar2093/charly_team_api/handlers/tokenclaimsgenerator"
	"github.com/stretchr/testify/assert"
)

func TestCognitoTokenGen(t *testing.T) {
	ctx := context.Background()
	userUUID := "user_uuid"
	event := events.CognitoEventUserPoolsPreTokenGen{
		CognitoEventUserPoolsHeader: events.CognitoEventUserPoolsHeader{
			UserName: userUUID,
		},
	}
	tokenClaimsGenerator := tokenclaimsgenerator.MockTokenClaimsGenerator{}
	returnedClaims := map[string]string{"data": "data"}
	expectedRunCall := tokenclaimsgenerator.TokenClaimsGeneratorRequest{UserUUID: userUUID}
	tokenClaimsGenerator.On("Run", ctx, &expectedRunCall).Return(
		&tokenclaimsgenerator.TokenClaimsGeneratorResponse{Claims: returnedClaims},
		nil,
	)
	service := NewGenerateCognitoTokenClaimsAWSLambda(&tokenClaimsGenerator)

	got, err := service.Handler(context.Background(), event)

	assert.Nil(t, err, "Should not response with an error")
	assert.Equal(t, returnedClaims, got.Response.ClaimsOverrideDetails.ClaimsToAddOrOverride)

}
