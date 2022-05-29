package services_test

import (
	"testing"

	"github.com/manicar2093/charly_team_api/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestPassGen(t *testing.T) {
	gen := services.PasswordGenerator{}
	pass, err := gen.Generate()
	assert.Nil(t, err, "should not be error")
	assert.NotEmpty(t, pass, "pass is empty. not generated")
}
