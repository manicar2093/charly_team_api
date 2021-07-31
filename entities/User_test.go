package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserEntity(t *testing.T) {
	user := User{
		Role:     Role{ID: 1},
		Name:     "Manuel",
		LastName: "Carbajal",
		Username: "manicar2093",
		Password: "12345678",
	}
	DB.Create(&user)
	assert.NotEmpty(t, user.ID, "ID should not be 0")
	DB.Delete(&user)
}
