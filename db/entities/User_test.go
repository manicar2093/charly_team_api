package entities

import (
	"context"
	"testing"
	"time"

	"github.com/go-rel/rel/where"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
)

var fake = faker.New()

func TestCustomerEntity(t *testing.T) {
	user := User{
		BiotypeID:     null.IntFrom(1),
		BoneDensityID: null.IntFrom(1),
		RoleID:        1,
		GenderID:      null.IntFrom(1),
		UserUUID:      fake.UUID().V4(),
		Name:          "Test",
		LastName:      "Test",
		Email:         "test@test.com",
		Birthday:      time.Now(),
	}
	ctx := context.Background()

	DB.Insert(ctx, &user)

	assert.NotEmpty(t, user.ID, "ID should not be empty. Customer was not created")
	DB.Delete(ctx, &user)

}

func TestCustomerEntity_RoleLoad(t *testing.T) {
	ctx := context.Background()
	userToSave := User{
		BiotypeID:     null.IntFrom(1),
		BoneDensityID: null.IntFrom(1),
		RoleID:        1,
		GenderID:      null.IntFrom(1),
		UserUUID:      fake.UUID().V4(),
		Name:          "Test",
		LastName:      "Test",
		Email:         "test@test.com",
		Birthday:      time.Now(),
	}
	DB.Insert(ctx, &userToSave)
	var userFound User

	DB.Find(ctx, &userFound, where.Eq("id", userToSave.ID))

	assert.NotEmpty(t, userFound.Role.Description, "Role was not load correctly")
	DB.Delete(ctx, &userToSave)
	DB.Delete(ctx, &userFound)

}
