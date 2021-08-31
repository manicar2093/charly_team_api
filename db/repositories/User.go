package repositories

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/db/entities"
)

type UserRepository interface {
	Save(context.Context, *entities.User) error
}

type UserRepositoryImpl struct {
	provider rel.Repository
}

func NewUserRepositoryImpl(provider rel.Repository) UserRepository {
	return UserRepositoryImpl{provider: provider}
}

func (c UserRepositoryImpl) Save(ctx context.Context, user *entities.User) error {

	err := c.provider.Insert(ctx, user)

	if err != nil {
		return err
	}

	return nil

}
