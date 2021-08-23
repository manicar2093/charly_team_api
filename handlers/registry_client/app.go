package main

import (
	"github.com/manicar2093/charly_team_api/connections"
	"github.com/manicar2093/charly_team_api/entities"
	"github.com/manicar2093/charly_team_api/models"
)

const clientDefaultRole = 3

func createClient(data models.CreateUserRequest, db connections.Repository) error {
	clientEntity := entities.User{
		Name:     data.Name,
		LastName: data.LastName,
		Email:    data.Email,
		Password: data.Password,
		Birthday: data.Birthday,
		RoleID:   clientDefaultRole,
	}

	result := db.Save(clientEntity)
	if result.Error != nil {
		return result.Error
	}

	return nil

}
