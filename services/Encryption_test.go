package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	pu := NewPasswordService()
	pass := "testingpass"
	p, e := pu.HashPassword([]byte(pass))

	assert.Nil(t, e, "No debió regresar error!")
	assert.False(t, pass == p, " Debió regresar un hash del password")
}

// TestValidatePassword checa el funcionamiento correcto de la validación de la contraseña
func TestValidatePassword(t *testing.T) {
	pu := NewPasswordService()
	pass := "testingpass"
	p, _ := pu.HashPassword([]byte(pass))
	e := pu.ValidatePassword(p, pass)
	assert.Nil(t, e, "No debió regresar error. La contraseña es valida")
}
