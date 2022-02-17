package tokenclaimsgenerator

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/pkg/logger"
)

type TokenClaimsGenerator interface {
	Run(ctx context.Context, req *TokenClaimsGeneratorRequest) (*TokenClaimsGeneratorResponse, error)
}

type tokenClaimsGeneratorImpl struct {
	userRepo repositories.UserRepository
}

func NewTokenClaimsGeneratorImpl(userRepo repositories.UserRepository) *tokenClaimsGeneratorImpl {
	return &tokenClaimsGeneratorImpl{userRepo: userRepo}
}

func (c *tokenClaimsGeneratorImpl) Run(
	ctx context.Context,
	req *TokenClaimsGeneratorRequest,
) (*TokenClaimsGeneratorResponse, error) {
	logger.Info(req)
	userToSign, err := c.userRepo.FindUserByUUID(ctx, req.UserUUID)

	if err != nil {
		logger.Error(err)
		return nil, errors.New("user was not found")
	}

	myClaims := map[string]string{
		"name_to_show": createNameToShow(userToSign.Name, userToSign.LastName),
		"avatar_url":   userToSign.AvatarUrl,
		"uuid":         userToSign.UserUUID,
		"id":           strconv.Itoa(int(userToSign.ID)),
		"role":         userToSign.Role.Description,
	}

	return &TokenClaimsGeneratorResponse{Claims: myClaims}, nil
}

// CreateNameToShow will split names to create a full name compose by first name and paternal surename
func createNameToShow(name, lastName string) string {
	nameSplitted := strings.Split(name, " ")
	first, _ := nameSplitted[0], ""

	sureNameSplitted := strings.Split(lastName, " ")

	paternal, _ := sureNameSplitted[0], ""

	return fmt.Sprintf("%s %s", first, paternal)

}
