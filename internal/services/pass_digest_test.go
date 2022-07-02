package services_test

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/manicar2093/charly_team_api/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestPassDigest(t *testing.T) {
	dig := services.BcryptImpl{}
	pass := faker.Password()
	pass, err := dig.Digest(pass)
	assert.Nil(t, err, "should not be error")
	assert.NotEmpty(t, pass, "pass is empty. not generated")
}
