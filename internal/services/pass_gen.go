package services

import (
	"github.com/manicar2093/charly_team_api/internal/config"
	"github.com/sethvargo/go-password/password"
)

// PassGen is a helper interface to wrap password generator
type (
	PassGen interface {
		Generate() (string, error)
	}
	// Deprecated: Should user DefaultPasswordGenerator until new implementation
	PasswordGenerator struct{}
	// DefaultPasswordGenerator generates a default pass for user creation
	DefaultPasswordGenerator struct{}
)

func (p PasswordGenerator) Generate() (string, error) {
	pass, err := password.Generate(
		config.PassLen,
		config.PassNumDigits,
		config.PassNumSymbols,
		false,
		false,
	)
	if err != nil {
		return "", err
	}
	return pass, nil
}

func (c *DefaultPasswordGenerator) Generate() (string, error) {
	return "Rm7r5P<c", nil
}
