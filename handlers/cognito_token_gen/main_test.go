package main

import (
	"context"
	"strconv"
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
	userID := int32(1)
	userIDAsStr := strconv.Itoa(int(userID))
	roleDescription := "ADescription"

	userFound := entities.User{
		ID:        userID,
		Name:      name,
		LastName:  lastName,
		Email:     email,
		AvatarUrl: avatarURL,
		UserUUID:  userUUID,
	}
	repo := reltest.New()
	event := events.CognitoEventUserPoolsPreTokenGen{
		CognitoEventUserPoolsHeader: events.CognitoEventUserPoolsHeader{
			UserName: userUUID,
		},
	}
	repo.ExpectFind(where.Eq("user_uuid", userUUID)).Result(userFound)
	repo.ExpectPreload("role").Result(entities.Role{Description: roleDescription})
	eventGot, err := CreateLambdaHandlerWDependencies(repo)(context.Background(), event)
	assert.Nil(t, err, "Should not response with an error")
	gotClaims := eventGot.Response.ClaimsOverrideDetails.ClaimsToAddOrOverride

	assert.Contains(t, gotClaims["name_to_show"], "Test", "Name is not in got name to show")
	assert.Contains(t, gotClaims["name_to_show"], "Great", "Last Name is not in name to show")
	assert.Equal(t, avatarURL, gotClaims["avatar_url"])
	assert.Equal(t, userUUID, gotClaims["uuid"])
	assert.Equal(t, userIDAsStr, gotClaims["id"])
	assert.Equal(t, roleDescription, gotClaims["role"])
}
