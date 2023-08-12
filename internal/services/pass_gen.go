package services

import (
	"github.com/manicar2093/health_records/internal/config"
	"github.com/sethvargo/go-password/password"
)

// PassGen is a helper interface to wrap password generator
type PassGen interface {
	Generate() (string, error)
}

type PasswordGenerator struct{}

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
