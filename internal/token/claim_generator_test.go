package token_test

import (
	"context"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/manicar2093/charly_team_api/internal/db/entities"
	"github.com/manicar2093/charly_team_api/internal/token"
	"github.com/manicar2093/charly_team_api/mocks"
)

var _ = Describe("ClaimGenerator", func() {

	var (
		name            string
		lastName        string
		avatarURL       string
		userUUID        string
		userID          int32
		userIDAsStr     string
		roleDescription string
		userFound       entities.User
		ctx             context.Context
		userRepo        mocks.UserRepository
	)

	BeforeEach(func() {
		name = "Test Testing"
		lastName = "Great System"
		avatarURL = "a_avatar_url"
		userUUID = "an_uuid"
		userID = int32(1)
		userIDAsStr = strconv.Itoa(int(userID))
		roleDescription = "ADescription"
		userFound = entities.User{
			ID:        userID,
			Name:      name,
			LastName:  lastName,
			AvatarUrl: avatarURL,
			UserUUID:  userUUID,
			Role:      entities.Role{ID: 1, Description: roleDescription},
		}
		ctx = context.Background()
		userRepo = mocks.UserRepository{}
	})

	Describe("TokenClaimsGeneratorImpl", func() {
		It("creates the data required to token", func() {
			userRepo.On("FindUserByUUID", ctx, userUUID).Return(&userFound, nil)
			request := token.TokenClaimsGeneratorRequest{UserUUID: userUUID}
			service := token.NewTokenClaimsGeneratorImpl(&userRepo)

			got, err := service.Run(ctx, &request)

			Expect(err).ToNot(HaveOccurred())
			Expect(got.Claims["name_to_show"]).To(ContainSubstring("Test"))
			Expect(got.Claims["name_to_show"]).To(ContainSubstring("Great"))
			Expect(got.Claims["avatar_url"]).To(Equal(avatarURL))
			Expect(got.Claims["uuid"]).To(Equal(userUUID))
			Expect(got.Claims["id"]).To(Equal(userIDAsStr))
			Expect(got.Claims["role"]).To(Equal(roleDescription))
		})
	})

	Describe("TokenClaimsFromUserImpl", func() {
		It("generates needed claims", func() {
			service := token.NewTokenClaimsFromUserImpl()

			got := service.Generate(ctx, &userFound)

			Expect(got["name_to_show"]).To(ContainSubstring("Test"))
			Expect(got["name_to_show"]).To(ContainSubstring("Great"))
			Expect(got["avatar_url"]).To(Equal(avatarURL))
			Expect(got["uuid"]).To(Equal(userUUID))
			Expect(got["id"]).To(Equal(userID))
			Expect(got["role"]).To(Equal(roleDescription))
		})
	})

})
