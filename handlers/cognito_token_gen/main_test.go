package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/go-rel/rel/reltest"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/stretchr/testify/assert"
)

func TestCreateNameToShow(t *testing.T) {
	name := "Test Testing"
	lastName := "Great System"

	got := CreateNameToShow(name, lastName)

	assert.Contains(t, got, "Test", "Name is not in got name to show")
	assert.Contains(t, got, "Great", "Last Name is not in name to show")
}

func TestCognitoTokenGen(t *testing.T) {
	email := "an_email@email.com"
	name := "Test Testing"
	lastName := "Great System"
	avatarURL := "a_avatar_url"
	userUUID := "an_uuid"

	userFound := entities.User{
		Name:      name,
		LastName:  lastName,
		Email:     email,
		AvatarUrl: avatarURL,
		UserUUID:  userUUID,
	}
	repo := reltest.New()
	event := events.CognitoEventUserPoolsPreTokenGen{
		Request: events.CognitoEventUserPoolsPreTokenGenRequest{
			UserAttributes: map[string]string{
				"email": email,
			},
		},
	}
	repo.ExpectFind(where.Eq("email", email)).Result(userFound)
	eventGot, err := CreateLambdaHandlerWDependencies(repo)(context.Background(), event)
	assert.Nil(t, err, "Should not response with an error")
	gotClaims := eventGot.Response.ClaimsOverrideDetails.ClaimsToAddOrOverride

	assert.Contains(t, gotClaims["name_to_show"], "Test", "Name is not in got name to show")
	assert.Contains(t, gotClaims["name_to_show"], "Great", "Last Name is not in name to show")
	assert.Equal(t, avatarURL, gotClaims["avatar_url"])
	assert.Equal(t, userUUID, gotClaims["uuid"])
}
