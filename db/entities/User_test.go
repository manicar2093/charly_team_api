package entities

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCustomerEntity(t *testing.T) {

	user := User{
		Biotype:     Biotype{ID: 1},
		BoneDensity: BoneDensity{ID: 1},
		Role:        Role{ID: 1},
		Name:        "Test",
		LastName:    "Test",
		Email:       "test@test.com",
		Birthday:    time.Now(),
	}

	ctx := context.Background()

	DB.Insert(ctx, &user)

	assert.NotEmpty(t, user.ID, "ID should not be empty. Customer was not created")

	DB.Delete(ctx, &user)

}
