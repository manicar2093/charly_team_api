package repositories

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/paginator"
)

type UserRepositoryRel struct {
	repo rel.Repository
}

func NewUserRepositoryRel(repo rel.Repository) *UserRepositoryRel {
	return &UserRepositoryRel{
		repo: repo,
	}
}

func (c *UserRepositoryRel) FindUserByUUID(ctx context.Context, userUUID string) (*entities.User, error) {
	userFound := entities.User{}
	if err := c.repo.Find(ctx, &userFound, where.Eq("user_uuid", userUUID)); err != nil {
		switch err.(type) {
		case rel.NotFoundError:
			return nil, NotFoundError{Entity: "User", Identifier: userUUID}
		}
		return nil, err
	}
	return &userFound, nil
}

func (c *UserRepositoryRel) FindUserLikeEmailOrNameOrLastName(ctx context.Context, parameter string) (*[]entities.User, error) {
	panic("not implemented") // TODO: Implement
}

func (c *UserRepositoryRel) FindAllUsers(ctx context.Context, pageSort *paginator.PageSort) (*paginator.Paginator, error) {
	panic("not implemented") // TODO: Implement
}

func (c *UserRepositoryRel) SaveUser(ctx context.Context, user *entities.User) error {
	panic("not implemented") // TODO: Implement
}

func (c *UserRepositoryRel) UpdateUser(ctx context.Context, user *entities.User) error {
	panic("not implemented") // TODO: Implement
}
