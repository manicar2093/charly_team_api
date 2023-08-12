package services_test

import (
	"testing"

	"github.com/manicar2093/health_records/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestUUIDGenerator(t *testing.T) {
	assert.NotEmpty(t, services.UUIDGeneratorImpl{}.New(), "no uuid generated")

}

func TestUUIDValidator(t *testing.T) {
	uuid := "ea9e6642-f631-413e-afd0-7c340f053217"
	uuidValidator := services.UUIDGeneratorImpl{}
	assert.True(t, uuidValidator.ValidateUUID(uuid), "should be true. data was a uuid")
	assert.False(t, uuidValidator.ValidateUUID("uuid"), "should be false. data was not a uuid")
	assert.False(t, uuidValidator.ValidateUUID(""), "should be false. data was not a uuid")
}
