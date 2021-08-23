package services

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PasswordService interface {
	HashPassword(password []byte) (string, error)
	ValidatePassword(hashed, password string) error
}

type PasswordServiceImpl struct{}

func NewPasswordService() PasswordService {
	return &PasswordServiceImpl{}
}

// HashPassword crea un hash con el algoritmo seleccionado para el sistema
func (p PasswordServiceImpl) HashPassword(password []byte) (string, error) {
	b, e := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if e != nil {
		return "", fmt.Errorf("an error occurred creating user's password")
	}
	return string(b), nil
}

// ValidatePassword valida que la contraseña recibida coincida con el hash proporcionado
func (p PasswordServiceImpl) ValidatePassword(hashed, password string) error {
	e := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if e != nil {
		return fmt.Errorf("invalid password")
	}
	return nil
}
