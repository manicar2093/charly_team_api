package repositories

import (
	"github.com/manicar2093/charly_team_api/db/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *entities.User) error
}

type UserRepositoryGorm struct {
	provider *gorm.DB
}

func NewUserRepositoryGorm(provider *gorm.DB) UserRepository {
	return UserRepositoryGorm{provider: provider}
}

func (c UserRepositoryGorm) Save(user *entities.User) error {

	err := c.provider.Save(user)

	if err.Error != nil {
		return err.Error
	}

	return nil

}
