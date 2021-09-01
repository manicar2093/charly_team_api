package entities

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
)

func TestCustomerEntity(t *testing.T) {

	user := User{
		BiotypeID:     null.IntFrom(1),
		BoneDensityID: null.IntFrom(1),
		RoleID:        1,
		GenderID:      1,
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
