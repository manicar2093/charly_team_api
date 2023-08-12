package token_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/manicar2093/health_records/internal/db/entities"
	"github.com/manicar2093/health_records/internal/token"
	"github.com/manicar2093/health_records/mocks"
	"github.com/stretchr/testify/assert"
)

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
		Role:      entities.Role{ID: 1, Description: roleDescription},
	}
	ctx := context.Background()
	userRepo := mocks.UserRepository{}
	userRepo.On("FindUserByUUID", ctx, userUUID).Return(&userFound, nil)
	request := token.TokenClaimsGeneratorRequest{UserUUID: userUUID}
	service := token.NewTokenClaimsGeneratorImpl(&userRepo)

	got, err := service.Run(ctx, &request)

	assert.Nil(t, err, "Should not response with an error")

	assert.Contains(t, got.Claims["name_to_show"], "Test", "Name is not in got name to show")
	assert.Contains(t, got.Claims["name_to_show"], "Great", "Last Name is not in name to show")
	assert.Equal(t, avatarURL, got.Claims["avatar_url"])
	assert.Equal(t, userUUID, got.Claims["uuid"])
	assert.Equal(t, userIDAsStr, got.Claims["id"])
	assert.Equal(t, roleDescription, got.Claims["role"])
}
