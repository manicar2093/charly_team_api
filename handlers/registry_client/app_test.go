package main

import (
	"testing"
	"time"

	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/manicar2093/charly_team_api/models"
)

func TestCreateClient(t *testing.T) {

	var dbMock mocks.Repository
	data := models.CreateUserRequest{
		Name:     "testing",
		LastName: "testing",
		Email:    "testing@testing.com",
		Password: "testing",
		Birthday: time.Now(),
	}

	t.Run("should create a client with the correct rol", func(t *testing.T) {

	})
}
