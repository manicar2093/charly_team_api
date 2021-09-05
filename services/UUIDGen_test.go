package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUIDGenerator(t *testing.T) {
	assert.NotEmpty(t, UUIDGeneratorImpl{}.New(), "no uuid generated")

}
