package token

import (
	"github.com/golang-jwt/jwt"
)

var (
	tokenCreationKey = []byte("6df2c7f0-72c6-4bc9-a06d-02d3a1e868d3")
)

type (
	TokenCreateable interface {
		Gen(claims map[string]interface{}) (string, error)
	}
	Creator struct{}
)

func NewCreator() *Creator {
	return &Creator{}
}

func (c *Creator) Gen(claims map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))

	return token.SignedString(tokenCreationKey)

}
