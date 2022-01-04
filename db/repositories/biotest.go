package repositories

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/paginator"
)

type BiotestRepositoryRel struct {
	repo rel.Repository
}

func NewBiotestRepositoryRel(repo rel.Repository) *BiotestRepositoryRel {
	return &BiotestRepositoryRel{repo: repo}
}

func (c *BiotestRepositoryRel) FindBiotestByUUID(ctx context.Context, biotestUUID string) (*entities.Biotest, error) {
	biotest := entities.Biotest{}
	if err := c.repo.Find(ctx, &biotest, where.Eq("biotest_uuid", biotestUUID)); err != nil {
		return nil, err
	}
	return &biotest, nil
}

func (c *BiotestRepositoryRel) GetAllUserBiotestByUserUUID(ctx context.Context, userUUID string) (*paginator.Paginator, error) {
	panic("not implemented") // TODO: Implement
}

func (c *BiotestRepositoryRel) GetAllUserBiotestByUserUUIDAsCatalog(ctx context.Context, userUUID string) (*paginator.Paginator, error) {
	panic("not implemented") // TODO: Implement
}

func (c *BiotestRepositoryRel) GetComparitionDataByUserUUID(ctx context.Context, userUUID string) (*BiotestComparisionResponse, error) {
	panic("not implemented") // TODO: Implement
}
