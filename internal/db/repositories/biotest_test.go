package repositories_test

import (
	"context"
	"fmt"
	"time"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/sort"
	"github.com/go-rel/rel/where"
	"github.com/go-rel/reltest"
	"github.com/jaswdr/faker"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/manicar2093/charly_team_api/internal/db/entities"
	"github.com/manicar2093/charly_team_api/internal/db/paginator"
	"github.com/manicar2093/charly_team_api/internal/db/repositories"
	"github.com/manicar2093/charly_team_api/mocks"
)

var _ = Describe("Biotest", func() {
	var (
		paginatorMock     *mocks.Paginable
		uuidGen           *mocks.UUIDGenerator
		repo              *reltest.Repository
		biotestRepository repositories.BiotestRepository
		ctx               context.Context
		fake              faker.Faker
	)

	BeforeEach(func() {
		repo = reltest.New()
		paginatorMock = &mocks.Paginable{}
		uuidGen = &mocks.UUIDGenerator{}
		biotestRepository = repositories.NewBiotestRepositoryRel(repo, paginatorMock, uuidGen)
		ctx = context.TODO()
		fake = faker.New()
	})

	AfterEach(func() {
		t := GinkgoT()
		repo.AssertExpectations(t)
		paginatorMock.AssertExpectations(t)
		uuidGen.AssertExpectations(t)
	})

	Describe("FindBiotestByUUID", func() {
		It("should be returned", func() {
			expectedBiotestUUID := fake.UUID().V4()
			expectedBiotest := entities.Biotest{
				BiotestUUID: expectedBiotestUUID,
			}
			repo.ExpectFind(
				where.Eq("biotest_uuid", expectedBiotestUUID),
			).Result(expectedBiotest)

			got, err := biotestRepository.FindBiotestByUUID(ctx, expectedBiotestUUID)

			Expect(err).ToNot(HaveOccurred())
			Expect(expectedBiotestUUID).To(Equal(got.BiotestUUID))

		})
		Context("if biotest does not exist", func() {
			It("should return a repositories.NotFoundError", func() {
				expectedBiotestUUID := fake.UUID().V4()
				repo.ExpectFind(
					where.Eq("biotest_uuid", expectedBiotestUUID),
				).NotFound()

				got, err := biotestRepository.FindBiotestByUUID(ctx, expectedBiotestUUID)

				Expect(err).To(BeAssignableToTypeOf(repositories.NotFoundError{}))
				Expect(got).To(BeNil())

			})
		})

	})

	Describe("GetAllUserBiotestByUserUUID", func() {
		It("should return all user biotest", func() {
			expectedUserUUID := fake.UUID().V4()
			expectedUserID := fake.Int32()
			userFindReturn := entities.User{
				ID:       expectedUserID,
				UserUUID: expectedUserUUID,
			}
			repo.ExpectFind(
				where.Eq("user_uuid", expectedUserUUID),
			).Result(userFindReturn)
			pageNumber := fake.Float64(2, 1, 2)
			pageNumberAsInt := int(pageNumber)
			pageSort := paginator.PageSort{
				Page: pageNumber,
			}
			pageSort.SetFiltersQueries(where.Eq("customer_id", expectedUserID), sort.Asc("created_at"))
			biotestPaginatorResponse := []entities.Biotest{
				{ID: 1, BiotestUUID: "uuid1", CreatedAt: time.Now()},
				{ID: 2, BiotestUUID: "uuid2", CreatedAt: time.Now()},
			}
			pageResponse := &paginator.Paginator{
				TotalPages:   2,
				CurrentPage:  pageNumberAsInt,
				PreviousPage: 0,
				NextPage:     2,
				Data:         biotestPaginatorResponse,
			}
			biotestHolder := []entities.Biotest{}
			paginatorMock.On(
				"CreatePagination",
				ctx,
				entities.BiotestTable,
				&biotestHolder,
				&pageSort,
			).Return(pageResponse, nil)

			got, err := biotestRepository.GetAllUserBiotestByUserUUID(ctx, &pageSort, expectedUserUUID)

			Expect(err).ToNot(HaveOccurred())
			Expect(got.Data.([]entities.Biotest)).To(HaveLen(2))
		})

		Context("if user does not exist", func() {
			It("should return a repositories.NotFoundError", func() {
				expectedUserUUID := fake.UUID().V4()
				repo.ExpectFind(
					where.Eq("user_uuid", expectedUserUUID),
				).Error(rel.ErrNotFound)
				pageNumber := fake.Float64(2, 1, 2)
				pageSort := paginator.PageSort{
					Page: pageNumber,
				}

				got, err := biotestRepository.GetAllUserBiotestByUserUUID(ctx, &pageSort, expectedUserUUID)

				Expect(err).To(BeAssignableToTypeOf(repositories.NotFoundError{}))
				Expect(got).To(BeNil())
			})
		})
	})

	Describe("GetAllUserBiotestByUserUUIDAsCatalog", func() {
		It("should return biotest data as catalog", func() {
			expectedUserUUID := fake.UUID().V4()
			expectedUserID := fake.Int32()
			userFindReturn := entities.User{
				ID:       expectedUserID,
				UserUUID: expectedUserUUID,
			}
			repo.ExpectFind(
				where.Eq("user_uuid", expectedUserUUID),
			).Result(userFindReturn)
			pageNumber := fake.Float64(2, 1, 2)
			pageNumberAsInt := int(pageNumber)
			pageSort := paginator.PageSort{
				Page: pageNumber,
			}
			pageSort.SetFiltersQueries(
				where.Eq("customer_id", expectedUserID),
				sort.Asc("created_at"),
				rel.Select("biotest_uuid", "created_at").From(entities.BiotestTable),
			)
			biotestPaginatorResponse := []repositories.BiotestDetails{
				{BiotestUUID: "uuid1", CreatedAt: time.Now()},
				{BiotestUUID: "uuid2", CreatedAt: time.Now()},
			}
			pageResponse := &paginator.Paginator{
				TotalPages:   2,
				CurrentPage:  pageNumberAsInt,
				PreviousPage: 0,
				NextPage:     2,
				Data:         biotestPaginatorResponse,
			}
			biotestHolder := []repositories.BiotestDetails{}
			paginatorMock.On(
				"CreatePagination",
				ctx,
				entities.BiotestTable,
				&biotestHolder,
				&pageSort,
			).Return(pageResponse, nil)

			got, err := biotestRepository.GetAllUserBiotestByUserUUIDAsCatalog(ctx, &pageSort, expectedUserUUID)

			Expect(err).ToNot(HaveOccurred())
			Expect(got.Data.([]repositories.BiotestDetails)).To(HaveLen(2))
		})

		Context("if user does not exist", func() {
			It("should return a repositories.NotFoundError", func() {
				expectedUserUUID := fake.UUID().V4()
				repo.ExpectFind(
					where.Eq("user_uuid", expectedUserUUID),
				).Error(rel.ErrNotFound)
				pageNumber := fake.Float64(2, 1, 2)
				pageSort := paginator.PageSort{
					Page: pageNumber,
				}

				got, err := biotestRepository.GetAllUserBiotestByUserUUIDAsCatalog(ctx, &pageSort, expectedUserUUID)

				Expect(err).To(BeAssignableToTypeOf(repositories.NotFoundError{}))
				Expect(got).To(BeNil())
			})
		})
	})

	Describe("GetComparitionDataByUserUUID", func() {
		It("Should return user comparition data", func() {
			expectedUserUUID := fake.UUID().V4()
			expectedUserID := fake.Int32()
			userReturned := entities.User{
				ID:       expectedUserID,
				UserUUID: expectedUserUUID,
			}
			repo.ExpectFind(
				where.Eq("user_uuid", expectedUserUUID),
			).Result(userReturned)
			biotestDetails := []repositories.BiotestDetails{
				{BiotestUUID: "uuid1", CreatedAt: time.Now()},
				{BiotestUUID: "uuid2", CreatedAt: time.Now()},
			}
			repo.ExpectFindAll(
				where.Eq("customer_id", expectedUserID),
				rel.Select("biotest_uuid", "created_at").From(entities.BiotestTable),
				sort.Asc("created_at"),
			).Result(biotestDetails)
			biotestResponses := []entities.Biotest{
				{ID: 1, BiotestUUID: "uuid1", CreatedAt: time.Now()},
				{ID: 2, BiotestUUID: "uuid2", CreatedAt: time.Now()},
			}
			repo.ExpectFind(
				where.Eq("customer_id", expectedUserID),
				sort.Asc("created_at"),
			).Result(biotestResponses[0])
			repo.ExpectFind(
				where.Eq("customer_id", expectedUserID),
				sort.Desc("created_at"),
			).Result(biotestResponses[1])

			got, err := biotestRepository.GetComparitionDataByUserUUID(ctx, expectedUserUUID)

			Expect(err).ToNot(HaveOccurred())
			Expect(*got.AllBiotestsDetails).To(HaveLen(2))
		})

		Context("if user does not exist", func() {
			It("should return a repositories.NotFoundError", func() {
				expectedUserUUID := fake.UUID().V4()
				repo.ExpectFind(
					where.Eq("user_uuid", expectedUserUUID),
				).NotFound()

				got, err := biotestRepository.GetComparitionDataByUserUUID(ctx, expectedUserUUID)

				Expect(err).To(BeAssignableToTypeOf(repositories.NotFoundError{}))
				Expect(got).To(BeNil())
			})
		})

		Context("if user has no biotests", func() {
			It("should return a respositories.NotFoundError", func() {
				expectedUserUUID := fake.UUID().V4()
				expectedUserID := fake.Int32()
				userReturned := entities.User{
					ID:       expectedUserID,
					UserUUID: expectedUserUUID,
				}
				repo.ExpectFind(
					where.Eq("user_uuid", expectedUserUUID),
				).Result(userReturned)
				biotestDetails := []repositories.BiotestDetails{}
				repo.ExpectFindAll(
					where.Eq("customer_id", expectedUserID),
					rel.Select("biotest_uuid", "created_at").From(entities.BiotestTable),
					sort.Asc("created_at"),
				).Result(biotestDetails)

				got, err := biotestRepository.GetComparitionDataByUserUUID(ctx, expectedUserUUID)

				Expect(err).To(BeAssignableToTypeOf(repositories.NotFoundError{}))
				Expect(got).To(BeNil())
			})
		})

		Context("if user has only one biotest", func() {
			It("should return LastBiotest as nil", func() {
				expectedUserUUID := fake.UUID().V4()
				expectedUserID := fake.Int32()
				userReturned := entities.User{
					ID:       expectedUserID,
					UserUUID: expectedUserUUID,
				}
				repo.ExpectFind(
					where.Eq("user_uuid", expectedUserUUID),
				).Result(userReturned)

				biotestDetails := []repositories.BiotestDetails{
					{BiotestUUID: "uuid1", CreatedAt: time.Now()},
				}
				repo.ExpectFindAll(
					where.Eq("customer_id", expectedUserID),
					rel.Select("biotest_uuid", "created_at").From(entities.BiotestTable),
					sort.Asc("created_at"),
				).Result(biotestDetails)
				biotestResponse := entities.Biotest{
					ID: 1, BiotestUUID: "uuid1", CreatedAt: time.Now(),
				}
				repo.ExpectFind(
					where.Eq("customer_id", expectedUserID),
					sort.Asc("created_at"),
				).Result(biotestResponse)

				got, err := biotestRepository.GetComparitionDataByUserUUID(ctx, expectedUserUUID)

				Expect(err).ToNot(HaveOccurred())
				Expect(got.LastBiotest).To(BeNil())
			})
		})
	})

	Describe("SaveBiotest", func() {
		It("should save a new biotest", func() {
			UUIDForBiotest := fake.UUID().V4()
			uuidGen.On("New").Return(UUIDForBiotest)
			biotestToSave := entities.Biotest{}
			repo.ExpectTransaction(func(r *reltest.Repository) {
				repo.ExpectInsert().ForType("entities.Biotest").Success()
				repo.ExpectInsert().ForType("entities.HigherMuscleDensity").Success()
				repo.ExpectInsert().ForType("entities.LowerMuscleDensity").Success()
				repo.ExpectInsert().ForType("entities.SkinFolds").Success()
			})

			err := biotestRepository.SaveBiotest(ctx, &biotestToSave)

			Expect(err).ToNot(HaveOccurred())
			Expect(biotestToSave.ID).ToNot(BeZero())
			Expect(biotestToSave.BiotestUUID).To(Equal(UUIDForBiotest))
		})

		Context("if a insertion error occurre", func() {
			It("should return the original error", func() {
				UUIDForBiotest := fake.UUID().V4()
				uuidGen.On("New").Return(UUIDForBiotest)
				biotestToSave := entities.Biotest{}
				repo.ExpectTransaction(func(r *reltest.Repository) {
					repo.ExpectInsert().ForType("entities.Biotest").Error(fmt.Errorf("an error occured"))
				})

				err := biotestRepository.SaveBiotest(ctx, &biotestToSave)

				Expect(err).To(HaveOccurred())
				Expect(biotestToSave.BiotestUUID).To(BeEmpty())
			})
		})
	})

	Describe("UpdateBiotest", func() {
		It("should update a entire biotest", func() {
			biotestToUpdate := entities.Biotest{}
			repo.ExpectTransaction(func(r *reltest.Repository) {
				r.ExpectUpdate().ForType("entities.Biotest").Success()
				r.ExpectUpdate().ForType("entities.HigherMuscleDensity").Success()
				r.ExpectUpdate().ForType("entities.LowerMuscleDensity").Success()
				r.ExpectUpdate().ForType("entities.SkinFolds").Success()
			})

			err := biotestRepository.UpdateBiotest(ctx, &biotestToUpdate)

			Expect(err).ToNot(HaveOccurred())

		})
	})
})
