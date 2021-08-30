package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassGen(t *testing.T) {
	gen := PasswordGenerator{}
	pass, err := gen.Generate()
	assert.Nil(t, err, "should not be error")
	assert.NotEmpty(t, pass, "pass is empty. not generated")
}
