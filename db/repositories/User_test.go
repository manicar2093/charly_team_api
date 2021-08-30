package repositories

import (
	"testing"
	"time"

	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	user := entities.User{
		Name:     "user",
		LastName: "repo",
		RoleID:   2,
		Email:    "testing@user-repo.com",
		Birthday: time.Now(),
	}
	repository := NewUserRepositoryGorm(DB)

	err := repository.Save(&user)

	assert.Nil(t, err, "should no be error")
	assert.Greater(t, user.ID, int32(0), "user was not saved")

}
