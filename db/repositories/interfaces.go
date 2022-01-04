package repositories

import (
	"context"

	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/paginator"
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
}
