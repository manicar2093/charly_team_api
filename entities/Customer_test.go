package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCustomerEntity(t *testing.T) {

	customer := Customer{
		Biotype:     Biotype{ID: 1},
		BoneDensity: BoneDensity{ID: 1},
		Role:        Role{ID: 1},
		Name:        "Test",
		LastName:    "Test",
		Email:       "test@test.com",
		Password:    "12345678",
		Birthday:    time.Now(),
	}

	DB.Create(&customer)

	assert.NotEmpty(t, customer.ID, "ID should not be empty. Customer was not created")

	DB.Delete(&customer)

}
