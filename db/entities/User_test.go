package entities

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCustomerEntity(t *testing.T) {

	user := User{
		BiotypeID:     sql.NullInt32{Valid: true, Int32: 1},
		BoneDensityID: sql.NullInt32{Valid: true, Int32: 1},
		RoleID:        1,
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
