package repositories_test

import (
	"context"
	"strings"

	"github.com/go-rel/rel/where"
	"github.com/go-rel/reltest"
	"github.com/jaswdr/faker"
	"github.com/manicar2093/charly_team_api/internal/db/entities"
	"github.com/manicar2093/charly_team_api/internal/db/paginator"
	"github.com/manicar2093/charly_team_api/internal/db/repositories"
	"github.com/manicar2093/charly_team_api/mocks"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("UserRepository", func() {
	var (
		paginatorMock  *mocks.Paginable
		repo           *reltest.Repository
		userRepository *repositories.UserRepositoryRel
		ctx            context.Context
		fake           faker.Faker
	)

	BeforeEach(func() {
		repo = reltest.New()
		paginatorMock = &mocks.Paginable{}
		userRepository = repositories.NewUserRepositoryRel(repo, paginatorMock)
		ctx = context.TODO()
		fake = faker.New()

	})

	AfterEach(func() {
		t := GinkgoT()
		repo.AssertExpectations(t)
		paginatorMock.AssertExpectations(t)
	})

	Describe("FindUserByUUID", func() {
		It("finds a user using given uuid", func() {
			expectedUserUUID := fake.UUID().V4()
			expectedUserID := fake.Int32()
			userReturned := entities.User{
				ID:       expectedUserID,
				UserUUID: expectedUserUUID,
			}
			repo.ExpectFind(
				where.Eq("user_uuid", expectedUserUUID),
			).Result(userReturned)

			got, err := userRepository.FindUserByUUID(ctx, expectedUserUUID)

			Expect(err).ToNot(HaveOccurred())
			Expect(got.UserUUID).To(Equal(expectedUserUUID))
		})

		When("user is not found", func() {
			It("returns a repositories.NotFound error", func() {
				expectedUserUUID := fake.UUID().V4()
				repo.ExpectFind(
					where.Eq("user_uuid", expectedUserUUID),
				).NotFound()

				got, err := userRepository.FindUserByUUID(ctx, expectedUserUUID)

				Expect(err).To(BeAssignableToTypeOf(&repositories.NotFoundError{}))
				Expect(err.Error()).To(ContainSubstring(expectedUserUUID))
				Expect(got).To(BeNil())
			})
		})
	})

	Describe("FindUserLikeEmailOrNameOrLastName", func() {
		It("find users by given data", func() {
			expectedSearchParam := "expectedSearchParam"
			expectSearchParamLower := strings.ToLower(expectedSearchParam)
			usersReturned := []entities.User{
				{},
				{},
				{},
				{},
			}
			expectedFilter := where.Like("LOWER(email)", "%"+expectSearchParamLower+"%").OrLike("LOWER(name)", "%"+expectSearchParamLower+"%").OrLike("LOWER(last_name)", "%"+expectSearchParamLower+"%")
			repo.ExpectFindAll(expectedFilter).Result(usersReturned)

			got, err := userRepository.FindUserLikeEmailOrNameOrLastName(ctx, expectedSearchParam)

			Expect(err).ToNot(HaveOccurred())
			Expect(got).ToNot(BeNil())
			Expect(*got).To(HaveLen(len(usersReturned)))
		})
	})

	Describe("FindAllUsers", func() {
		It("creates a paginator to return data", func() {
			pageSort := paginator.PageSort{
				Page:         1,
				ItemsPerPage: 10,
			}
			usersHolder := []entities.User{}
			paginationReturn := paginator.Paginator{Data: []entities.User{{}, {}}}
			paginatorMock.On(
				"CreatePagination",
				ctx,
				entities.UserTable,
				&usersHolder,
				&pageSort,
			).Return(&paginationReturn, nil)

			got, err := userRepository.FindAllUsers(ctx, &pageSort)

			Expect(err).ToNot(HaveOccurred())
			Expect(got).To(BeAssignableToTypeOf(&paginator.Paginator{}))
			Expect(got.Data).To(BeAssignableToTypeOf([]entities.User{}))

		})
	})

	Describe("SaveUser", func() {
		It("saves given user", func() {
			expectedUserUUID := fake.UUID().V4()
			expectedUserToSave := entities.User{
				UserUUID: expectedUserUUID,
			}

			repo.ExpectTransaction(func(r *reltest.Repository) {
				r.ExpectInsert().For(&expectedUserToSave)
			})

			err := userRepository.SaveUser(ctx, &expectedUserToSave)

			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("UpdateUser", func() {
		It("updates a given user", func() {
			expectedUserUUID := fake.UUID().V4()
			expectedUserToUpdate := entities.User{
				UserUUID: expectedUserUUID,
			}
			repo.ExpectTransaction(func(r *reltest.Repository) {
				r.ExpectUpdate().For(&expectedUserToUpdate)
			})

			err := userRepository.UpdateUser(ctx, &expectedUserToUpdate)

			Expect(err).ToNot(HaveOccurred())
		})
	})
})
