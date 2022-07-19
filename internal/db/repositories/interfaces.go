package repositories

import (
	"context"

	"github.com/manicar2093/charly_team_api/internal/db/entities"
	"github.com/manicar2093/charly_team_api/internal/db/paginator"
)

type BiotestRepository interface {
	FindBiotestByUUID(
		ctx context.Context,
		biotestUUID string,
	) (*entities.Biotest, error)
	GetAllUserBiotestByUserUUID(
		ctx context.Context,
		pageSort *paginator.PageSort,
		userUUID string,
	) (*paginator.Paginator, error)
	GetAllUserBiotestByUserUUIDAsCatalog(
		ctx context.Context,
		pageSort *paginator.PageSort,
		userUUID string,
	) (*paginator.Paginator, error)
	GetComparitionDataByUserUUID(
		ctx context.Context,
		userUUID string,
	) (*BiotestComparisionResponse, error)
	SaveBiotest(
		ctx context.Context,
		biotest *entities.Biotest,
	) error
	UpdateBiotest(
		ctx context.Context,
		biotest *entities.Biotest,
	) error
}

type UserRepository interface {
	FindUserByUUID(
		ctx context.Context,
		userUUID string,
	) (*entities.User, error)
	FindUserLikeEmailOrNameOrLastName(
		ctx context.Context,
		parameter string,
	) (*[]entities.User, error)
	FindAllUsers(
		ctx context.Context,
		pageSort *paginator.PageSort,
	) (*paginator.Paginator, error)
	SaveUser(
		ctx context.Context,
		user *entities.User,
	) error
	UpdateUser(
		ctx context.Context,
		user *entities.User,
	) error
	FindUserByEmail(
		ctx context.Context,
		email string,
	) (*entities.User, error)
}
