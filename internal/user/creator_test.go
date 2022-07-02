package user_test

import (
	"context"
	"fmt"
	"time"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gopkg.in/guregu/null.v4"

	"github.com/manicar2093/charly_team_api/internal/config"
	"github.com/manicar2093/charly_team_api/internal/db/entities"
	"github.com/manicar2093/charly_team_api/internal/user"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/manicar2093/charly_team_api/pkg/validators"
)

var _ = Describe("Creator", func() {

	var (
		validator      *mocks.ValidatorService
		passGenMock    *mocks.PassGen
		uuidGen        *mocks.UUIDGenerator
		userRepo       *mocks.UserRepository
		userCreator    *user.UserCreatorImpl
		passHasherMock *mocks.HashPassword
		ctx            context.Context
		userEmail      string
	)

	BeforeEach(func() {
		validator = &mocks.ValidatorService{}
		passGenMock = &mocks.PassGen{}
		uuidGen = &mocks.UUIDGenerator{}
		userRepo = &mocks.UserRepository{}
		passHasherMock = &mocks.HashPassword{}
		userCreator = user.NewUserCreatorImpl(passGenMock, userRepo, uuidGen, validator, passHasherMock)
		ctx = context.Background()
		userEmail = "testing@main-func.com"
	})

	AfterEach(func() {
		T := GinkgoT()
		validator.AssertExpectations(T)
		passGenMock.AssertExpectations(T)
		uuidGen.AssertExpectations(T)
		userRepo.AssertExpectations(T)
	})

	It("Creates a user with given data", func() {
		var (
			userRequest = user.UserCreatorRequest{
				Name:          "testing",
				LastName:      "main",
				Email:         userEmail,
				GenderID:      1,
				Birthday:      time.Now(),
				RoleID:        1,
				BiotypeID:     2,
				BoneDensityID: 3,
			}
			userUUID         = faker.UUIDDigit()
			userAvatar       = faker.UUIDDigit()
			passReturn       = faker.Password()
			passDigestReturn = faker.Jwt()
		)
		validator.On("Validate", &userRequest).Return(validators.ValidateOutput{IsValid: true, Err: nil})
		passGenMock.On("Generate").Return(passReturn, nil)
		passHasherMock.On("Digest", passReturn).Return(passDigestReturn, nil)
		uuidGen.On("New").Return(userUUID).Once()
		uuidGen.On("New").Return(userAvatar).Once()
		userRepo.On("SaveUser", ctx, &entities.User{
			Name:          userRequest.Name,
			LastName:      userRequest.LastName,
			RoleID:        int32(userRequest.RoleID),
			GenderID:      null.IntFrom(int64(userRequest.GenderID)),
			Email:         userRequest.Email,
			Password:      passDigestReturn,
			Birthday:      userRequest.Birthday,
			AvatarUrl:     fmt.Sprintf("%s%s.svg", config.AvatarURLSrc, userAvatar),
			UserUUID:      userUUID,
			BiotypeID:     null.NewInt(int64(userRequest.BiotypeID), true),
			BoneDensityID: null.NewInt(int64(userRequest.BoneDensityID), true),
		}).Return(nil)

		res, err := userCreator.Run(ctx, &userRequest)

		Expect(err).ToNot(HaveOccurred())
		Expect(res.UserCreated.UserUUID).To(Equal(userUUID))
		Expect(res.UserCreated.AvatarUrl).To(ContainSubstring(userAvatar))

	})

})
