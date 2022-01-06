package repositories

import (
	"context"
	"strings"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/paginator"
)

type UserRepositoryRel struct {
	repo      rel.Repository
	paginator paginator.Paginable
}

func NewUserRepositoryRel(repo rel.Repository, paginator paginator.Paginable) *UserRepositoryRel {
	return &UserRepositoryRel{
		repo:      repo,
		paginator: paginator,
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
	parameterLower := strings.ToLower(parameter)

	filter := where.Like(
		"LOWER(email)", "%"+parameterLower+"%",
	).OrLike(
		"LOWER(name)", "%"+parameterLower+"%",
	).OrLike(
		"LOWER(last_name)", "%"+parameterLower+"%",
	)

	usersFound := []entities.User{}
	if err := c.repo.FindAll(ctx, &usersFound, filter); err != nil {
		return nil, err
	}

	return &usersFound, nil
}

func (c *UserRepositoryRel) FindAllUsers(ctx context.Context, pageSort *paginator.PageSort) (*paginator.Paginator, error) {
	usersFound := []entities.User{}
	return c.paginator.CreatePagination(ctx, entities.UserTable, &usersFound, pageSort)
}

func (c *UserRepositoryRel) SaveUser(ctx context.Context, user *entities.User) error {
	return c.repo.Transaction(ctx, func(ctx context.Context) error {
		return c.repo.Insert(ctx, user)
	})
}

func (c *UserRepositoryRel) UpdateUser(ctx context.Context, user *entities.User) error {
	return c.repo.Transaction(ctx, func(ctx context.Context) error {
		return c.repo.Update(ctx, user)
	})
}
